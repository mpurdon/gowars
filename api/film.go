package api

import (
	"context"
	"net/url"
)

// Film represents a film
type Film struct {
	CharacterURLs []string `json:"characters"`    //
	CreatedAt     string   `json:"created"`       // RFC 3339 format.
	DirectorName  string   `json:"director"`      //
	EditedAt      string   `json:"edited"`        // RFC 3339 format.
	EpisodeID     int64    `json:"episode_id"`    //
	OpeningCrawl  string   `json:"opening_crawl"` //
	PlanetURLs    []string `json:"planets"`       //
	ProducerNames string   `json:"producer"`      // Comma-separated if more than 1.
	ReleaseDate   string   `json:"release_date"`  // RFC 3339 format.
	SpeciesURLs   []string `json:"species"`       //
	StarshipURLs  []string `json:"starships"`     //
	Title         string   `json:"title"`         //
	URL           string   `json:"url"`           //
	VehicleURLs   []string `json:"vehicles"`      //
}

// FilmPage represents a page of films
type FilmPage struct {
	Count int64  `json:"count"`
	Films []Film `json:"results"`
}

// URLs for a page of films
func (p FilmPage) URLs() []string {
	urls := make([]string, len(p.Films))
	for i, f := range p.Films {
		urls[i] = f.URL
	}

	return urls
}

// Film gets a film from the API
func (c *Client) Film(ctx context.Context, url string) (Film, error) {
	r, err := c.NewRequest(ctx, url)
	if err != nil {
		return Film{}, err
	}

	var f Film
	if _, err = c.Do(r, &f); err != nil {
		return Film{}, err
	}

	return f, nil
}

// Films searches for films using the API
func (c *Client) Films(ctx context.Context, title string) (FilmPage, error) {
	q := url.Values{"search": {title}}
	r, err := c.NewRequest(ctx, "/films?"+q.Encode())
	if err != nil {
		return FilmPage{}, err
	}

	var fp FilmPage
	if _, err = c.Do(r, &fp); err != nil {
		return FilmPage{}, err
	}

	return fp, nil
}
