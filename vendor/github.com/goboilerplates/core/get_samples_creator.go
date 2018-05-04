package core

// CreateDefaultGetSamples .
func CreateDefaultGetSamples(names ...string) GetSamplesInteractor {
	var samples []*SampleEntity
	for _, item := range names {
		samples = append(samples, &SampleEntity{Name: item})
	}
	return GetSamplesInteractorImpl{Samples: samples}
}
