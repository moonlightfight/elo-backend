package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	c "github.com/moonlightfight/elo-backend/config"
	a "github.com/moonlightfight/elo-backend/routes/admin"
	to "github.com/moonlightfight/elo-backend/routes/tournament"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

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
	port := fmt.Sprintf(":%d", configuration.Server.Port)
	mongo.Connect(ctx, clientOptions)
	router := mux.NewRouter()
	router.HandleFunc("/api/admin", a.CreateAdminEndpoint).Methods("POST")
	router.HandleFunc("/api/admin/login", a.AdminLoginEndpoint).Methods("POST")
	router.HandleFunc("/api/tournament/getfromweb", to.GetTournamentData).Methods("GET")
	http.ListenAndServe(port, router)
}
