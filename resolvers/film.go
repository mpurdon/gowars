package resolvers

import (
	"context"
	"strings"
	"time"

	"github.com/graph-gophers/graphql-go"
	"github.com/mpurdon/gowars/api"
	"github.com/mpurdon/gowars/errors"
	"github.com/mpurdon/gowars/loaders"
)

// FilmResolver resolves the Film type.
type FilmResolver struct {
	film api.Film
}

type NewFilmsArgs struct {
	Page api.FilmPage
	URLs []string
}

type NewFilmArgs struct {
	Film api.Film
	URL  string
}

func NewFilm(ctx context.Context, args NewFilmArgs) (*FilmResolver, error) {
	var film api.Film
	var err error

	switch {
	case args.Film.URL != "":
		film = args.Film
	case args.URL != "":
		film, err = loaders.LoadFilm(ctx, args.URL)
	default:
		err = errors.UnableToResolve
	}

	if err != nil {
		return nil, err
	}

	return &FilmResolver{film: film}, nil
}

func NewFilms(ctx context.Context, args NewFilmsArgs) (*[]*FilmResolver, error) {
	err := loaders.PrimeFilms(ctx, args.Page)
	if err != nil {
		return nil, err
	}

	results, err := loaders.LoadFilms(ctx, append(args.URLs, args.Page.URLs()...))
	if err != nil {
		return nil, err
	}

	var (
		films     = results.WithoutErrors()
		resolvers = make([]*FilmResolver, 0, len(films))
		errs      errors.Errors
	)

	for i, film := range films {
		resolver, err := NewFilm(ctx, NewFilmArgs{Film: film})
		if err != nil {
			errs = append(errs, errors.WithIndex(err, i))
		}

		resolvers = append(resolvers, resolver)
	}

	return &resolvers, errs.Err()
}

// ID resolves the film's unique identifier.
func (r *FilmResolver) ID() graphql.ID {
	return extractID(r.film.URL)
}

// Episode resolves the episode number of this film.
func (r *FilmResolver) Episode() int32 {
	return int32(r.film.EpisodeID)
}

// DirectorName resolves the name this film's director.
func (r *FilmResolver) DirectorName() string {
	return r.film.DirectorName
}

// ProducerNames resolves a list of names of this film's producers.
func (r *FilmResolver) ProducerNames() []string {
	return strings.Split(r.film.ProducerNames, ", ")
}

// ReleaseDate resolves the time of the film release in the original creator country.
func (r *FilmResolver) ReleaseDate() (graphql.Time, error) {
	t, err := time.Parse("2006-01-02", r.film.ReleaseDate)
	return graphql.Time{Time: t}, err
}

// OpeningCrawl resolves the opening crawl text
func (r *FilmResolver) OpeningCrawl() string {
	return r.film.OpeningCrawl
}

// // Characters resolves a list of characters that are in this film.
// func (r *FilmResolver) Characters(ctx context.Context) (*[]*PersonResolver, error) {
// 	return NewPeople(ctx, NewPeopleArgs{URLs: r.film.CharacterURLs})
// }

// CreatedAt resolves the RFC3339 date format of the time this resource was created.
func (r *FilmResolver) CreatedAt(ctx context.Context) (graphql.Time, error) {
	t, err := time.Parse(time.RFC3339, r.film.CreatedAt)
	return graphql.Time{Time: t}, err
}

// EditedAt resolves the RFC3339 date format of the time this resource was created.
func (r *FilmResolver) EditedAt(ctx context.Context) (*graphql.Time, error) {
	if r.film.EditedAt == "" {
		return nil, nil
	}

	t, err := time.Parse(time.RFC3339, r.film.EditedAt)
	return &graphql.Time{Time: t}, err
}
