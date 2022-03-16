package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/moonlightfight/elo-backend/constants"
	"github.com/moonlightfight/elo-backend/graph/model"
	"github.com/moonlightfight/elo-backend/types"
)

func GetChallongeBracket(tournamentId string, subDomain interface{}) *model.APIReturnedTournament {
	apiKey := constants.GetEnvVar("CHALLONGE_API_KEY")
	var apiUrl string
	/// define the structs in order to morph the data into universal data
	var bracketInfo model.APIReturnedTournament
	var matches []*model.APIReturnedMatch
	var players []*model.APIReturnedPlayer
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

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	var challongeBracket types.ChallongeBracket

	json.Unmarshal(bodyBytes, &challongeBracket)

	// morph players into frontend format
	for _, participant := range challongeBracket.Tournament.Participants {
		player := model.APIReturnedPlayer{
			ID:    participant.Participant.ID,
			Name:  participant.Participant.DisplayName,
			Place: participant.Participant.FinalRank,
		}
		players = append(players, &player)
	}

	// generate the match info
	for _, match := range challongeBracket.Tournament.Matches {
		var winnerIndex int
		var loserIndex int
		for i, player := range players {
			if player.ID == match.Match.WinnerID {
				winnerIndex = i
				break
			}
		}
		for i, player := range players {
			if player.ID == match.Match.LoserID {
				loserIndex = i
				break
			}
		}
		player1Score, _ := strconv.Atoi(strings.TrimRight(match.Match.ScoresCsv, "-"))
		player2Score, _ := strconv.Atoi(strings.TrimLeft(match.Match.ScoresCsv, "-"))
		var winnerScore int
		var loserScore int
		if player1Score > player2Score {
			winnerScore = player1Score
			loserScore = player2Score
		} else {
			winnerScore = player2Score
			loserScore = player1Score
		}
		set := model.APIReturnedMatch{
			WinnerID:    match.Match.WinnerID,
			LoserID:     match.Match.LoserID,
			WinnerName:  players[winnerIndex].Name,
			LoserName:   players[loserIndex].Name,
			WinnerScore: winnerScore,
			LoserScore:  loserScore,
			MatchDate:   match.Match.StartedAt,
		}
		matches = append(matches, &set)
	}

	// generate the full bracket info
	bracketInfo = model.APIReturnedTournament{
		Title:          challongeBracket.Tournament.Name,
		NumPlayers:     challongeBracket.Tournament.ParticipantsCount,
		TournamentDate: challongeBracket.Tournament.StartedAt,
		Players:        players,
		Matches:        matches,
	}

	// return the data to the API endpoint
	return &bracketInfo
}

func GetSmashBracket(slug string) *model.APIReturnedTournament {
	apiKey := constants.GetEnvVar("SMASH_API_KEY")
	// set an http client since we need to pack request headers
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	// define the structs in order to morph the data into universal data
	var bracketInfo model.APIReturnedTournament
	var matches []*model.APIReturnedMatch
	var players []*model.APIReturnedPlayer
	// set the endpoint
	apiUrl := "https://api.smash.gg/gql/alpha"
	// generate the authorization header value
	authHeader := fmt.Sprintf("Bearer %s", apiKey)
	// create the GQL query and variables to pass

	variables := types.SmashVariables{
		Slug: slug,
	}
	query := types.SmashQuery{
		Query:     "query EventQuery($slug: String!) { event(slug: $slug) { id name startAt standings(query: {page: 1, perPage: 500}) { nodes { id placement entrant { id name } } } sets { nodes { id slots { entrant { id name } } winnerId displayScore completedAt } } videogame { id name } tournament { id name } } }",
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
	json.Unmarshal(bodyBytes, &smashBracket)

	// generate the player info
	for _, player := range smashBracket.Data.Event.Standings.Nodes {
		insertedPlayer := model.APIReturnedPlayer{
			ID:    player.Entrant.ID,
			Name:  player.Entrant.Name,
			Place: player.Placement,
		}
		players = append(players, &insertedPlayer)
	}

	// generate the match info
	for _, set := range smashBracket.Data.Event.Sets.Nodes {
		setPlayers := strings.Split(set.Displayscore, " - ")
		leftPlayerScore, _ := strconv.Atoi(strings.TrimLeft(setPlayers[0], " "))
		rightPlayerScore, _ := strconv.Atoi(strings.TrimLeft(setPlayers[1], " "))
		var winnerScore int
		var loserScore int
		var winnerName string
		var loserName string
		var winnerId int
		var loserId int
		if leftPlayerScore > rightPlayerScore {
			winnerScore = leftPlayerScore
			loserScore = rightPlayerScore
			winnerName = set.Slots[0].Entrant.Name
			loserName = set.Slots[1].Entrant.Name
			winnerId = set.Slots[0].Entrant.ID
			loserId = set.Slots[1].Entrant.ID
		} else {
			winnerScore = rightPlayerScore
			loserScore = leftPlayerScore
			winnerName = set.Slots[1].Entrant.Name
			loserName = set.Slots[0].Entrant.Name
			winnerId = set.Slots[1].Entrant.ID
			loserId = set.Slots[0].Entrant.ID
		}
		match := model.APIReturnedMatch{
			WinnerID:    winnerId,
			LoserID:     loserId,
			WinnerName:  winnerName,
			LoserName:   loserName,
			WinnerScore: winnerScore,
			LoserScore:  loserScore,
			MatchDate:   time.Unix(set.CompletedAt, 0),
		}
		matches = append(matches, &match)
	}

	// generate the fully formatted bracket
	bracketInfo = model.APIReturnedTournament{
		Title:          smashBracket.Data.Event.Tournament.Name,
		NumPlayers:     len(smashBracket.Data.Event.Standings.Nodes),
		TournamentDate: time.Unix(smashBracket.Data.Event.StartAt, 0),
		Players:        players,
		Matches:        matches,
	}

	// return the bracket info to endpoint
	return &bracketInfo
}
