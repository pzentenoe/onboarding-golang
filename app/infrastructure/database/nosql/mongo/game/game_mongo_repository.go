package game

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type gameMongoRepository struct {
	mongoClient *mongo.Client
}

func NewGameMongoRepository(mongoClient *mongo.Client) *gameMongoRepository {
	return &gameMongoRepository{mongoClient: mongoClient}
}

func (r *gameMongoRepository) getGameCollection() *mongo.Collection {
	collection := r.mongoClient.Database("games_store").Collection("games")
	return collection
}
