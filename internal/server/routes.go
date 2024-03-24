package server

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/markbates/goth/gothic"
	"github.com/rs/cors"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", s.HelloWorldHandler)
	r.Get("/health", s.healthHandler)
	r.Get("/auth/{provider}/callback", s.getAuthCallbackFunction)
	r.Get("/auth/{provider}", s.beginAuthProviderCall)
	r.Get("/logout/{provider}", s.logoutHandler)

	return cors.Default().Handler(r)
}

func (s *Server) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, _ := json.Marshal(s.db.Health())
	_, _ = w.Write(jsonResp)
}

func (s *Server) getAuthCallbackFunction(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")

	r = r.WithContext(context.WithValue(context.Background(), "provider", provider))

	log.Println("Starting user authentication...")

	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		log.Printf("Error during user authentication: %v", err)
		fmt.Fprintln(w, err)
		return
	}

	log.Printf("User authenticated successfully: %v", user)

	http.Redirect(w, r, "http://localhost:5173", http.StatusFound)
}

func (s *Server) beginAuthProviderCall(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")

	r = r.WithContext(context.WithValue(context.Background(), "provider", provider))

	log.Printf("Starting authentication process for provider: %s", provider)

	gothic.BeginAuthHandler(w, r)
}

func (s *Server) logoutHandler(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")

	r = r.WithContext(context.WithValue(context.Background(), "provider", provider))
	log.Printf("Starting logout process for provider: %s", provider)

	gothic.Logout(w, r)
	http.Redirect(w, r, "http://localhost:5173", http.StatusFound)
}
