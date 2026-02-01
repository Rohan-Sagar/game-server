package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/rohan-sagar/game-server/internal/engine"
	"github.com/rohan-sagar/game-server/internal/types"
)

func main() {
	mode := os.Getenv("MODE")
	if mode == "" {
		mode = "cli"
	}

	gameEngine := engine.NewEngine()

	switch mode {
	case "cli":
		parseCLI(gameEngine)
	default:
		fmt.Printf("Unknown mode: %s\n", mode)
		os.Exit(1)
	}

}

func parseCLI(e *engine.Engine) {
	action := flag.String("action", "ENTER", "Action")
	player_id := flag.String("player_id", "rohan", "Player ID")
	skill_rating := flag.Int("skill_rating", 1500, "Skill Rating")
	var region types.Region = types.UsEast

	allowedRegions := []string{string(types.UsEast), string(types.UsWest)}
	usage := fmt.Sprintf("Region must be one of the %v", allowedRegions)
	flag.Func("region", usage, func(val string) error {
		for _, allowed := range allowedRegions {
			if val == allowed {
				region = types.Region(val)
				return nil
			}
		}
		return fmt.Errorf("Invalid region: %s", region)
	})

	flag.Parse()

	result := e.HandleAction(*action, *player_id, *skill_rating, region)

	fmt.Printf("%s\n", result.Message)
	if result.Success {
		if result.Player != nil {
			fmt.Printf("  Player: %+v\n", *result.Player)
		}
	}

	if *action == "STATUS" || *action == "ENTER" {
		fmt.Println()
		e.PrintWaitingRoom()
	}
}
