package game

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

	var minFoodDistanceCoord Coord = Coord{1000, 1000}

	for i := 0; i < amountOfFood; i++ {
		if battlesnake.Head.distanceToOther(minFoodDistanceCoord) > battlesnake.Head.distanceToOther(board.Food[i]) {
			minFoodDistanceCoord = Coord{
				board.Food[i].X,
				board.Food[i].Y,
			}
		}
	}

	if minFoodDistanceCoord.X > battlesnake.Head.X {
		return SnakeDirection.RIGHT
	}

	if minFoodDistanceCoord.X < battlesnake.Head.X {
		return SnakeDirection.LEFT
	}

	if minFoodDistanceCoord.Y > battlesnake.Head.Y {
		return SnakeDirection.UP
	}

	if minFoodDistanceCoord.Y < battlesnake.Head.Y {
		return SnakeDirection.DOWN
	}

	return SnakeDirection.UP
}

func getSafeMove(battlesnake Battlesnake, board Board) SnakeDirectionType {
	for _, v := range possibleMoves {
		if battlesnake.Head.newCoordFromMove(v).isSafe(battlesnake, board) {
			return v
		}
	}

	println("No safe move found")
	return SnakeDirection.UP
}
