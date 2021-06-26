package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/flutter-clutter/starter-snake-go/game"
)

var snake *game.StrategicBattlesnake

type Game struct {
	ID      string `json:"id"`
	Timeout int32  `json:"timeout"`
}

type BattlesnakeInfoResponse struct {
	APIVersion string `json:"apiversion"`
	Author     string `json:"author"`
	Color      string `json:"color"`
	Head       string `json:"head"`
	Tail       string `json:"tail"`
}

type GameRequest struct {
	Game  Game             `json:"game"`
	Turn  int              `json:"turn"`
	Board game.Board       `json:"board"`
	You   game.Battlesnake `json:"you"`
}

type MoveResponse struct {
	Move  game.SnakeDirectionType `json:"move"`
	Shout string                  `json:"shout,omitempty"`
}

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

	snake = &game.StrategicBattlesnake{
		Snake:    request.You,
		Action:   game.CollectNearestFood{},
		Strategy: game.NearestFoodStrategy{},
	}

	w.WriteHeader(http.StatusOK)
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

	snake.Snake = request.You
	snake.Strategy = game.NearestFoodStrategy{}
	snake.Action = snake.Strategy.ExecuteNextStep(snake.Snake, request.Board)

	response := MoveResponse{
		Move: snake.Action.Execute(request.You, request.Board),
	}

	//fmt.Printf("MOVE: %s\n", response.Move)

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Fatal(err)
	}
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

func Start() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}

	s := &http.Server{
		Addr:           fmt.Sprintf(":%s", port),
		Handler:        setupRouter(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	fmt.Printf("Starting Battlesnake Server at http://0.0.0.0:%s...\n", port)

	log.Fatal(s.ListenAndServe())
}

func setupRouter() http.Handler {
	handler := http.NewServeMux()
	handler.HandleFunc("/", HandleIndex)
	handler.HandleFunc("/start", HandleStart)
	handler.HandleFunc("/move", HandleMove)
	handler.HandleFunc("/end", HandleEnd)

	return handler
}
