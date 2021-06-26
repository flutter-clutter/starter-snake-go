package game

type SnakeDirectionType string

var SnakeDirection = struct {
	UP    SnakeDirectionType
	DOWN  SnakeDirectionType
	LEFT  SnakeDirectionType
	RIGHT SnakeDirectionType
}{
	UP:    "up",
	DOWN:  "down",
	LEFT:  "left",
	RIGHT: "right",
}
