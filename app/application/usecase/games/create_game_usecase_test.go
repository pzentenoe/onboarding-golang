package games

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func Test_gameUsecase_CreateGame(t *testing.T) {
	t.Parallel()
	t.Run("when create game OK return game with id", func(t *testing.T) {

		game := getGameStub()
		mockGameRepository := new(gameRepositoryMock)
		mockGameRepository.On("CreateGame", mock.AnythingOfType("*models.Game")).
			Return(game, nil)

		gameUsecase := NewGameUsecase(mockGameRepository)
		gameCreated, err := gameUsecase.CreateGame(game)

		assert.NotNil(t, gameCreated)
		assert.NoError(t, err)
		assert.Equal(t, game.ID, gameCreated.ID)

	})

	t.Run("when create game fail return an error", func(t *testing.T) {

		game := getGameStub()
		mockGameRepository := new(gameRepositoryMock)
		mockGameRepository.On("CreateGame", mock.AnythingOfType("*models.Game")).
			Return(nil, errors.New("error to create game"))

		gameUsecase := NewGameUsecase(mockGameRepository)
		gameCreated, err := gameUsecase.CreateGame(game)

		assert.Nil(t, gameCreated)
		assert.Error(t, err)

	})
}
