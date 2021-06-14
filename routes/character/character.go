package character

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/moonlightfight/elo-backend/database"
	"github.com/moonlightfight/elo-backend/models"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateCharacterEndpoint(response http.ResponseWriter, request *http.Request) {
	client, err := database.ConfigDB()
	if err != nil {
		fmt.Println(err)
	}
	var newCharacter models.Character
	response.Header().Set("content-type", "application/json")
	_ = json.NewDecoder(request.Body).Decode(&newCharacter)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	if cancel != nil {
		fmt.Println(cancel)
	}
	collection := client.Database("test").Collection("Character")
	result, err := collection.InsertOne(ctx, newCharacter)
	if err != nil {
		fmt.Println(err)
	}
	json.NewEncoder(response).Encode(result)
}

func GetCharactersEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var characters []models.Character
	client, err := database.ConfigDB()
	if err != nil {
		fmt.Println(err)
	}
	collection := client.Database("test").Collection("Character")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var character models.Character
		cursor.Decode(&character)
		characters = append(characters, character)
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(characters)
}
