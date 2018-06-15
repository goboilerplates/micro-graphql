package api

import (
	"log"

	"github.com/goboilerplates/core"
	"github.com/graphql-go/graphql"
)

// GetSamplesAPI is the interface for GetSamplesAPI.
type GetSamplesAPI interface {
	All(params graphql.ResolveParams) (interface{}, error)
}

// GetSamplesAPIImpl is the implementation of GetSamplesAPI interface.
type GetSamplesAPIImpl struct {
	Interactor core.GetSamplesInteractor
}

// All gets all samples.
func (api GetSamplesAPIImpl) All(params graphql.ResolveParams) (interface{}, error) {
	keyword := params.Args["keyword"]
	log.Println("request to get samples with keyword: ", keyword)
	if keyword != nil {
		samples, err := api.Interactor.AllByName(keyword.(string))
		return samples, err
	}
	samples, err := api.Interactor.All()
	return samples, err
}
