package octokit

import (
	"github.com/bmizerany/assert"
	"net/http"
	"strings"
	"testing"
)

func TestResponseError_Error_400(t *testing.T) {
	setup()
	defer tearDown()

	mux.HandleFunc("/error", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, `{"message":"Problems parsing JSON"}`, 400)
	})

	_, err := client.Request("GET", testURLOf("error"), nil, nil)
	assert.Tf(t, strings.Contains(err.Error(), "400 - Problems parsing JSON"), "%s", err.Error())
}

func TestResponseError_Error_422_error(t *testing.T) {
	setup()
	defer tearDown()

	mux.HandleFunc("/error", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, `{"error":"No repository found for hubtopic"}`, 422)
	})

	_, err := client.Request("GET", testURLOf("error"), nil, nil)
	assert.Tf(t, strings.Contains(err.Error(), "Error: No repository found for hubtopic"), "%s", err.Error())
}

func TestResponseError_Error_422_error_summary(t *testing.T) {
	setup()
	defer tearDown()

	mux.HandleFunc("/error", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, `{"message":"Validation Failed", "errors": [{"resource":"Issue", "field": "title", "code": "missing_field"}]}`, 422)
	})

	_, err := client.Request("GET", testURLOf("error"), nil, nil)
	assert.Tf(t, strings.Contains(err.Error(), "422 - Validation Failed"), "%s", err.Error())
	assert.Tf(t, strings.Contains(err.Error(), "missing_field error caused by title field on Issue resource"), "%s", err.Error())
}

func TestResponseError_Error_415(t *testing.T) {
	setup()
	defer tearDown()

	mux.HandleFunc("/error", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, `{"message":"Unsupported Media Type", "documentation_url":"http://developer.github.com/v3"}`, 415)
	})

	_, err := client.Request("GET", testURLOf("error"), nil, nil)
	assert.Tf(t, strings.Contains(err.Error(), "415 - Unsupported Media Type"), "%s", err.Error())
	assert.Tf(t, strings.Contains(err.Error(), "// See: http://developer.github.com/v3"), "%s", err.Error())
}