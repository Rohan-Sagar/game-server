package engine

import (
	"fmt"
	"time"

	"github.com/rohan-sagar/game-server/internal/types"
)

type Player struct {
	Id             string
	SkillRating    int
	Region         types.Region
	QueueEntryTime time.Time
}

// create and validate new player
// checks if id is empty string and skill_rating in the correct age
func NewPlayer(id string, skillRating int, region types.Region) (*Player, error) {
	if id == "" {
		return nil, fmt.Errorf("playerId cannot be empty\n")
	}

	if skillRating < 1000 || skillRating > 3000 {
		return nil, fmt.Errorf("skillRating needs to be between 1000 and 3000")
	}

	return &Player{
		Id:             id,
		SkillRating:    skillRating,
		Region:         region,
		QueueEntryTime: time.Now(),
	}, nil
}
