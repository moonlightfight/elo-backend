package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	c "./config"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

type Admin struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	email    string             `json:"email,omitempty" bson:"email,omitempty"`
	username string             `json:"username,omitempty" bson:"username,omitempty"`
	password string             `json:"password,omitempty" bson:"password,omitempty"`
}

func CreateAdminEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var admin Admin
	_ = json.NewDecoder(request.Body).Decode(&admin)
	collection := client.Database("").Collection("Admin")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, _ := collection.InsertOne(ctx, admin)
	json.NewEncoder(response).Encode(result)
}

func main() {
	// Set the file name of the configurations file
	viper.SetConfigName("config")

	// Set the path to look for the configurations file
	viper.AddConfigPath(".")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	viper.SetConfigType("yml")
	var configuration c.Configurations

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}
	fmt.Println("Start writing code!")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb+srv://%s:%s@cluster0.ucnph.mongodb.net/%s?retryWrites=true&w=majority", configuration.Database.DBUser, configuration.Database.DBPass, configuration.Database.DBName))
	client, _ := mongo.Connect(ctx, clientOptions)
	router := mux.NewRouter()
	router.HandleFunc("/admin", CreateAdminEndpoint).Methods("POST")
	http.ListenAndServe(":6000", router)
}
