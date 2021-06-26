package main

type Battlesnake struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Health int32   `json:"health"`
	Body   []Coord `json:"body"`
	Head   Coord   `json:"head"`
	Length int32   `json:"length"`
	Shout  string  `json:"shout"`
}
