package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
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

func CalculateTournamentPoints(numPlayers, placing int) int {
	switch {
	case numPlayers < 5:
		switch placing {
		case 1:
			return 6
		case 2:
			return 3
		case 3:
			return 2
		default:
			return 1
		}
	case numPlayers > 4 && numPlayers < 7:
		switch placing {
		case 1:
			return 9
		case 2:
			return 6
		case 3:
			return 3
		case 4:
			return 2
		default:
			return 1
		}
	case numPlayers > 6 && numPlayers < 9:
		switch placing {
		case 1:
			return 12
		case 2:
			return 9
		case 3:
			return 6
		case 4:
			return 3
		case 5:
			return 2
		default:
			return 1
		}
	case numPlayers > 8 && numPlayers < 13:
		switch placing {
		case 1:
			return 15
		case 2:
			return 12
		case 3:
			return 9
		case 4:
			return 6
		case 5:
			return 3
		case 7:
			return 2
		default:
			return 1
		}
	case numPlayers > 12 && numPlayers < 17:
		switch placing {
		case 1:
			return 20
		case 2:
			return 15
		case 3:
			return 12
		case 4:
			return 9
		case 5:
			return 6
		case 7:
			return 3
		case 9:
			return 2
		default:
			return 1
		}
	case numPlayers > 16 && numPlayers < 25:
		switch placing {
		case 1:
			return 30
		case 2:
			return 20
		case 3:
			return 15
		case 4:
			return 12
		case 5:
			return 9
		case 7:
			return 6
		case 9:
			return 3
		case 13:
			return 2
		default:
			return 1
		}
	case numPlayers > 24 && numPlayers < 33:
		switch placing {
		case 1:
			return 40
		case 2:
			return 30
		case 3:
			return 20
		case 4:
			return 15
		case 5:
			return 12
		case 7:
			return 9
		case 9:
			return 6
		case 13:
			return 3
		case 17:
			return 2
		default:
			return 1
		}
	case numPlayers > 32 && numPlayers < 49:
		switch placing {
		case 1:
			return 60
		case 2:
			return 40
		case 3:
			return 30
		case 4:
			return 20
		case 5:
			return 15
		case 7:
			return 12
		case 9:
			return 9
		case 13:
			return 6
		case 17:
			return 3
		case 25:
			return 2
		default:
			return 1
		}
	case numPlayers > 48 && numPlayers < 65:
		switch placing {
		case 1:
			return 80
		case 2:
			return 60
		case 3:
			return 40
		case 4:
			return 30
		case 5:
			return 20
		case 7:
			return 15
		case 9:
			return 12
		case 13:
			return 9
		case 17:
			return 6
		case 25:
			return 3
		case 33:
			return 2
		default:
			return 1
		}
	case numPlayers > 48 && numPlayers < 65:
		switch placing {
		case 1:
			return 80
		case 2:
			return 60
		case 3:
			return 40
		case 4:
			return 30
		case 5:
			return 20
		case 7:
			return 15
		case 9:
			return 12
		case 13:
			return 9
		case 17:
			return 6
		case 25:
			return 3
		case 33:
			return 2
		default:
			return 1
		}
	case numPlayers > 64 && numPlayers < 97:
		switch placing {
		case 1:
			return 90
		case 2:
			return 70
		case 3:
			return 60
		case 4:
			return 40
		case 5:
			return 30
		case 7:
			return 20
		case 9:
			return 15
		case 13:
			return 12
		case 17:
			return 9
		case 25:
			return 6
		case 33:
			return 3
		case 49:
			return 2
		default:
			return 1
		}
	case numPlayers > 96 && numPlayers < 129:
		switch placing {
		case 1:
			return 120
		case 2:
			return 90
		case 3:
			return 70
		case 4:
			return 60
		case 5:
			return 40
		case 7:
			return 30
		case 9:
			return 20
		case 13:
			return 15
		case 17:
			return 12
		case 25:
			return 9
		case 33:
			return 6
		case 49:
			return 3
		case 65:
			return 2
		default:
			return 1
		}
	case numPlayers > 128 && numPlayers < 193:
		switch placing {
		case 1:
			return 160
		case 2:
			return 120
		case 3:
			return 90
		case 4:
			return 70
		case 5:
			return 60
		case 7:
			return 40
		case 9:
			return 30
		case 13:
			return 20
		case 17:
			return 15
		case 25:
			return 12
		case 33:
			return 9
		case 49:
			return 6
		case 65:
			return 3
		case 97:
			return 2
		default:
			return 1
		}
	case numPlayers > 192 && numPlayers < 257:
		switch placing {
		case 1:
			return 200
		case 2:
			return 160
		case 3:
			return 120
		case 4:
			return 90
		case 5:
			return 70
		case 7:
			return 60
		case 9:
			return 40
		case 13:
			return 30
		case 17:
			return 20
		case 25:
			return 15
		case 33:
			return 12
		case 49:
			return 9
		case 65:
			return 6
		case 97:
			return 3
		case 129:
			return 2
		default:
			return 1
		}
	case numPlayers > 256 && numPlayers < 385:
		switch placing {
		case 1:
			return 250
		case 2:
			return 200
		case 3:
			return 160
		case 4:
			return 120
		case 5:
			return 90
		case 7:
			return 70
		case 9:
			return 60
		case 13:
			return 40
		case 17:
			return 30
		case 25:
			return 20
		case 33:
			return 15
		case 49:
			return 12
		case 65:
			return 9
		case 97:
			return 6
		case 129:
			return 3
		case 193:
			return 2
		default:
			return 1
		}
	case numPlayers > 384 && numPlayers < 513:
		switch placing {
		case 1:
			return 300
		case 2:
			return 250
		case 3:
			return 200
		case 4:
			return 160
		case 5:
			return 120
		case 7:
			return 90
		case 9:
			return 70
		case 13:
			return 60
		case 17:
			return 40
		case 25:
			return 30
		case 33:
			return 20
		case 49:
			return 15
		case 65:
			return 12
		case 97:
			return 9
		case 129:
			return 6
		case 193:
			return 3
		case 257:
			return 2
		default:
			return 1
		}
	case numPlayers > 512 && numPlayers < 769:
		switch placing {
		case 1:
			return 400
		case 2:
			return 300
		case 3:
			return 250
		case 4:
			return 200
		case 5:
			return 160
		case 7:
			return 120
		case 9:
			return 90
		case 13:
			return 70
		case 17:
			return 60
		case 25:
			return 40
		case 33:
			return 30
		case 49:
			return 20
		case 65:
			return 15
		case 97:
			return 12
		case 129:
			return 9
		case 193:
			return 6
		case 257:
			return 3
		case 385:
			return 2
		default:
			return 1
		}
	case numPlayers > 768 && numPlayers < 1025:
		switch placing {
		case 1:
			return 500
		case 2:
			return 400
		case 3:
			return 300
		case 4:
			return 250
		case 5:
			return 200
		case 7:
			return 160
		case 9:
			return 120
		case 13:
			return 90
		case 17:
			return 70
		case 25:
			return 60
		case 33:
			return 40
		case 49:
			return 30
		case 65:
			return 20
		case 97:
			return 15
		case 129:
			return 12
		case 193:
			return 9
		case 257:
			return 6
		case 385:
			return 3
		case 513:
			return 2
		default:
			return 1
		}
	default:
		switch placing {
		case 1:
			return 700
		case 2:
			return 500
		case 3:
			return 400
		case 4:
			return 300
		case 5:
			return 250
		case 7:
			return 200
		case 9:
			return 160
		case 13:
			return 120
		case 17:
			return 90
		case 25:
			return 70
		case 33:
			return 60
		case 49:
			return 40
		case 65:
			return 30
		case 97:
			return 20
		case 129:
			return 15
		case 193:
			return 12
		case 257:
			return 9
		case 385:
			return 6
		case 513:
			return 3
		case 769:
			return 2
		default:
			return 1
		}
	}
}

func CalculateElo(winnerElo, loserElo int) (updatedWinnerElo, updatedLoserElo int) {
	winnerTransElo := math.Pow10(winnerElo / 400)
	loserTransElo := math.Pow10(loserElo / 400)
	winnerExpectedScore := winnerTransElo / (winnerTransElo + loserTransElo)
	loserExpectedScore := loserTransElo / (winnerTransElo + loserTransElo)
	updatedWinnerElo = int(math.Round(float64(winnerElo) + 32*(1-winnerExpectedScore)))
	updatedLoserElo = int(math.Round(float64(winnerElo) + 32*(0-loserExpectedScore)))
	return
}
