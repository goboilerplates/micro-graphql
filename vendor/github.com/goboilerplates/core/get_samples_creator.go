package core

import "time"

// CreateDefaultGetSamples .
func CreateDefaultGetSamples(names ...string) GetSamplesInteractor {
	var samples []*SampleEntity
	for _, item := range names {
		samples = append(samples, &SampleEntity{
			ID:   time.Now().Format(time.RFC850),
			Name: item,
		})
	}
	return GetSamplesInteractorImpl{Samples: samples}
}
