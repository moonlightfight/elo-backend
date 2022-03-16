package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/moonlightfight/elo-backend/constants"
	"github.com/moonlightfight/elo-backend/database"
	"github.com/moonlightfight/elo-backend/graph/generated"
	"github.com/moonlightfight/elo-backend/graph/model"
	"github.com/moonlightfight/elo-backend/helpers"
)

func (r *mutationResolver) CreatePlayer(ctx context.Context, input model.NewPlayer) (*model.Player, error) {
	lowerName := strings.ToLower(input.Username)

	specialCharRegex, err := regexp.Compile(`([^A-Za-z0-9\s_-])`)

	if err != nil {
		log.Println(err)
	}

	re := strings.NewReplacer("_", "-", " ", "-")

	noSpecialChar := specialCharRegex.ReplaceAllString(lowerName, "")

	slug := re.Replace(noSpecialChar)

	player := model.Player{
		Username: input.Username,
		Rating:   1200,
		Score:    0,
		Slug:     slug,
	}

	return db.CreatePlayer(player), nil
}

func (r *mutationResolver) CreateAdmin(ctx context.Context, input model.NewAdmin) (*model.Admin, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) LoginAdmin(ctx context.Context, input model.LoginAdmin) (*model.Jwt, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateTeam(ctx context.Context, input model.NewTeam) (*model.Team, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateTournament(ctx context.Context, input model.NewTournament) (*model.Tournament, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Players(ctx context.Context) ([]*model.Player, error) {
	return db.GetPlayers(), nil
}

func (r *queryResolver) Characters(ctx context.Context) ([]*model.Character, error) {
	return db.GetCharacters(), nil
}

func (r *queryResolver) Tournaments(ctx context.Context) ([]*model.Tournament, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Matches(ctx context.Context) ([]*model.Match, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Teams(ctx context.Context) ([]*model.Team, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) TournamentFromAPI(ctx context.Context, input model.TournamentFromAPI) (*model.APIReturnedTournament, error) {
	if strings.Contains(input.URL, "challonge") {
		var tournamentId string
		var subDomain interface{}
		if strings.Contains(input.URL, "https://challonge.com/") {
			subDomain = nil
			tournamentId = strings.Replace(input.URL, "https://challonge.com/", "", -1)
		} else {
			re := strings.NewReplacer("https://", "", ".challonge.com", "")
			trunc := re.Replace(input.URL)
			subDomain = strings.TrimRight(trunc, "/")
			tournamentId = strings.TrimLeft(input.URL, "/")
		}
		return helpers.GetChallongeBracket(tournamentId, subDomain), nil
	} else if strings.Contains(input.URL, "smash") {
		// trim the url down to the obscenely long event slug bc fucking smashers
		re := strings.NewReplacer("https://smash.gg/", "", "/overview", "")
		slug := re.Replace(input.URL)
		return helpers.GetSmashBracket(slug), nil
	} else {
		return nil, fmt.Errorf("invalid bracket url %s", input.URL)
	}
}

func (r *queryResolver) Player(ctx context.Context, input model.SinglePlayer) (*model.Player, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Team(ctx context.Context, input model.SingleTeam) (*model.Team, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Tournament(ctx context.Context, input model.SingleTournament) (*model.Tournament, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
var mongodbUri = constants.GetEnvVar("MONGODB_URI")
var db = database.Connect(mongodbUri)
