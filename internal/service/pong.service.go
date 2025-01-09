package service

import "src/internal/repo"

type PongService struct {
	// truyền vô thằng repo
	pongRepo *repo.PongRepo
}

func NewPongService() *PongService {
	return &PongService{
		pongRepo: repo.NewPongRepo(),
	}
}

func (ps *PongService) Pong() string {
	return ps.pongRepo.Pong()
}
