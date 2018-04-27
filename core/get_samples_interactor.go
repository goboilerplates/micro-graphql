package core

import (
	"errors"
	"strings"
)

// GetSamplesInteractor .
type GetSamplesInteractor interface {
	All() ([]*SampleEntity, error)
	AllByName(keyword string) ([]*SampleEntity, error)
}

// GetSamplesInteractorImpl is the implementation of GetSamplesInteractor interface.
type GetSamplesInteractorImpl struct {
	Samples []*SampleEntity
}

// All .
func (interactor GetSamplesInteractorImpl) All() ([]*SampleEntity, error) {
	return interactor.Samples, nil
}

// AllByName .
func (interactor GetSamplesInteractorImpl) AllByName(keyword string) ([]*SampleEntity, error) {
	keyword = strings.ToLower(keyword)
	var list []*SampleEntity
	for _, item := range interactor.Samples {
		name := item.NameLower()
		if strings.Contains(name, keyword) {
			list = append(list, item)
		}
	}
	if len(list) == 0 {
		return list, errors.New("Not Found")
	}
	return list, nil
}
