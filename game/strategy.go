package game

var possibleMoves []SnakeDirectionType = []SnakeDirectionType{SnakeDirection.UP, SnakeDirection.DOWN, SnakeDirection.LEFT, SnakeDirection.RIGHT}

type Strategy interface {
	ExecuteNextStep(Battlesnake, Board) Action
}

type NearestFoodStrategy struct{}

func (s NearestFoodStrategy) ExecuteNextStep(snake Battlesnake, board Board) Action {
	return CollectNearestFood{}
}
