package games

import "github.com/pzentenoe/onboarding-golang/app/domain/models"

func (u *gameUsecase) CreateGame(game *models.Game) (*models.Game, error) {
	return u.gameRepository.CreateGame(game)
}
