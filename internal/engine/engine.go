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
	Match   *Match
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
	case "TICK":
		return e.handleTick()
	default:
		return ActionResult{
			Success: false,
			Message: fmt.Sprintf("Unknown action: %s\n", action),
		}
	}
}

func (e *Engine) handleEnter(playerId string, skillRating int, region types.Region) ActionResult {
	if _, ok := e.WaitingRoom[playerId]; ok {
		return ActionResult{
			Success: false,
			Message: "Player already exists in waiting room",
		}
	}
	player, err := NewPlayer(playerId, skillRating, region)
	if err != nil {
		return ActionResult{
			Success: false,
			Message: err.Error(),
		}
	}

	e.WaitingRoom[player.Id] = player

	return ActionResult{
		Success: true,
		Message: "Player added to waiting room",
		Player:  player,
	}
}

func (e *Engine) handleTick() ActionResult {
	var players []Player
	for _, value := range e.WaitingRoom {
		players = append(players, *value)
	}
	match := NewMatch(players)

	return ActionResult{
		Success: true,
		Message: "Started Match",
		Match:   match,
	}
}

func (e *Engine) PrintWaitingRoom() {
	fmt.Println("  Waiting room:")
	for id, player := range e.WaitingRoom {
		fmt.Printf("    %s: %+v\n", id, *player)
	}
}
