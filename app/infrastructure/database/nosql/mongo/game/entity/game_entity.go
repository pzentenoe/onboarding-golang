package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type GameEntity struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Title    string             `bson:"title"`
	Price    int                `bson:"price"`
	ImageURL string             `bson:"image_url"`
	Year     int                `bson:"year"`
}
