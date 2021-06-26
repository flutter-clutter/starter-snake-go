package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/flutter-clutter/starter-snake-go/game"
)

func TestIndexReturnsCorrectResponse(t *testing.T) {
	server := httptest.NewServer(setupRouter())

	resp, err := http.Get(server.URL)

	if err != nil {
		t.Fatal(err)
		return
	}

	if resp.StatusCode != 200 {
		t.Errorf("Expected status code: 200. Got %d", resp.StatusCode)
		return
	}
}
func TestStartReturnsCorrectResponse(t *testing.T) {
	server := httptest.NewServer(setupRouter())
	request := createGameRequest()
	resp := sendGameRequest(t, request, server.URL, "start")

	if resp.StatusCode != 200 {
		t.Errorf("Expected status code: 200. Got %d", resp.StatusCode)
		return
	}
}
func TestMoveReturnsCorrectResponse(t *testing.T) {
	server := httptest.NewServer(setupRouter())

	request := createGameRequest()
	resp := sendGameRequest(t, request, server.URL, "start")
	resp.Body.Close()

	request = createGameRequest()
	resp = sendGameRequest(t, request, server.URL, "move")
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Errorf("Expected status code: 200. Got %d", resp.StatusCode)
		return
	}

	bytes, err := io.ReadAll(resp.Body)

	if err != nil {
		t.Fatal(err)
		return
	}

	var response MoveResponse
	err = json.Unmarshal(bytes, &response)

	if err != nil {
		t.Error(err)
		return
	}

	if response.Move != game.SnakeDirection.UP && response.Move != game.SnakeDirection.DOWN && response.Move != game.SnakeDirection.LEFT && response.Move != game.SnakeDirection.RIGHT {
		t.Errorf("Got unexpected direction: %s", response.Move)
		return
	}
}

func createGameRequest() GameRequest {
	var snakeGame Game = Game{
		"1",
		int32(60),
	}

	var snake game.Battlesnake = game.Battlesnake{
		ID:     "1",
		Name:   "Battlesnake",
		Health: int32(100),
		Body:   []game.Coord{{X: 0, Y: 0}},
		Head:   game.Coord{X: 0, Y: 1},
		Length: int32(2),
		Shout:  "",
	}

	var board game.Board = game.Board{
		Height: 10,
		Width:  10,
		Food:   []game.Coord{},
		Snakes: []game.Battlesnake{snake},
	}

	var request GameRequest = GameRequest{
		Game:  snakeGame,
		Turn:  1,
		Board: board,
		You:   snake,
	}

	return request
}

func sendGameRequest(t *testing.T, request GameRequest, url string, route string) *http.Response {
	requestBytes, err := json.Marshal(request)

	if err != nil {
		t.Fatal(requestBytes)
	}

	r := bytes.NewReader(requestBytes)

	resp, err := http.Post(fmt.Sprintf("%s/%s", url, route), "application/json", r)
	if err != nil {
		t.Fatal(err)
		return nil
	}

	return resp
}
