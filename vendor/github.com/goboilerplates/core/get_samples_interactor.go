package core

import (
	"errors"
	"strings"
)

// GetSamplesInteractor is the interface of GetSamplesInteractor.
type GetSamplesInteractor interface {
	All() ([]*SampleEntity, error)
	AllByName(keyword string) ([]*SampleEntity, error)
}

// GetSamplesInteractorImpl is the implementation of GetSamplesInteractor interface.
type GetSamplesInteractorImpl struct {
	Samples []*SampleEntity
}

// All gets all samples.
func (interactor GetSamplesInteractorImpl) All() ([]*SampleEntity, error) {
	return interactor.Samples, nil
}

// AllByName gets all samples that have name matched with keyword.
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
