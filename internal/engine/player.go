package engine

import (
	"time"

	"github.com/rohan-sagar/game-server/internal/types"
)

type Player struct {
	Id          string
	SkillRating int
	Region      types.Region
	StartTime   int64
}

func NewPlayer(id string, skillRating int, region types.Region) *Player {
	return &Player{
		Id:          id,
		SkillRating: skillRating,
		Region:      region,
		StartTime:   time.Now().Unix(),
	}
}
