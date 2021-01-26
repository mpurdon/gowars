package resolvers

import (
	"context"

	"github.com/mpurdon/gowars/api"
	"github.com/mpurdon/gowars/errors"
)

// The QueryResolver is the entry point for all top-level read operations.
type QueryResolver struct {
	client *api.Client
}

func NewRoot(client *api.Client) (*QueryResolver, error) {
	if client == nil {
		return nil, errors.UnableToResolve
	}

	return &QueryResolver{client: client}, nil
}

// FilmsQueryArgs are the arguments for the "films" query.
type FilmsQueryArgs struct {
	// Title of the film. When nil, all films are fetched.
	Title *string
}

// Films resolves a list of films. If no arguments are provided, all films are fetched.
func (r QueryResolver) Films(ctx context.Context, args FilmsQueryArgs) (*[]*FilmResolver, error) {
	page, err := r.client.Films(ctx, strValue(args.Title))
	if err != nil {
		return nil, err
	}

	return NewFilms(ctx, NewFilmsArgs{Page: page})
}
