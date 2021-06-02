package tournament

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
	c "github.com/moonlightfight/elo-backend/config"
	"github.com/moonlightfight/elo-backend/routes/tournament/types"
	"github.com/spf13/viper"
)

func getChallongeBracket(tournamentId string, subDomain interface{}, apiKey string) types.BracketInfo {
	var apiUrl string
	/// define the structs in order to morph the data into universal data
	var bracketInfo types.BracketInfo
	var matches []types.Match
	var players []types.Player
	if subDomain == nil {
		// if there's no subdomain, we only need to pack the tournament ID and api key into the query params
		apiUrl = fmt.Sprintf("https://api.challonge.com/v1/tournaments/%s.json?api_key=%s&include_participants=1&include_matches=1", tournamentId, apiKey)
	} else {
		// if there's a subdomain, we need to concatenate the subdomain with the tournament ID and also include the api key
		apiUrl = fmt.Sprintf("https://api.challonge.com/v1/tournaments/%s-%s.json?api_key=%s&include_participants=1&include_matches=1", subDomain, tournamentId, apiKey)
	}
	// run the api request
	resp, err := http.Get(apiUrl)
	if err != nil {
		log.Fatalln(err)
	}

	// unpack the json and unload it into the bracket struct
	defer resp.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	var challongeBracket types.ChallongeBracket

	json.Unmarshal(bodyBytes, challongeBracket)

	// morph players into frontend format
	for _, participant := range challongeBracket.Tournament.Participants {
		player := types.Player{
			ID:    participant.Participant.ID,
			Name:  participant.Participant.DisplayName,
			Place: participant.Participant.FinalRank,
		}
		players = append(players, player)
	}

	// generate the match info

	// generate the full bracket info
	bracketInfo = types.BracketInfo{
		Title:          challongeBracket.Tournament.Name,
		NumPlayers:     challongeBracket.Tournament.ParticipantsCount,
		TournamentDate: challongeBracket.Tournament.StartedAt,
		Players:        players,
		Matches:        matches,
	}

	// return the data to the API endpoint
	return bracketInfo
}

func getSmashBracket(slug, apiKey string) types.BracketInfo {
	// set an http client since we need to pack request headers
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	// define the structs in order to morph the data into universal data
	var bracketInfo types.BracketInfo
	var matches []types.Match
	var players []types.Player
	// set the endpoint
	apiUrl := "https://api.smash.gg/gql/alpha"
	// generate the authorization header value
	authHeader := fmt.Sprintf("Bearer %s", apiKey)
	// create the GQL query and variables to pass
	var query types.SmashQuery
	var variables types.SmashVariables
	variables = types.SmashVariables{
		Slug: slug,
	}
	query = types.SmashQuery{
		Query:     "query EventQuery($slug: String!) { event(slug: $slug) { id name startAt standings(query: {page: 1, perPage: 500}) { nodes { id placement entrant { id name } } } sets { nodes { id slots { entrant { id name } } winnerId displayScore } } videogame { id name } tournament { id name } } }",
		Variables: variables,
	}
	// create the json
	jsonBody, _ := json.Marshal(query)
	// generate the api request (POST is recommended on GraphQL queries/mutations from REST)
	req, err := http.NewRequest("POST", apiUrl, bytes.NewReader(jsonBody))
	if err != nil {
		panic("error formatting json!")
	}
	// set the headers on the api request
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", authHeader)
	// execute the API request
	resp, err := client.Do(req)
	if err != nil {
		panic("POST error")
	}
	// read the json data and unpack it into the bracket struct
	defer resp.Body.Close()
	var smashBracket types.SmashBracket
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(bodyBytes, smashBracket)

	// generate the player info
	for _, player := range smashBracket.Data.Event.Standings.Nodes {
		insertedPlayer := types.Player{
			ID:    player.Entrant.ID,
			Name:  player.Entrant.Name,
			Place: player.Placement,
		}
		players = append(players, insertedPlayer)
	}

	// generate the match info

	// generate the fully formatted bracket
	bracketInfo = types.BracketInfo{
		Title:          smashBracket.Data.Event.Tournament.Name,
		NumPlayers:     len(smashBracket.Data.Event.Standings.Nodes),
		TournamentDate: time.Unix(smashBracket.Data.Event.StartAt, 0),
		Players:        players,
		Matches:        matches,
	}

	// return the bracket info to endpoint
	return bracketInfo
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
		// trim the url down to the obscenely long event slug that IDK what Smash.gg was thinking when they created it
		re := strings.NewReplacer("https://smash.gg/", "", "/overview", "")
		slug := re.Replace(url)
		getSmashBracket(slug, configuration.ApiKeys.Smash)
	} else {
		panic("unsupported bracket URL")
	}
}