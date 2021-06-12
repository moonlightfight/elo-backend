package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	c "github.com/moonlightfight/elo-backend/config"
	"github.com/moonlightfight/elo-backend/routes/admin"
	"github.com/moonlightfight/elo-backend/routes/character"
	"github.com/moonlightfight/elo-backend/routes/tournament"
	"github.com/spf13/viper"
)

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
	port := fmt.Sprintf(":%d", configuration.Server.Port)
	router := mux.NewRouter()

	// Admin routes
	router.HandleFunc("/api/admin", admin.CreateAdminEndpoint).Methods("POST")
	router.HandleFunc("/api/admin/login", admin.AdminLoginEndpoint).Methods("POST")

	// Tournament routes
	router.HandleFunc("/api/tournament/getfromweb", tournament.GetTournamentData).Queries("url", "{url}").Methods("GET")
	router.HandleFunc("/api/tournaments", tournament.CreateTournament).Methods("POST")

	// Character routes
	router.HandleFunc("/api/character", character.CreateCharacterEndpoint).Methods("POST")
	router.HandleFunc("/api/character", character.GetCharactersEndpoint).Methods("GET")

	// run the server
	fmt.Printf("server listening on http://localhost%v", port)
	http.ListenAndServe(port, router)
}
