package player

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/moonlightfight/elo-backend/database"
	"github.com/moonlightfight/elo-backend/models"
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

}
