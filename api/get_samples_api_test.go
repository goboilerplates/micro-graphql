package api_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/bmizerany/assert"
	"github.com/gin-gonic/gin"
	"github.com/goboilerplates/micro-graphql/config"
)

// SetupRouter .
func SetupRouter() *gin.Engine {
	router := gin.Default()
	gin.SetMode(gin.TestMode)
	return router
}

// TestGetAllSamples .
func TestGetAllSamples(test *testing.T) {
	router := SetupRouter()

	config.SetAPIPaths(router)

	url := "/api?query={samples{name,id}}"
	req, _ := http.NewRequest("GET", url, nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(test, resp.Code, 200)

	url = "/api"
	req, _ = http.NewRequest("POST", url, strings.NewReader("query={samples{name,id}}"))
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(test, resp.Code, 200)

	url = "/api"
	req, _ = http.NewRequest("GET", url, strings.NewReader("query=AAAA"))
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(test, resp.Code, 404)

	url = "/api"
	req, _ = http.NewRequest("POST", url, strings.NewReader("query=AAAA"))
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(test, resp.Code, 404)
}

// TestAllByNameSamples .
func TestAllByNameSamples(test *testing.T) {
	router := SetupRouter()

	config.SetAPIPaths(router)

	url := `/api?query={samples(keyword:"ka"){name,id}}`
	req, _ := http.NewRequest("GET", url, nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(test, resp.Code, 200)

	url = `/api?query={samples(keyword:"ro"){name,id}}`
	req, _ = http.NewRequest("GET", url, strings.NewReader(`query={samples(keyword:"ro"){name,id}}`))
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(test, resp.Code, 200)

	url = "/api"
	req, _ = http.NewRequest("POST", url, strings.NewReader(`query={samples(keyword:"ka"){name,id}}`))
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(test, resp.Code, 200)

	req, _ = http.NewRequest("POST", url, strings.NewReader(`query={samples(keyword:"ro"){name,id}}`))
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(test, resp.Code, 200)
}
