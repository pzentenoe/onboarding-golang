package games

import (
	"github.com/pzentenoe/onboarding-golang/app/domain/repository"
)

type gameUsecase struct {
	gameRepository repository.GameRepository
}

func NewGameUsecase(gameRepository repository.GameRepository) *gameUsecase {
	return &gameUsecase{gameRepository: gameRepository}
}


