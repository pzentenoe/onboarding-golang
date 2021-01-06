package games

import "github.com/pzentenoe/onboarding-golang/app/domain/models"

func (u *gameUsecase) GetGames() ([]*models.Game, error) {
	return u.gameRepository.FindGames()
}
