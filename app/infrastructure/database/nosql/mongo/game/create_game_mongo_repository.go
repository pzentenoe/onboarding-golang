package game

import (
	"context"
	"github.com/labstack/gommon/log"
	"github.com/pzentenoe/onboarding-golang/app/domain/models"
	"github.com/pzentenoe/onboarding-golang/app/infrastructure/database/nosql/mongo/game/mapper"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *gameMongoRepository) CreateGame(game *models.Game) (*models.Game, error) {
	collection := r.getGameCollection()
	gameEntity := mapper.GameModelToGameEntity(game)
	result, err := collection.InsertOne(context.Background(), gameEntity)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	objectId := result.InsertedID.(primitive.ObjectID)

	game.ID = objectId.Hex()
	return game, nil

}
