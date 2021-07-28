package test

import (
	"bytes"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// APITestCase represents the data needed to describe an API test case.
type APITestCase struct {
	Name         string
	Method, URL  string
	Body         string
	Header       http.Header
	WantStatus   int
	WantResponse string
	Role         int
}

// Endpoint tests an HTTP endpoint using the given APITestCase spec.
func Endpoint(t *testing.T, router *mux.Router, tc APITestCase) {
	t.Run(tc.Name, func(t *testing.T) {
		req, _ := http.NewRequest(tc.Method, tc.URL, bytes.NewBufferString(tc.Body))
		if tc.Header != nil {
			req.Header = tc.Header
		}
		res := httptest.NewRecorder()
		if req.Header.Get("Content-Type") == "" {
			req.Header.Set("Content-Type", "application/json")
		}

		var role string
		switch tc.Role {
		case 0:
			role = "NOT NECESSARY"
		case 1:
			role = "ADMIN"
		case 2:
			role = "USER"
		}

		if tc.Role != 0 {
			context.Set(req, "user_role", role)
		}
		t.Log("Role Used: ", role)

		t.Log("Request: ", req)
		router.ServeHTTP(res, req)
		t.Log("Response Code: ", res.Code, " Body: ", res.Body.String())

		assert.Equal(t, tc.WantStatus, res.Code, "status mismatch: ", res.Body.String())
		if tc.WantResponse != "" {
			pattern := strings.Trim(tc.WantResponse, "*")
			if pattern != tc.WantResponse {
				assert.Contains(t, res.Body.String(), pattern, "response mismatch")
			} else {
				assert.JSONEq(t, tc.WantResponse, res.Body.String(), "response mismatch")
			}
		}
	})
}
