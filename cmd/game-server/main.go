package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

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
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println()

	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}
		line := strings.Split(scanner.Text(), " ")
		action := line[0]
		player_id := strings.Split(line[1], "=")[1]
		skill_rating, err := strconv.Atoi(strings.Split(line[2], "=")[1])
		if err != nil {
			panic("what the helly")

		}
		region := types.Region(strings.Split(line[3], "=")[1])

		result := e.HandleAction(action, player_id, skill_rating, region)

		if result.Success {
			if result.Player != nil {
				fmt.Printf("QUEUED player_id=%s, skill_rating=%d, region=%s\n", player_id, skill_rating, region)
			}
		}

		if action == "ENTER" {
			fmt.Println()
			e.PrintWaitingRoom()
		}

	}
}
