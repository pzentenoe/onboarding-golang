package usecase

import "github.com/pzentenoe/onboarding-golang/app/domain/models"

type GameUsecase interface {
	GetGames() ([]*models.Game, error)
	CreateGame(game *models.Game) (*models.Game, error)
}
