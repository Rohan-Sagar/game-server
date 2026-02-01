package engine

import (
	"fmt"

	"github.com/rohan-sagar/game-server/internal/types"
)

type Engine struct {
	WaitingRoom map[string]*Player
}

type ActionResult struct {
	Success bool
	Message string
	Player  *Player
}

func NewEngine() *Engine {
	return &Engine{
		WaitingRoom: make(map[string]*Player),
	}
}

func (e *Engine) HandleAction(action, playerId string, skillRating int, region types.Region) ActionResult {
	switch action {
	case "ENTER":
		return e.handleEnter(playerId, skillRating, region)
	default:
		return ActionResult{
			Success: false,
			Message: fmt.Sprintf("Unknown action: %s\n", action),
		}
	}
}

func (e *Engine) handleEnter(playerId string, skillRating int, region types.Region) ActionResult {
	player := NewPlayer(playerId, skillRating, region)
	e.WaitingRoom[player.Id] = player

	return ActionResult{
		Success: true,
		Message: "Player added to waiting room",
		Player:  player,
	}
}

func (e *Engine) PrintWaitingRoom() {
	fmt.Println("  Waiting room:")
	for id, player := range e.WaitingRoom {
		fmt.Printf("    %s: %+v\n", id, *player)
	}
}
