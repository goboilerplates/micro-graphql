package api

import (
	"github.com/gin-gonic/gin"
	"github.com/goboilerplates/micro-rest/core"
)

// GetSamplesAPI .
type GetSamplesAPI interface {
	All(context *gin.Context)
	AllByName(context *gin.Context)
}

// GetSamplesAPIImpl .
type GetSamplesAPIImpl struct {
	Interactor core.GetSamplesInteractor
}

// All .
func (api GetSamplesAPIImpl) All(context *gin.Context) {
	body, err := api.Interactor.All()
	if err != nil {
		context.JSON(404, err)
		return
	}

	context.JSON(200, body)
}

// AllByName .
func (api GetSamplesAPIImpl) AllByName(context *gin.Context) {
	keyword := context.Param("keyword")
	body, err := api.Interactor.AllByName(keyword)
	if err != nil {
		context.JSON(404, err)
		return
	}

	context.JSON(200, body)
}
