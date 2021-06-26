package game

import "testing"

func TestCollectNearestFoodActionWithFoodInRange(t *testing.T) {
	tests := []struct {
		Food        []Coord
		SnakeCoords []Coord
		Expected    SnakeDirectionType
	}{
		{
			Food:        []Coord{{X: 1, Y: 2}},
			SnakeCoords: []Coord{{X: 2, Y: 2}, {X: 2, Y: 1}},
			Expected:    SnakeDirection.LEFT,
		},
		{
			Food:        []Coord{{X: 3, Y: 2}},
			SnakeCoords: []Coord{{X: 2, Y: 2}, {X: 2, Y: 1}},
			Expected:    SnakeDirection.RIGHT,
		},
		{
			Food:        []Coord{{X: 2, Y: 3}},
			SnakeCoords: []Coord{{X: 2, Y: 2}, {X: 2, Y: 1}},
			Expected:    SnakeDirection.UP,
		},
		{
			Food:        []Coord{{X: 2, Y: 0}},
			SnakeCoords: []Coord{{X: 2, Y: 1}, {X: 2, Y: 2}},
			Expected:    SnakeDirection.DOWN,
		},
	}

	for _, tt := range tests {
		action := CollectNearestFood{}

		var snake Battlesnake = Battlesnake{
			ID:     "1",
			Name:   "Battlesnake",
			Health: int32(100),
			Body:   tt.SnakeCoords[1:],
			Head:   tt.SnakeCoords[0],
			Length: int32(len(tt.SnakeCoords)),
			Shout:  "",
		}

		move := action.Execute(
			snake,
			Board{
				Height: 10,
				Width:  10,
				Food:   tt.Food,
				Snakes: []Battlesnake{snake},
			},
		)

		if move != tt.Expected {
			t.Errorf("Snake does not move in direcion of food (%s), %s instead", tt.Expected, move)
			return
		}
	}
}
