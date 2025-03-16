package main

import (
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain1(t *testing.T) {
	tests := []struct {
		description   string
		route         string
		expectedError bool
		expectedCode  int
		expectedBody  string
	}{
		{
			description:   "get all blogs",
			route:         "/api/v1/blog-posts",
			expectedError: false,
			expectedCode:  500,
			expectedBody:  "",
		},
		{
			description:   "non existing route",
			route:         "/i-dont-exist",
			expectedError: false,
			expectedCode:  404,
			expectedBody:  "Cannot GET /i-dont-exist",
		},
	}

	app := setup()

	for _, test := range tests {
		req, _ := http.NewRequest(http.MethodGet, test.route, nil)
		resp, err := app.Test(req, -1)
		assert.NoError(t, err)
		assert.Equal(t, test.expectedCode, resp.StatusCode)
		if test.description == "non existing route" {
			body, _ := io.ReadAll(resp.Body)
			assert.Equal(t, test.expectedBody, string(body))
		}
	}
}
