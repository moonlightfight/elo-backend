package database

import (
	"context"
	"log"
	"time"

	"github.com/moonlightfight/elo-backend/constants"
	"github.com/moonlightfight/elo-backend/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var databaseName = constants.GetEnvVar("DATABASE_NAME")

type DB struct {
	client *mongo.Client
}

func Connect(dbUrl string) *DB {
	client, err := mongo.NewClient(options.Client().ApplyURI(dbUrl))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	return &DB{
		client: client,
	}
}

func (db *DB) GetPlayers() []*model.Player {
	playerColl := db.client.Database(databaseName).Collection("Player")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cursor, err := playerColl.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	var players []*model.Player

	for cursor.Next(ctx) {
		var player *model.Player

		cursor.Decode(&player)

		players = append(players, player)
	}

	return players
}

func (db *DB) GetCharacters() []*model.Character {
	charactersColl := db.client.Database(databaseName).Collection("Character")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cursor, err := charactersColl.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	var characters []*model.Character

	for cursor.Next(ctx) {
		var character *model.Character

		cursor.Decode(&character)

		characters = append(characters, character)
	}

	return characters
}

func (db *DB) CreatePlayer(player model.Player) *model.Player {
	playerColl := db.client.Database(databaseName).Collection("Player")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	result, err := playerColl.InsertOne(ctx, bson.D{{Key: "username", Value: player.Username}, {Key: "slug", Value: player.Slug}, {Key: "rating", Value: player.Rating}, {Key: "score", Value: player.Score}})
	if err != nil {
		log.Println(err)
	}
	id := result.InsertedID.(primitive.ObjectID).Hex()
	newPlayer := model.Player{
		ID:       id,
		Username: player.Username,
		Slug:     player.Slug,
		Rating:   player.Rating,
		Score:    player.Score,
	}

	return &newPlayer
}
