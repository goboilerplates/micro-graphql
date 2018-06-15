package api

import "github.com/graphql-go/graphql"

// GetSampleType .
func GetSampleType() *graphql.Object {
	fields := graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
	}
	objectConfig := graphql.ObjectConfig{Name: "sample", Fields: fields}
	return graphql.NewObject(objectConfig)
}
