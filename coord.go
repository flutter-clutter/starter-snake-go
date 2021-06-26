package main

type Coord struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func (currentCoord Coord) isOutsideOfArea(board Board) bool {
	return currentCoord.X < 0 || currentCoord.Y < 0 || currentCoord.X >= board.Width || currentCoord.Y >= board.Height
}

func (current Coord) equals(other Coord) bool {
	return current.X == other.X && current.Y == other.Y
}

func (currentCoord Coord) newCoordFromMove(move SnakeDirectionType) Coord {
	if move == SnakeDirection.UP {
		return Coord{currentCoord.X, currentCoord.Y + 1}
	}
	if move == SnakeDirection.DOWN {
		return Coord{currentCoord.X, currentCoord.Y - 1}
	}
	if move == SnakeDirection.LEFT {
		return Coord{currentCoord.X - 1, currentCoord.Y}
	}
	if move == SnakeDirection.RIGHT {
		return Coord{currentCoord.X + 1, currentCoord.Y}
	}
	return currentCoord
}

func (coord Coord) distanceToOther(other Coord) int {
	return Abs(coord.X-other.X) + Abs(coord.Y-other.Y)
}

func (currentCoord Coord) isInSnakes(board Board) bool {
	for _, snake := range board.Snakes {
		if currentCoord.isInSnake(snake) {
			return true
		}
	}

	return false
}

func (currentCoord Coord) isInSnake(battlesnake Battlesnake) bool {
	if currentCoord.equals(battlesnake.Head) {
		println("Coord is in snake")
		return true
	}
	for _, bodyCoord := range battlesnake.Body /*[:len(battlesnake.Body)-1]*/ {
		if currentCoord.equals(bodyCoord) {
			println("Coord is in snake")
			return true
		}
	}

	return false
}

func (currentCoord Coord) isSafe(battlesnake Battlesnake, board Board) bool {
	return !currentCoord.isOutsideOfArea(board) && !currentCoord.isInSnakes(board)
}
