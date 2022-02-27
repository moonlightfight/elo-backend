// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type Admin struct {
	ID       string `json:"_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Character struct {
	ID        string `json:"_id"`
	Name      string `json:"name"`
	ImageLink string `json:"imageLink"`
}

type Jwt struct {
	Jwt string `json:"jwt"`
}

type LoginAdmin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Match struct {
	ID                       string    `json:"_id"`
	WinningPlayer            *Player   `json:"winningPlayer"`
	LosingPlayer             *Player   `json:"losingPlayer"`
	Date                     time.Time `json:"date"`
	WinningPlayerStartingElo int       `json:"winningPlayerStartingElo"`
	WinningPlayerEndingElo   int       `json:"winningPlayerEndingElo"`
	LosingPlayerStartingElo  int       `json:"losingPlayerStartingElo"`
	LosingPlayerEndingElo    int       `json:"losingPlayerEndingElo"`
}

type NewAdmin struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type NewPlayer struct {
	Username string `json:"username"`
	Rating   *int   `json:"rating"`
	Score    *int   `json:"score"`
}

type Player struct {
	ID            string        `json:"_id"`
	Slug          string        `json:"slug"`
	Username      string        `json:"username"`
	Rating        int           `json:"rating"`
	Score         int           `json:"score"`
	Country       *string       `json:"country"`
	Twitter       *string       `json:"twitter"`
	Twitch        *string       `json:"twitch"`
	Instagram     *string       `json:"instagram"`
	RealName      *string       `json:"realName"`
	Team          *Team         `json:"team"`
	MainCharacter *Character    `json:"mainCharacter"`
	SubCharacters []*Character  `json:"subCharacters"`
	Picture       *string       `json:"picture"`
	Controller    *string       `json:"controller"`
	Tournaments   []*Tournament `json:"tournaments"`
	Matches       []*Match      `json:"matches"`
}

type Team struct {
	ID           string  `json:"_id"`
	Slug         string  `json:"slug"`
	Name         string  `json:"name"`
	Abbreviation string  `json:"abbreviation"`
	Logo         *string `json:"logo"`
	Twitter      *string `json:"twitter"`
	Website      *string `json:"website"`
}

type Tournament struct {
	ID          string              `json:"_id"`
	Name        string              `json:"name"`
	Slug        string              `json:"slug"`
	Location    *string             `json:"location"`
	BracketURL  *string             `json:"bracketUrl"`
	NoBracket   bool                `json:"noBracket"`
	NumPlayers  int                 `json:"numPlayers"`
	Date        time.Time           `json:"date"`
	DateAdded   time.Time           `json:"dateAdded"`
	Replay      *string             `json:"replay"`
	Results     []*TournamentResult `json:"results"`
	Matches     []*Match            `json:"matches"`
	BracketType BracketType         `json:"bracketType"`
}

type TournamentResult struct {
	Place          int          `json:"place"`
	Points         int          `json:"points"`
	Player         *Player      `json:"player"`
	CharactersUsed []*Character `json:"CharactersUsed"`
}

type BracketType string

const (
	BracketTypeDoubleElim BracketType = "DOUBLE_ELIM"
	BracketTypeSingleElim BracketType = "SINGLE_ELIM"
	BracketTypeSwiss      BracketType = "SWISS"
	BracketTypeRoundRobin BracketType = "ROUND_ROBIN"
)

var AllBracketType = []BracketType{
	BracketTypeDoubleElim,
	BracketTypeSingleElim,
	BracketTypeSwiss,
	BracketTypeRoundRobin,
}

func (e BracketType) IsValid() bool {
	switch e {
	case BracketTypeDoubleElim, BracketTypeSingleElim, BracketTypeSwiss, BracketTypeRoundRobin:
		return true
	}
	return false
}

func (e BracketType) String() string {
	return string(e)
}

func (e *BracketType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = BracketType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid BracketType", str)
	}
	return nil
}

func (e BracketType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}