package main

import (
	"flag"
	"fmt"

	"github.com/rohan-sagar/game-server/internal/engine"
	"github.com/rohan-sagar/game-server/internal/types"
)

func main() {
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

	// new_engine := engine.NewEngine()
	player := engine.NewPlayer(*player_id, *skill_rating, region)
	fmt.Printf("action: %s\n", *action)
	fmt.Printf("%v\n", player)

	waiting_room := make(map[string]*engine.Player)
	waiting_room[player.Id] = player

	fmt.Printf("Waiting room:\n")
	for id, player := range waiting_room {
		fmt.Printf("  %s: %+v\n", id, *player)
	}

}
