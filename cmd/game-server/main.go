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
		var result engine.ActionResult

		line := strings.Split(scanner.Text(), " ")
		action := line[0]

		switch action {
		case "ENTER":
			player_id := strings.Split(line[1], "=")[1]
			skill_rating, err := strconv.Atoi(strings.Split(line[2], "=")[1])

			if err != nil {
				panic("what the helly")

			}
			region := types.Region(strings.Split(line[3], "=")[1])

			result = e.HandleAction(action, player_id, skill_rating, region)
		case "TICK":
			result = e.HandleAction(action, "", 0, "us-east")
		default:
			panic("wrong code")
		}

		if result.Success {
			if result.Player != nil {
				fmt.Printf("QUEUED player_id=%s skill_rating=%d region=%s queue_join_time=%v\n", result.Player.Id, result.Player.SkillRating, result.Player.Region, result.Player.QueueEntryTime)
			}
			if result.Match != nil {
				fmt.Printf("MATCH match_id=%d players=%v average_skill=%f formationTime=%v CreatedAt=%v", result.Match.Id, result.Match.Players, result.Match.AverageSkill, result.Match.FormationTime, result.Match.CreatedAt)
			}
		}

		if action == "ENTER" {
			fmt.Println()
			e.PrintWaitingRoom()
		}

	}
}
