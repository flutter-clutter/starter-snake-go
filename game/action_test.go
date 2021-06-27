package game

import "testing"

func TestCollectNearestFoodActionWithFoodInRange(t *testing.T) {
	tests := []struct {
		Name        string
		Food        []Coord
		SnakeCoords []Coord
		Expected    SnakeDirectionType
	}{
		{
			Name:        "Go left when food is left",
			Food:        []Coord{{X: 1, Y: 2}},
			SnakeCoords: []Coord{{X: 2, Y: 2}, {X: 2, Y: 1}},
			Expected:    SnakeDirection.LEFT,
		},
		{
			Name:        "Go right when food is right",
			Food:        []Coord{{X: 3, Y: 2}},
			SnakeCoords: []Coord{{X: 2, Y: 2}, {X: 2, Y: 1}},
			Expected:    SnakeDirection.RIGHT,
		},
		{
			Name:        "Go up when food is above",
			Food:        []Coord{{X: 2, Y: 3}},
			SnakeCoords: []Coord{{X: 2, Y: 2}, {X: 2, Y: 1}},
			Expected:    SnakeDirection.UP,
		},
		{
			Name:        "Go down when food is below",
			Food:        []Coord{{X: 2, Y: 0}},
			SnakeCoords: []Coord{{X: 2, Y: 1}, {X: 2, Y: 2}},
			Expected:    SnakeDirection.DOWN,
		},
	}

	action := CollectNearestFood{}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
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
		})
	}
}

func TestApproachBorder(t *testing.T) {
	tests := []struct {
		Name        string
		Obstacles   []Coord
		SnakeCoords []Coord
		Expected    SnakeDirectionType
	}{
		{
			Name:        "With left border being closest, expect to go left",
			SnakeCoords: []Coord{{X: 1, Y: 5}, {X: 1, Y: 6}},
			Expected:    SnakeDirection.LEFT,
		},
		{
			Name:        "With right border being closest, expect to go right",
			SnakeCoords: []Coord{{X: 8, Y: 5}, {X: 8, Y: 6}},
			Expected:    SnakeDirection.RIGHT,
		},
		{
			Name:        "With bottom border being closest, expect to go down",
			SnakeCoords: []Coord{{X: 5, Y: 1}, {X: 6, Y: 1}},
			Expected:    SnakeDirection.DOWN,
		},
		{
			Name:        "With top border being closest, expect to go up",
			SnakeCoords: []Coord{{X: 5, Y: 8}, {X: 6, Y: 8}},
			Expected:    SnakeDirection.UP,
		},
		{
			Name:        "Expect not to target obstacle above snake",
			Obstacles:   []Coord{{X: 5, Y: 9}},
			SnakeCoords: []Coord{{X: 5, Y: 8}, {X: 6, Y: 8}},
			Expected:    SnakeDirection.DOWN,
		},
		{
			Name:        "Expect to follow border when in top right corner",
			SnakeCoords: []Coord{{X: 9, Y: 9}, {X: 9, Y: 8}},
			Expected:    SnakeDirection.LEFT,
		},
		{
			Name:        "Expect to follow border when next to top right corner",
			SnakeCoords: []Coord{{X: 8, Y: 9}, {X: 9, Y: 9}, {X: 9, Y: 8}},
			Expected:    SnakeDirection.LEFT,
		},
		{
			Name:        "Expect to choose safe move when shortest path to border is blocked",
			SnakeCoords: []Coord{{X: 1, Y: 1}, {X: 0, Y: 1}, {X: 0, Y: 0}, {X: 1, Y: 0}, {X: 2, Y: 0}, {X: 2, Y: 1}},
			Expected:    SnakeDirection.UP,
		},
	}

	action := ApproachBorder{}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			var snake Battlesnake = Battlesnake{
				ID:     "1",
				Name:   "Battlesnake",
				Health: int32(100),
				Body:   tt.SnakeCoords[1:],
				Head:   tt.SnakeCoords[0],
				Length: int32(len(tt.SnakeCoords)),
				Shout:  "",
			}

			var battlesnakes []Battlesnake = []Battlesnake{snake}

			if len(tt.Obstacles) > 0 {
				var enemyBody []Coord
				if len(tt.Obstacles) > 1 {
					enemyBody = tt.Obstacles[1:]
				}

				var enemy = Battlesnake{
					ID:     "2",
					Name:   "Battlesnake 2",
					Health: int32(100),
					Body:   enemyBody,
					Head:   tt.Obstacles[0],
					Length: int32(len(tt.Obstacles)),
					Shout:  "",
				}

				battlesnakes = append(battlesnakes, enemy)
			}

			move := action.Execute(
				snake,
				Board{
					Height: 10,
					Width:  10,
					Food:   []Coord{},
					Snakes: battlesnakes,
				},
			)

			if move != tt.Expected {
				t.Errorf("Snake does not move in direction of border (%s), %s instead", tt.Expected, move)
				return
			}
		})
	}
}
