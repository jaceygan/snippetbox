package main

import (
	"net/http"
	"testing"

	"github.com/jaceygan/snippetbox/internal/assert"
)

func TestPing(t *testing.T) {
	app := newTestApplication(t)

	ts := newTestServer(t, app.routes())
	defer ts.Close()

	statusCode, _, body := ts.get(t, "/ping")

	assert.Equal(t, statusCode, http.StatusOK)
	assert.Equal(t, body, "OK")

}

func TestSnippetView(t *testing.T) {
	app := newTestApplication(t)

	ts := newTestServer(t, app.routes())
	defer ts.Close()

	tests := []struct {
		name               string
		urlPath            string
		expectedStatusCode int
		expectedBody       string
	}{
		{
			name:               "Valid Snippet ID",
			urlPath:            "/snippet/view/1",
			expectedStatusCode: http.StatusOK,
			expectedBody:       "Mock Snippet",
		},
		{
			name:               "Non-Existent Snippet ID",
			urlPath:            "/snippet/view/2",
			expectedStatusCode: http.StatusNotFound,
		},
		{
			name:               "Negative Snippet ID",
			urlPath:            "/snippet/view/-1",
			expectedStatusCode: http.StatusNotFound,
		},
		{
			name:               "Decimal Snippet ID",
			urlPath:            "/snippet/view/1.5",
			expectedStatusCode: http.StatusNotFound,
		},
		{
			name:               "String Snippet ID",
			urlPath:            "/snippet/view/abc",
			expectedStatusCode: http.StatusNotFound,
		},
		{
			name:               "Empty Snippet ID",
			urlPath:            "/snippet/view/",
			expectedStatusCode: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			statusCode, _, body := ts.get(t, tt.urlPath)

			assert.Equal(t, statusCode, tt.expectedStatusCode)
			if tt.expectedBody != "" {
				assert.StringContains(t, body, tt.expectedBody)
			}
		})
	}
}
