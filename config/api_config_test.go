package config_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bmizerany/assert"
	"github.com/gin-gonic/gin"
	"github.com/goboilerplates/micro-rest/config"
)

// SetupRouter .
func SetupRouter() *gin.Engine {
	router := gin.Default()
	gin.SetMode(gin.TestMode)
	return router
}

// TestSetAPIPaths .
func TestSetAPIPaths(test *testing.T) {
	router := SetupRouter()

	config.SetAPIPaths(router)

	url := "/api/v1/samples"
	req, _ := http.NewRequest("GET", url, nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(test, resp.Code, 200)

	url = "/api/v1/samples/ka"
	req, _ = http.NewRequest("GET", url, nil)
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(test, resp.Code, 200)

	url = "/api/v1/samples/b"
	req, _ = http.NewRequest("GET", url, nil)
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(test, resp.Code, 404)

	url = "/api/v1/samples/nal"
	req, _ = http.NewRequest("GET", url, nil)
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(test, resp.Code, 200)
}
