package repository

import "github.com/pzentenoe/onboarding-golang/app/domain/models"

type GameRepository interface {
	FindGames() ([]*models.Game, error)
	CreateGame(game *models.Game) (*models.Game, error)
}
