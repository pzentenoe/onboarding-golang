package main

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pzentenoe/onboarding-golang/app/application/usecase/games"
	gameRepo "github.com/pzentenoe/onboarding-golang/app/infrastructure/database/nosql/mongo/game"
	"github.com/pzentenoe/onboarding-golang/app/infrastructure/http/handlers"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func main() {

	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.CORS())

	client := getMongoClient()

	initDependencies(client, e)
	e.Logger.Fatal(e.Start(":8088"))

}

func initDependencies(client *mongo.Client, e *echo.Echo) {
	gameMongoRepository := gameRepo.NewGameMongoRepository(client)
	gameUsecase := games.NewGameUsecase(gameMongoRepository)
	_ = handlers.NewGameHandler(e, gameUsecase)

	handlers.NewHealthHandler(e)
}

func getMongoClient() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://chat_user:WbnBnWLRtCOZlXVT@pzenteno-cluster.q4adt.mongodb.net/games_store?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatalf("error to create client mongodb :%s", err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("error to connect to mongodb :%s", err.Error())
	}
	return client
}
