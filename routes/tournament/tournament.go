package tournament

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	c "github.com/moonlightfight/elo-backend/config"
	"github.com/moonlightfight/elo-backend/routes/tournament/types"
	"github.com/spf13/viper"
)

func getChallongeBracket(url, apiKey string) types.ChallongeBracket {
	fmt.Println("test")
}

func getSmashBracket(url, apiKey string) {}

func GetTournamentData(response http.ResponseWriter, request *http.Request) {
	// Set the file name of the configurations file
	viper.SetConfigName("config")

	// Set the path to look for the configurations file
	viper.AddConfigPath("../..")

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
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	url, _ := params["url"]

	// check the bracket url, if it is valid, send it to the proper function for retrieval and formatting, else, throw an error
	if strings.Contains(url, "challonge") {
		bracket := getChallongeBracket(url, configuration.ApiKeys.Challonge)
	} else if strings.Contains(url, "smash") {
		getSmashBracket(url, configuration.ApiKeys.Smash)
	}
}
