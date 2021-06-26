package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Game struct {
	ID      string `json:"id"`
	Timeout int32  `json:"timeout"`
}

type Board struct {
	Height int           `json:"height"`
	Width  int           `json:"width"`
	Food   []Coord       `json:"food"`
	Snakes []Battlesnake `json:"snakes"`
}

type BattlesnakeInfoResponse struct {
	APIVersion string `json:"apiversion"`
	Author     string `json:"author"`
	Color      string `json:"color"`
	Head       string `json:"head"`
	Tail       string `json:"tail"`
}

type GameRequest struct {
	Game  Game        `json:"game"`
	Turn  int         `json:"turn"`
	Board Board       `json:"board"`
	You   Battlesnake `json:"you"`
}

type MoveResponse struct {
	Move  SnakeDirectionType `json:"move"`
	Shout string             `json:"shout,omitempty"`
}

var possibleMoves []SnakeDirectionType = []SnakeDirectionType{SnakeDirection.UP, SnakeDirection.DOWN, SnakeDirection.LEFT, SnakeDirection.RIGHT}

// HandleIndex is called when your Battlesnake is created and refreshed
// by play.battlesnake.com. BattlesnakeInfoResponse contains information about
// your Battlesnake, including what it should look like on the game board.
func HandleIndex(w http.ResponseWriter, r *http.Request) {
	response := BattlesnakeInfoResponse{
		APIVersion: "1",
		Author:     "schnodderfahne",
		Color:      "#ff5978",
		Head:       "gamer",
		Tail:       "mouse",
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Fatal(err)
	}
}

// HandleStart is called at the start of each game your Battlesnake is playing.
// The GameRequest object contains information about the game that's about to start.
// TODO: Use this function to decide how your Battlesnake is going to look on the board.
func HandleStart(w http.ResponseWriter, r *http.Request) {
	request := GameRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Fatal(err)
	}

	if len(request.Board.Snakes) == 1 {
		println("I am the only snake here")
	} else {
		println("Oh, other snakes here, too.")
	}

	// Nothing to respond with here
	fmt.Print("START\n")
}

// HandleMove is called for each turn of each game.
// Valid responses are "up", "down", "left", or "right".
// TODO: Use the information in the GameRequest object to determine your next move.
func HandleMove(w http.ResponseWriter, r *http.Request) {
	request := GameRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Fatal(err)
	}

	var move SnakeDirectionType

	move = approachNearestFood(request.Board, request.You)

	var newCoord Coord = request.You.Head.newCoordFromMove(move)

	fmt.Printf("Is %s safe?\n", move)
	if !newCoord.isSafe(request.You, request.Board) {
		move = getSafeMove(request.You, request.Board)
	}

	response := MoveResponse{
		Move: move,
	}

	fmt.Printf("MOVE: %s\n", response.Move)

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Fatal(err)
	}
}

func getSafeMove(battlesnake Battlesnake, board Board) SnakeDirectionType {
	for _, v := range possibleMoves {
		if battlesnake.Head.newCoordFromMove(v).isSafe(battlesnake, board) {
			fmt.Printf("Safe move: %s\n", v)
			return v
		}
	}

	println("No safe move found")
	return SnakeDirection.UP
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func approachNearestFood(board Board, battlesnake Battlesnake) SnakeDirectionType {
	println("Looking for food ...")
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

// HandleEnd is called when a game your Battlesnake was playing has ended.
// It's purely for informational purposes, no response required.
func HandleEnd(w http.ResponseWriter, r *http.Request) {
	request := GameRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Fatal(err)
	}

	// Nothing to respond with here
	fmt.Print("END\n")
}

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}

	http.HandleFunc("/", HandleIndex)
	http.HandleFunc("/start", HandleStart)
	http.HandleFunc("/move", HandleMove)
	http.HandleFunc("/end", HandleEnd)

	fmt.Printf("Starting Battlesnake Server at http://0.0.0.0:%s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
