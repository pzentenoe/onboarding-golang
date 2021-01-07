package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/pzentenoe/onboarding-golang/app/domain/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	jsonResponseGameListOK      = "[{\"id\":\"1\",\"title\":\"god of war\",\"price\":40000,\"image_url\":\"url\",\"year\":2018}]\n"
	jsonResponseGameListError   = "{\"message\":\"Fail to get games\"}\n"
	jsonResponseCreateGameOK    = "{\"id\":\"1\",\"title\":\"god of war\",\"price\":40000,\"image_url\":\"url\",\"year\":2018}\n"
	jsonResponseCreateGameError = "{\"message\":\"Error to create game\"}\n"
)

type gameUsecaseMock struct {
	mock.Mock
}

func (u *gameUsecaseMock) GetGames() ([]*models.Game, error) {
	args := u.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*models.Game), args.Error(1)

}
func (u *gameUsecaseMock) CreateGame(game *models.Game) (*models.Game, error) {
	args := u.Called(game)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Game), args.Error(1)
}

func Test_gameHandler_getGames(t *testing.T) {
	t.Parallel()
	t.Run("when get games OK return list of games", func(t *testing.T) {
		//Arrange
		mockGameUsecase := new(gameUsecaseMock)
		mockGameUsecase.On("GetGames").
			Return(getGamesStub(), nil)

		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/games", nil)
		resp := httptest.NewRecorder()
		echoContext := e.NewContext(req, resp)

		handler := NewGameHandler(e, mockGameUsecase)

		//Act
		handler.getGames(echoContext)

		//Assert
		assert.Equal(t, jsonResponseGameListOK, resp.Body.String())
		assert.Equal(t, http.StatusOK, resp.Code)
	})

	t.Run("when get games fails return an error message", func(t *testing.T) {

		mockGameUsecase := new(gameUsecaseMock)
		mockGameUsecase.On("GetGames").
			Return(nil, errors.New("Fail to get games"))

		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/games", nil)
		resp := httptest.NewRecorder()
		echoContext := e.NewContext(req, resp)

		handler := NewGameHandler(e, mockGameUsecase)

		//Act
		handler.getGames(echoContext)

		//Assert
		assert.Equal(t, jsonResponseGameListError, resp.Body.String())
		assert.Equal(t, http.StatusBadGateway, resp.Code)
	})

}

func Test_gameHandler_createGame(t *testing.T) {
	t.Parallel()
	t.Run("when create game works OK return new Game with ID", func(t *testing.T) {
		mockGameUsecase := new(gameUsecaseMock)
		gameResponse := getGameStub()
		mockGameUsecase.On("CreateGame", mock.AnythingOfType("*models.Game")).
			Return(gameResponse, nil)

		gameRequest := getGameStub()
		requestData, _ := json.Marshal(gameRequest)

		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/games", bytes.NewBuffer(requestData))
		req.Header.Add(echo.HeaderContentType, echo.MIMEApplicationJSON)
		response := httptest.NewRecorder()
		echoContext := e.NewContext(req, response)

		handler := NewGameHandler(e, mockGameUsecase)

		handler.createGame(echoContext)

		assert.Equal(t, jsonResponseCreateGameOK, response.Body.String())
		assert.Equal(t, http.StatusCreated, response.Code)
	})

	t.Run("when create game fail return and error response", func(t *testing.T) {
		mockGameUsecase := new(gameUsecaseMock)
		mockGameUsecase.On("CreateGame", mock.AnythingOfType("*models.Game")).
			Return(nil, errors.New("Error to create game"))

		gameRequest := getGameStub()
		requestData, _ := json.Marshal(gameRequest)

		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/games", bytes.NewBuffer(requestData))
		req.Header.Add(echo.HeaderContentType, echo.MIMEApplicationJSON)
		response := httptest.NewRecorder()
		echoContext := e.NewContext(req, response)

		handler := NewGameHandler(e, mockGameUsecase)

		handler.createGame(echoContext)

		assert.Equal(t, jsonResponseCreateGameError, response.Body.String())
		assert.Equal(t, http.StatusBadGateway, response.Code)
	})

	t.Run("when create game fail for invalid body", func(t *testing.T) {
		mockGameUsecase := new(gameUsecaseMock)
		gameResponse := getGameStub()
		mockGameUsecase.On("CreateGame", mock.AnythingOfType("*models.Game")).
			Return(gameResponse, nil)

		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/games", bytes.NewBufferString("<name>asas<name>"))
		req.Header.Add(echo.HeaderContentType, echo.MIMEApplicationJSON)
		response := httptest.NewRecorder()
		echoContext := e.NewContext(req, response)

		handler := NewGameHandler(e, mockGameUsecase)

		handler.createGame(echoContext)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})
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
