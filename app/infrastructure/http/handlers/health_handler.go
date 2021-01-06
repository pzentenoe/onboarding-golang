package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type healthHandler struct {
}

type healthResponse struct {
	Status string `json:"status" xml:"status"`
}

func NewHealthHandler(e *echo.Echo) *healthHandler {
	h := new(healthHandler)
	e.GET("/health", h.healthCheck)

	return h
}

func (h *healthHandler) healthCheck(c echo.Context) error {
	response := healthResponse{
		Status: "OK",
	}

	return c.JSON(http.StatusOK, response)
}
