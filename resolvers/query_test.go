package resolvers

import (
	"testing"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/mpurdon/gowars/schema"
)

func TestResolversSatisfySchema(t *testing.T) {
	rootResolver := &QueryResolver{}
	_, err := graphql.ParseSchema(schema.String(), rootResolver)
	if err != nil {
		t.Error(err)
	}
}
