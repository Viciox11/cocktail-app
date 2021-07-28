package test

import (
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/ioartigiano/ioartigiano-be/pkg/email"
	gomail "gopkg.in/mail.v2"
	"net/http"
)

// MockRouter creates a routing.Router for testing APIs.
func MockRouter() *mux.Router {
	return mux.NewRouter()
}

// MockHeader creates an http.Header for testing APIs.
func MockHeader() http.Header {
	return http.Header{}
}

// MockValidator creates an validator for testing APIs.
func MockValidator() *validator.Validate {
	return validator.New()
}

// MockEmailClient creates an emailClient for testing APIs.
func MockEmailClient() *email.EmailClient {
	return email.NewEmailHandler(gomail.NewMessage())
}
