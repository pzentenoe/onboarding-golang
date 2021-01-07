package games

import (
	"github.com/pzentenoe/onboarding-golang/app/domain/models"
	"github.com/stretchr/testify/mock"
)

type gameRepositoryMock struct {
	mock.Mock
}

func (m *gameRepositoryMock) FindGames() ([]*models.Game, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*models.Game), args.Error(1)
}
func (m *gameRepositoryMock) CreateGame(game *models.Game) (*models.Game, error) {
	args := m.Called(game)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Game), args.Error(1)
}


func getGamesStub() []*models.Game {
	games := make([]*models.Game, 0)
	games = append(games, getGameStub())
	return games
}

func getGameStub() *models.Game {
	return &models.Game{
		ID:       "1",
		Title:    "god of war",
		Price:    40000,
		ImageURL: "url",
		Year:     2018,
	}
}


