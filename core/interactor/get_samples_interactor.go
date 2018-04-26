package interactor

import (
	"errors"
	"strings"

	"github.com/goboilerplates/micro-rest/core/entity"
)

// GetSamplesInteractor .
type GetSamplesInteractor interface {
	All() ([]*entity.SampleEntity, error)
	AllByName(keyword string) ([]*entity.SampleEntity, error)
}

// GetSamplesInteractorImpl is the implementation of GetSamplesInteractor interface.
type GetSamplesInteractorImpl struct {
	Samples []*entity.SampleEntity
}

// All .
func (interactor GetSamplesInteractorImpl) All() ([]*entity.SampleEntity, error) {
	return interactor.Samples, nil
}

// AllByName .
func (interactor GetSamplesInteractorImpl) AllByName(keyword string) ([]*entity.SampleEntity, error) {
	keyword = strings.ToLower(keyword)
	var list []*entity.SampleEntity
	for _, item := range interactor.Samples {
		name := strings.ToLower(item.Name)
		if strings.Contains(name, keyword) {
			list = append(list, item)
		}
	}
	if len(list) == 0 {
		return list, errors.New("Not Found")
	}
	return list, nil
}
