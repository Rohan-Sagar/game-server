package engine

import (
	"math/rand"
	"time"
)

type Match struct {
	Id           int
	Players      []Player
	AverageSkill float64
	// QualityScore  float64
	FormationTime time.Time
	CreatedAt     time.Time
	// Region        types.Region
}

func NewMatch(players []Player) *Match {
	id := rand.Intn(100000)
	var total int
	for _, player := range players {
		total += player.SkillRating
	}
	avg_skill := float64(total / len(players))

	return &Match{
		Id:            id,
		Players:       players,
		AverageSkill:  avg_skill,
		FormationTime: time.Now(),
		CreatedAt:     time.Now(),
	}

}
