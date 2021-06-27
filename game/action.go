package game

import "sort"

var possibleMoves []SnakeDirectionType = []SnakeDirectionType{SnakeDirection.UP, SnakeDirection.RIGHT, SnakeDirection.DOWN, SnakeDirection.LEFT}

type Action interface {
	Execute(Battlesnake, Board) SnakeDirectionType
}

type CollectNearestFood struct{}

func (CollectNearestFood) Execute(snake Battlesnake, board Board) SnakeDirectionType {
	var move SnakeDirectionType

	move = approachNearestFood(snake, board)

	var newCoord Coord = snake.Head.newCoordFromMove(move)
	if !newCoord.isSafe(snake, board) {
		move = getSafeMove(snake, board)
	}

	return move
}

func approachNearestFood(battlesnake Battlesnake, board Board) SnakeDirectionType {
	var amountOfFood int = len(board.Food)

	if amountOfFood == 0 {
		// TODO: Ungefährlicher Move (keine Wand, nicht selber fressen)
		return SnakeDirection.UP
	}

	return moveTowardsNearestCoord(battlesnake.Head, board.Food)

}

func getSafeMove(battlesnake Battlesnake, board Board) SnakeDirectionType {
	for _, v := range possibleMoves {
		newCoord := battlesnake.Head.newCoordFromMove(v)
		if newCoord.isSafe(battlesnake, board) {
			return v
		}
	}

	println("No safe move found")
	return SnakeDirection.UP
}

func getNextMoveAlongBorder(battlesnake Battlesnake, board Board) SnakeDirectionType {
	/*for _, v := range possibleMoves {
		newCoord := battlesnake.Head.newCoordFromMove(v)
		if newCoord.isSafe(battlesnake, board) && newCoord.isAtEdge(battlesnake, board) {
			return v
		}
	}*/

	if battlesnake.Head.X == 0 {
		return SnakeDirection.DOWN
	}

	if battlesnake.Head.X == board.Width-1 {
		return SnakeDirection.UP
	}

	if battlesnake.Head.Y == 0 {
		return SnakeDirection.LEFT
	}

	if battlesnake.Head.Y == board.Height-1 {
		return SnakeDirection.RIGHT
	}

	println("No safe border move found")
	return getSafeMove(battlesnake, board)
}

type MakeSafeMove struct{}

func (MakeSafeMove) Execute(snake Battlesnake, board Board) SnakeDirectionType {
	return getSafeMove(snake, board)
}

type MakeSafeBorderMove struct{}

func (MakeSafeBorderMove) Execute(snake Battlesnake, board Board) SnakeDirectionType {
	return getNextMoveAlongBorder(snake, board)
}

type FollowBorder struct{}

func (FollowBorder) Execute(snake Battlesnake, board Board) SnakeDirectionType {
	return getNextMoveAlongBorder(snake, board)
}

type ApproachBorder struct{}

func (ApproachBorder) Execute(snake Battlesnake, board Board) SnakeDirectionType {
	safeBorderPieces := createListOfSafeBorderPieces(snake, board)

	byDistance := ByDistance{snake.Head, safeBorderPieces}
	sort.Sort(ByDistance(byDistance))
	sortedBorderPieces := byDistance.Coords

	move := moveTowardsNearestCoord(snake.Head, sortedBorderPieces)

	if snake.Head.newCoordFromMove(move).isSafe(snake, board) {
		return move
	}
	return getSafeMove(snake, board)
}

func createListOfSafeBorderPieces(snake Battlesnake, board Board) []Coord {
	var safeBorderPieces []Coord = []Coord{}

	for i := 0; i < board.Height; i++ {
		leftCoord := Coord{0, i}
		rightCoord := Coord{board.Width - 1, i}

		if leftCoord.isSafe(snake, board) {
			safeBorderPieces = append(safeBorderPieces, leftCoord)
		}

		if rightCoord.isSafe(snake, board) {
			safeBorderPieces = append(safeBorderPieces, rightCoord)
		}
	}

	for i := 0; i < board.Width; i++ {
		upperCoord := Coord{i, 0}
		lowerCoord := Coord{i, board.Height - 1}

		if upperCoord.isSafe(snake, board) {
			safeBorderPieces = append(safeBorderPieces, upperCoord)
		}

		if lowerCoord.isSafe(snake, board) {
			safeBorderPieces = append(safeBorderPieces, lowerCoord)
		}
	}

	return safeBorderPieces
}

func moveTowardsNearestCoord(snakeCoord Coord, allowedCoords []Coord) SnakeDirectionType {
	var minDistanceCoord Coord = allowedCoords[0]

	for _, v := range allowedCoords {
		if snakeCoord.distanceToOther(minDistanceCoord) > snakeCoord.distanceToOther(v) {
			minDistanceCoord = Coord{v.X, v.Y}
		}
	}
	if minDistanceCoord.X > snakeCoord.X {
		return SnakeDirection.RIGHT
	}

	if minDistanceCoord.X < snakeCoord.X {
		return SnakeDirection.LEFT
	}

	if minDistanceCoord.Y > snakeCoord.Y {
		return SnakeDirection.UP
	}

	if minDistanceCoord.Y < snakeCoord.Y {
		return SnakeDirection.DOWN
	}

	return SnakeDirection.UP
}
