package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/pzentenoe/onboarding-golang/app/application/usecase"
	"github.com/pzentenoe/onboarding-golang/app/domain/models"
	"net/http"
)

type gameHandler struct {
	gameUsecase usecase.GameUsecase
}

func NewGameHandler(e *echo.Echo, gameUsecase usecase.GameUsecase) *gameHandler {
	h := &gameHandler{gameUsecase: gameUsecase}

	g := e.Group("/games")
	g.GET("", h.getGames)
	g.POST("", h.createGame)
	g.GET("/:id", h.getGameById)

	return h

}

func (h *gameHandler) getGames(c echo.Context) error {

	gameList, err := h.gameUsecase.GetGames()
	if err != nil {
		return c.JSON(http.StatusBadGateway, echo.Map{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, gameList)
}

func (h *gameHandler) getGameById(c echo.Context) error {
	return nil
}

func (h *gameHandler) createGame(c echo.Context) error {

	var game *models.Game
	if err := c.Bind(&game); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}

	gameCreated, err := h.gameUsecase.CreateGame(game)
	if err != nil {
		return c.JSON(http.StatusBadGateway, echo.Map{"message": err.Error()})
	}

	return c.JSON(http.StatusCreated, gameCreated)
}
