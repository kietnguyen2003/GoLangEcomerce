package repo

type PongRepo struct{}

func NewPongRepo() *PongRepo {
	return &PongRepo{}
}

func (pr *PongRepo) Pong() string {
	return "pong"
}
