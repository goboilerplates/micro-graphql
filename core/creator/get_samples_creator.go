package creator

import (
	"github.com/goboilerplates/micro-rest/core/entity"
	"github.com/goboilerplates/micro-rest/core/interactor"
)

// DefaultGetSamples .
func DefaultGetSamples() interactor.GetSamplesInteractor {
	var samples []*entity.SampleEntity
	names := [2]string{"Kaka", "Ronaldo"}
	for _, item := range names {
		samples = append(samples, &entity.SampleEntity{Name: item})
	}
	return interactor.GetSamplesInteractorImpl{Samples: samples}
}
