package engine

type Engine struct {
	WaitingRoom map[string]*Player
}

func NewEngine() *Engine {
	return &Engine{
		WaitingRoom: make(map[string]*Player),
	}
}
