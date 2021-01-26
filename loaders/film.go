package loaders

import (
	"context"
	"sync"

	"github.com/mpurdon/gowars/api"
	"github.com/mpurdon/gowars/errors"
	"github.com/nicksrandall/dataloader"
)

func LoadFilm(ctx context.Context, url string) (api.Film, error) {
	var film api.Film

	ldr, err := extract(ctx, filmLoaderKey)
	if err != nil {
		return film, err
	}

	data, err := ldr.Load(ctx, dataloader.StringKey(url))()
	if err != nil {
		return film, err
	}

	film, ok := data.(api.Film)
	if !ok {
		return film, errors.WrongType(film, data)
	}

	return film, nil
}

func LoadFilms(ctx context.Context, urls []string) (FilmResults, error) {
	var results []FilmResult
	ldr, err := extract(ctx, filmLoaderKey)
	if err != nil {
		return results, err
	}

	data, errs := ldr.LoadMany(ctx, dataloader.NewKeysFromStrings(urls))()
	results = make([]FilmResult, 0, len(urls))

	for i, d := range data {
		var e error
		if errs != nil {
			e = errs[i]
		}

		film, ok := d.(api.Film)
		if !ok && e == nil {
			e = errors.WrongType(film, d)
		}

		results = append(results, FilmResult{Film: film, Error: e})
	}

	return results, nil
}

// FilmResult is the (data, error) pair result of loading a specific key.
type FilmResult struct {
	Film  api.Film
	Error error
}

// FilmResults is a named type, so methods can be attached to []FilmResult.
type FilmResults []FilmResult

// WithoutErrors filters any result pairs with non-nil errors.
func (results FilmResults) WithoutErrors() []api.Film {
	var films = make([]api.Film, 0, len(results))

	for _, r := range results {
		if r.Error != nil {
			continue
		}

		films = append(films, r.Film)
	}

	return films
}

// PrimeFilms primes the cache with the provided films
func PrimeFilms(ctx context.Context, page api.FilmPage) error {
	ldr, err := extract(ctx, filmLoaderKey)
	if err != nil {
		return err
	}

	for _, f := range page.Films {
		ldr.Prime(ctx, dataloader.StringKey(f.URL), f)
	}

	return nil
}

// filmGetter specifies an interface for getting films
type filmGetter interface {
	Film(ctx context.Context, url string) (api.Film, error)
}

// FilmLoader contains the client required to load film resources.
type filmLoader struct {
	get filmGetter
}

// newFilmLoader returns a file loader
func newFilmLoader(client filmGetter) dataloader.BatchFunc {
	return filmLoader{get: client}.loadBatch
}

// loadBatch load a batch of films with the loader
func (ldr filmLoader) loadBatch(ctx context.Context, urls dataloader.Keys) []*dataloader.Result {
	var (
		n       = len(urls)
		results = make([]*dataloader.Result, n)
		wg      sync.WaitGroup
	)

	wg.Add(n)

	// @Todo change results to a channel - duh
	for i, url := range urls {
		// load data from the API asynchronously
		go func(i int, url dataloader.Key) {
			defer wg.Done()

			resp, err := ldr.get.Film(ctx, url.String())
			results[i] = &dataloader.Result{Data: resp, Error: err}
		}(i, url)
	}

	wg.Wait()

	return results
}
