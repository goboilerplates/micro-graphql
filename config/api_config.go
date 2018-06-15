package config

import (
	"log"

	"strings"

	"github.com/gin-gonic/gin"
	"github.com/goboilerplates/core"
	"github.com/goboilerplates/micro-graphql/api"
	"github.com/graphql-go/graphql"
)

var schema, _ = CreateSchema()

// SetAPIPaths sets API paths.
func SetAPIPaths(router *gin.Engine) gin.IRoutes {
	routes := router.GET("/api", HanldeAPIRequest)
	routes = router.POST("/api", HanldeAPIPOSTRequest)
	return routes
}

// CreateSchema .
func CreateSchema() (graphql.Schema, error) {
	var getSamplesAPI api.GetSamplesAPI

	// inject GetSamplesAPIImpl .
	getSamplesAPI = api.GetSamplesAPIImpl{
		Interactor: core.CreateDefaultGetSamples("Kaka", "Ronaldo"),
	}

	sampleObject := api.GetSampleType()

	fields := graphql.Fields{
		"samples": &graphql.Field{
			Type: graphql.NewList(sampleObject),
			Args: graphql.FieldConfigArgument{
				"keyword": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: getSamplesAPI.All,
		},
	}
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	return graphql.NewSchema(schemaConfig)
}

// HanldeAPIRequest .
func HanldeAPIRequest(context *gin.Context) {
	query := context.Query("query")
	log.Println("request with query: ", query)
	result := ExecuteQuery(query)
	if len(result.Errors) > 0 {
		context.JSON(404, result.Errors[0])
		return
	}
	context.JSON(200, result.Data)
}

// HanldeAPIPOSTRequest .
func HanldeAPIPOSTRequest(context *gin.Context) {
	data, _ := context.GetRawData()
	query := string(data)
	query = strings.Split(query, "=")[1]
	log.Println("request with query: ", query)
	result := ExecuteQuery(query)
	if len(result.Errors) > 0 {
		context.JSON(404, result.Errors[0])
		return
	}
	context.JSON(200, result.Data)
}

// ExecuteQuery .
func ExecuteQuery(query string) *graphql.Result {
	return graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
}
