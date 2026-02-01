package engine

import (
	"time"

	"github.com/rohan-sagar/game-server/internal/types"
)

type Player struct {
	Id             string
	SkillRating    int
	Region         types.Region
	QueueEntryTime time.Time
}

// create new player
func NewPlayer(id string, skillRating int, region types.Region) *Player {
	return &Player{
		Id:             id,
		SkillRating:    skillRating,
		Region:         region,
		QueueEntryTime: time.Now(),
	}
}
