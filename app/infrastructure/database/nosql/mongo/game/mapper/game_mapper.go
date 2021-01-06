package mapper

import (
	"github.com/pzentenoe/onboarding-golang/app/domain/models"
	"github.com/pzentenoe/onboarding-golang/app/infrastructure/database/nosql/mongo/game/entity"
)

func GameEntityToGameModel(gameEntity *entity.GameEntity) *models.Game {
	game := &models.Game{
		ID:       gameEntity.ID.Hex(),
		Title:    gameEntity.Title,
		Price:    gameEntity.Price,
		ImageURL: gameEntity.ImageURL,
		Year:     gameEntity.Year,
	}
	return game
}

func GameModelToGameEntity(gameModel *models.Game) *entity.GameEntity {
	game := &entity.GameEntity{
		Title:    gameModel.Title,
		Price:    gameModel.Price,
		ImageURL: gameModel.ImageURL,
		Year:     gameModel.Year,
	}
	return game
}
