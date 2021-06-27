package game

type Strategy interface {
	ExecuteNextStep(Battlesnake, Board) Action
}

type NearestFoodStrategy struct{}

func (NearestFoodStrategy) ExecuteNextStep(snake Battlesnake, board Board) Action {
	return CollectNearestFood{}
}

type FoodOnlyWhenHealthLow struct{}

func (FoodOnlyWhenHealthLow) ExecuteNextStep(snake Battlesnake, board Board) Action {
	if snake.Health > int32(board.Height) {
		return MakeSafeMove{}
	}
	return CollectNearestFood{}
}

type CircleInnerBorder struct{}

func (CircleInnerBorder) ExecuteNextStep(snake Battlesnake, board Board) Action {
	if snake.Health < int32(board.Height) {
		return CollectNearestFood{}
	}
	if !snake.Head.isAtEdge(snake, board) {
		return ApproachBorder{}
	}

	return FollowBorder{}
}
