package player

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/moonlightfight/elo-backend/database"
	"github.com/moonlightfight/elo-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreatePlayerEndpoint(response http.ResponseWriter, request *http.Request) {
	type PlayerData struct {
		Name string `json:"name"`
	}

	client, err := database.ConfigDB()
	if err != nil {
		log.Println(err)
	}

	response.Header().Set("content-type", "application/json")

	var data PlayerData

	jsonErr := json.NewDecoder(request.Body).Decode(&data)
	if jsonErr != nil {
		log.Println(err)
	}

	lowerName := strings.ToLower(data.Name)

	slug := strings.Replace(lowerName, " ", "-", -1)

	player := models.Player{
		Username:      data.Name,
		Country:       "",
		Points:        0,
		Ranking:       1200,
		Slug:          slug,
		RealName:      "",
		Controller:    "",
		Twitter:       "",
		Twitch:        "",
		Picture:       "",
		Tournaments:   []primitive.ObjectID{},
		Matches:       []primitive.ObjectID{},
		SubCharacters: []primitive.ObjectID{},
	}

	collection := client.Database("test").Collection("Player")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	if cancel != nil {
		log.Println(cancel)
	}

	result, err := collection.InsertOne(ctx, player)
	if err != nil {
		log.Println(err)
	}

	json.NewEncoder(response).Encode(result)
}

func GetPlayersEndPoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var players []models.Player
	client, err := database.ConfigDB()
	if err != nil {
		fmt.Println(err)
	}
	collection := client.Database("test").Collection("Player")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var player models.Player
		cursor.Decode(&player)
		players = append(players, player)
	}
	if len(players) == 0 {
		players = []models.Player{}
	}
	type ReturnData struct {
		Players []models.Player `json:"players"`
	}
	returnData := ReturnData{
		Players: players,
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(returnData)
}
