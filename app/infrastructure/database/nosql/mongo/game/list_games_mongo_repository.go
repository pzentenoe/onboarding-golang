package game

import (
	"context"
	"github.com/labstack/gommon/log"
	"github.com/pzentenoe/onboarding-golang/app/domain/models"
	"github.com/pzentenoe/onboarding-golang/app/infrastructure/database/nosql/mongo/game/entity"
	"github.com/pzentenoe/onboarding-golang/app/infrastructure/database/nosql/mongo/game/mapper"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *gameMongoRepository) FindGames() ([]*models.Game, error) {

	collection := r.getGameCollection()

	cursor, err := collection.Find(context.Background(), primitive.D{})
	if err != nil {

	}
	defer cursor.Close(context.Background())

	games := make([]*models.Game, 0)

	for cursor.Next(context.Background()) {
		gameEntity := &entity.GameEntity{}
		err := cursor.Decode(gameEntity)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		game := mapper.GameEntityToGameModel(gameEntity)
		games = append(games, game)
	}

	return games, nil

}
