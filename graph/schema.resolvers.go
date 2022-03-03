package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/moonlightfight/elo-backend/constants"
	"github.com/moonlightfight/elo-backend/database"
	"github.com/moonlightfight/elo-backend/graph/generated"
	"github.com/moonlightfight/elo-backend/graph/model"
)

func (r *mutationResolver) CreatePlayer(ctx context.Context, input model.NewPlayer) (*model.Player, error) {
	panic(fmt.Errorf("not implemented"))
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
var mongodbUri = constants.GetDbUri()
var db = database.Connect(mongodbUri)
