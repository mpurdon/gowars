package main

import (
	"net/http"
	"time"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/mpurdon/gowars/api"
	"github.com/mpurdon/gowars/handlers"
	"github.com/mpurdon/gowars/loaders"
	"github.com/mpurdon/gowars/log"
	"github.com/mpurdon/gowars/resolvers"
	"github.com/mpurdon/gowars/schema"
)

func main() {
	// Tweak configuration values here.
	var (
		addr              = ":8000"
		readHeaderTimeout = 1 * time.Second
		writeTimeout      = 10 * time.Second
		idleTimeout       = 90 * time.Second
		maxHeaderBytes    = http.DefaultMaxHeaderBytes
	)

	c := api.NewClient(http.DefaultClient) // TODO: don't use the default client.

	root, err := resolvers.NewRoot(c)
	if err != nil {
		log.Error().
			Err(err).
			Msg("could not set up the resolvers")
	}

	// Create the request handler; inject dependencies.
	h := handlers.GraphQL{
		// Parse and validate schema. Panic if unable to do so.
		Schema:  graphql.MustParseSchema(schema.String(), root),
		Loaders: loaders.Initialize(c),
		Logger: log.Logger(),
	}

	// Register handlers to routes.
	mux := http.NewServeMux()
	mux.Handle("/", handlers.GraphiQL{})
	mux.Handle("/graphql/", h)
	mux.Handle("/graphql", h) // Register without a trailing slash to avoid redirect.

	// Configure the HTTP server.
	s := &http.Server{
		Addr:              addr,
		Handler:           mux,
		ReadHeaderTimeout: readHeaderTimeout,
		WriteTimeout:      writeTimeout,
		IdleTimeout:       idleTimeout,
		MaxHeaderBytes:    maxHeaderBytes,
	}

	// begin listening for requests
	log.Info().
		Str("server_address", s.Addr).
		Dur("write_timeout", writeTimeout).
		Dur("idle_timeout", idleTimeout).
		Msg("listening for requests")

	if err = s.ListenAndServe(); err != nil {
		log.Error().
			Err(err).
			Msg("ListenAndServe failed to start")
	}

	// TODO: intercept shutdown signals for cleanup of connections.
	log.Info().Msg("shutting down server")
}
