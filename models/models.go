package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Admin struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Email    string             `json:"email,omitempty" bson:"email,omitempty"`
	Username string             `json:"username,omitempty" bson:"username,omitempty"`
	Password string             `json:"password,omitempty" bson:"password,omitempty"`
}

type Player struct {
	ID            primitive.ObjectID   `json:"_id,omitempty" bson:"_id,omitempty"`
	Slug          string               `json:"slug,omitempty" bson:"slug,omitempty"`
	Username      string               `json:"username,omitempty" bson:"username,omitempty"`
	Country       string               `json:"country,omitempty" bson:"country,omitempty"`
	Ranking       int                  `json:"ranking,omitempty" bson:"ranking,omitempty"`
	Points        int                  `json:"points,omitempty" bson:"points,omitempty"`
	Controller    string               `json:"controller,omitempty" bson:"controller,omitempty"`
	RealName      string               `json:"realName,omitempty" bson:"realName,omitempty"`
	Twitter       string               `json:"twitter,omitempty" bson:"twitter,omitempty"`
	Twitch        string               `json:"twitch,omitempty" bson:"twitch,omitempty"`
	Picture       string               `json:"picture,omitempty" bson:"picture,omitempty"`
	Tournaments   []primitive.ObjectID `json:"tournaments,omitempty" bson:"tournaments,omitempty"`
	Matches       []primitive.ObjectID `json:"matches,omitempty" bson:"matches,omitempty"`
	MainCharacter primitive.ObjectID   `json:"mainCharacter,omitempty" bson:"mainCharacter,omitempty"`
	SubCharacters []primitive.ObjectID `json:"subCharacters,omitempty" bson:"subCharacters,omitempty"`
}

type Team struct {
	ID           primitive.ObjectID   `json:"_id,omitempty" bson:"_id,omitempty"`
	Slug         string               `json:"slug,omitempty" bson:"slug,omitempty"`
	Name         string               `json:"name,omitempty" bson:"name,omitempty"`
	Abbreviation string               `json:"abbreviation,omitempty" bson:"abbreviation,omitempty"`
	Members      []primitive.ObjectID `json:"members,omitempty" bson:"members,omitempty"`
	Logo         string               `json:"logo,omitempty" bson:"logo,omitempty"`
	Twitter      string               `json:"twitter,omitempty" bson:"twitter,omitempty"`
	Website      string               `json:"website,omitempty" bson:"website,omitempty"`
}

type Tournament struct {
	ID         primitive.ObjectID   `json:"_id,omitempty" bson:"_id,omitempty"`
	Slug       string               `json:"slug,omitempty" bson:"slug,omitempty"`
	Location   string               `json:"location,omitempty" bson:"location,omitempty"`
	BracketUrl string               `json:"bracketUrl,omitempty" bson:"bracketUrl,omitempty"`
	NumPlayers int                  `json:"numPlayers,omitempty" bson:"numPlayers,omitempty"`
	Date       string               `json:"date,omitempty" bson:"date,omitempty"`
	Replay     string               `json:"replay,omitempty" bson:"replay,omitempty"`
	Results    []TournamentResults  `json:"results,omitempty" bson:"results,omitempty"`
	Matches    []primitive.ObjectID `json:"matches,omitempty" bson:"matches,omitempty"`
}

type TournamentResults struct {
	Place          int                  `json:"place,omitempty" bson:"place,omitempty"`
	Points         int                  `json:"points,omitempty" bson:"points,omitempty"`
	Player         primitive.ObjectID   `json:"player,omitempty" bson:"player,omitempty"`
	CharactersUsed []primitive.ObjectID `json:"charactersUsed,omitempty" bson:"charactersUsed,omitempty"`
}

type Character struct {
	ID   primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Slug string             `json:"slug,omitempty" bson:"slug,omitempty"`
	Name string             `json:"name,omitempty" bson:"name,omitempty"`
}

type Match struct {
	ID                       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	WinningPlayer            primitive.ObjectID `json:"winningPlayer,omitempty" bson:"winningPlayer,omitempty"`
	LosingPlayer             primitive.ObjectID `json:"losingPlayer,omitempty" bson:"losingPlayer,omitempty"`
	Date                     string             `json:"date,omitempty" bson:"date,omitempty"`
	WinningPlayerStartingElo int                `json:"winningPlayerStartingElo,omitempty" bson:"winningPlayerStartingElo,omitempty"`
	WinningPlayerEndingElo   int                `json:"winningPlayerEndingElo,omitempty" bson:"winningPlayerEndingElo,omitempty"`
	LosingPlayerStartingElo  int                `json:"losingPlayerStartingElo,omitempty" bson:"losingPlayerStartingElo,omitempty"`
	LosingPlayerEndingElo    int                `json:"losingPlayerEndingElo,omitempty" bson:"losingPlayerEndingElo,omitempty"`
}

type LoginData struct {
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}
