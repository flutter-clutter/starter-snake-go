package game

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
	return abs(coord.X-other.X) + abs(coord.Y-other.Y)
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
		return true
	}
	for _, bodyCoord := range battlesnake.Body /*[:len(battlesnake.Body)-1]*/ {
		if currentCoord.equals(bodyCoord) {
			return true
		}
	}

	return false
}

func (currentCoord Coord) isSafe(battlesnake Battlesnake, board Board) bool {
	return !currentCoord.isOutsideOfArea(board) && !currentCoord.isInSnakes(board)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (currentCoord Coord) isAtEdge(battlesnake Battlesnake, board Board) bool {
	if currentCoord.X == 0 || currentCoord.X == board.Width-1 {
		return true
	}

	if currentCoord.Y == 0 || currentCoord.Y == board.Height-1 {
		return true
	}

	return false
}

type ByDistance struct {
	SnakePosition Coord
	Coords        []Coord
}

func (a ByDistance) Len() int      { return len(a.Coords) }
func (a ByDistance) Swap(i, j int) { a.Coords[i], a.Coords[j] = a.Coords[j], a.Coords[i] }
func (a ByDistance) Less(i, j int) bool {
	return a.Coords[i].distanceToOther(a.SnakePosition) < a.Coords[j].distanceToOther(a.SnakePosition)
}
