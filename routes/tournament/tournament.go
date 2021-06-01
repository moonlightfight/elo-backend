package tournament

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	c "github.com/moonlightfight/elo-backend/config"
	"github.com/moonlightfight/elo-backend/routes/tournament/types"
	"github.com/spf13/viper"
)

func getChallongeBracket(tournamentId string, subDomain interface{}, apiKey string) types.BracketInfo {
	var apiUrl string
	var bracketInfo types.BracketInfo
	var matches []types.Match
	var players []types.Player
	if subDomain == nil {
		apiUrl = fmt.Sprintf("https://api.challonge.com/v1/tournaments/%s.json?api_key=%s&include_participants=1&include_matches=1", tournamentId, apiKey)
	} else {
		apiUrl = fmt.Sprintf("https://api.challonge.com/v1/tournaments/%s-%s.json?api_key=%s&include_participants=1&include_matches=1", subDomain, tournamentId, apiKey)
	}
	resp, err := http.Get(apiUrl)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	var challongeBracket types.ChallongeBracket

	json.Unmarshal(bodyBytes, challongeBracket)
}

func getSmashBracket(slug, apiKey string) types.BracketInfo {
	var bracketInfo types.BracketInfo
	var matches []types.Match
	var players []types.Player
	apiUrl := "https://api.smash.gg/gql/alpha"
	var query types.SmashQuery
	var variables types.SmashVariables
	variables = types.SmashVariables{
		Slug: slug,
	}
	query = types.SmashQuery{
		Query:     "query EventQuery($slug: String!) { event(slug: $slug) { id name standings(query: {page: 1, perPage: 500}) { nodes { id placement entrant { id name } } } sets { nodes { id slots { entrant { id name } } winnerId displayScore } } videogame { id name } tournament { id name } } }",
		Variables: variables,
	}
}

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

	var bracket types.BracketInfo

	// check the bracket url, if it is valid, send it to the proper function for retrieval and formatting, else, throw an error
	if strings.Contains(url, "challonge") {
		var tournamentId string
		var subDomain interface{}
		if strings.Contains(url, "https://challonge.com/") {
			subDomain = nil
			tournamentId = strings.Replace(url, "https://challonge.com/", "", -1)
		} else {
			trunc := strings.Replace(url, "https://", "", 1)
			subDomain = strings.TrimRight(trunc, ".challonge.com")
			tournamentId = strings.TrimLeft(url, fmt.Sprintf("https://%s.challonge.com/", subDomain))
		}
		bracket = getChallongeBracket(tournamentId, subDomain, configuration.ApiKeys.Challonge)
	} else if strings.Contains(url, "smash") {
		slug := strings.Replace(url, "https://smash.gg/", "", -1)
		getSmashBracket(slug, configuration.ApiKeys.Smash)
	}
}
