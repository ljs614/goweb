package myapp

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndexPathHandler(t *testing.T) {
	assert := assert.New(t)
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)

	mux := NewHTTPHandler()
	mux.ServeHTTP(res, req)

	indexHandler(res, req)
	assert.Equal(http.StatusOK, res.Code)

	data, _ := ioutil.ReadAll(res.Body)
	assert.Equal("Hello World", string(data))
}

func TestBarPathHandler_WithoutName(t *testing.T) {
	assert := assert.New(t)
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/bar", nil)

	mux := NewHTTPHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)

	data, _ := ioutil.ReadAll(res.Body)
	assert.Equal("Hello World", string(data))
}