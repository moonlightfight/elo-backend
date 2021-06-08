package character

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/moonlightfight/elo-backend/database"
	"github.com/moonlightfight/elo-backend/models"
)

func CreateCharacterEndpoint(response http.ResponseWriter, request *http.Request) {
	client, err := database.ConfigDB()
	if err != nil {
		fmt.Println(err)
	}
	var newCharacter models.Admin
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
