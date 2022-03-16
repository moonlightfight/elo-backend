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
	result, err := playerColl.InsertOne(ctx, player)
	if err != nil {
		log.Println(err)
	}
	id := result.InsertedID.(primitive.ObjectID).Hex()
	player.ID = id

	return &player
}

func (db *DB) InsertMatch(match model.Match) *model.Match {
	matchColl := db.client.Database(databaseName).Collection("Match")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	result, err := matchColl.InsertOne(ctx, match)
	if err != nil {
		log.Println(err)
	}
	id := result.InsertedID.(primitive.ObjectID).Hex()
	match.ID = id
	return &match
}

func (db *DB) InsertTournament(tournament model.Tournament) *model.Tournament {
	tournamentColl := db.client.Database(databaseName).Collection("Tournament")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	result, err := tournamentColl.InsertOne(ctx, tournament)
	if err != nil {
		log.Println(err)
	}
	id := result.InsertedID.(primitive.ObjectID).Hex()
	tournament.ID = id
	return &tournament
}

func (db *DB) GetPlayerByID(playerID string) *model.Player {
	objectID, err := primitive.ObjectIDFromHex(playerID)
	if err != nil {
		log.Fatal(err)
	}
	playerColl := db.client.Database(databaseName).Collection("Player")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	result := playerColl.FindOne(ctx, bson.M{"_id": objectID})
	var player model.Player
	result.Decode(player)
	return &player
}

func (db *DB) GetCharacter(characterId string) *model.Character {
	objectId, err := primitive.ObjectIDFromHex(characterId)
	if err != nil {
		log.Fatal(err)
	}
	characterColl := db.client.Database(databaseName).Collection("Character")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	result := characterColl.FindOne(ctx, bson.M{"_id": objectId})
	var character model.Character
	result.Decode(character)
	return &character
}
