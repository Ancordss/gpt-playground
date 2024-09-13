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
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", s.HelloWorldHandler)

	r.Get("/health", s.healthHandler)
	r.Get("/auth/{provider}/callback", s.getAuthCallback)
	r.Get("/logout/{provider}", s.logout)
	r.Get("/auth/{provider}", s.getAuth)

	return r
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

func (s *Server) getAuthCallback(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")
	r = r.WithContext(context.WithValue(context.Background(), "provider", provider))

	// Autenticación con Gothic
	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		fmt.Fprintln(w, "Error en la autenticación:", err)
		return
	}

	// Extraer los datos relevantes del usuario autenticado
	email := user.Email
	id := user.UserID // Este sería el "sub" (subject) en OAuth
	token := user.AccessToken

	// Guardar en la base de datos
	err = s.db.InsertUser(id, email, token)
	if err != nil {
		fmt.Fprintln(w, "Error al guardar en la base de datos:", err)
		return
	}

	fmt.Println("Usuario autenticado y guardado:", user)
	http.Redirect(w, r, "http://localhost:8080/", http.StatusFound)
}

func (s *Server) logout(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")

	r = r.WithContext(context.WithValue(context.Background(), "provider", provider))

	gothic.Logout(w, r)
	http.Redirect(w, r, "http://localhost:8080/health", http.StatusTemporaryRedirect)
}

func (s *Server) getAuth(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")

	r = r.WithContext(context.WithValue(context.Background(), "provider", provider))

	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		gothic.BeginAuthHandler(w, r)
		return
	}

	fmt.Println("test", user)
	http.Redirect(w, r, "http://localhost:8080/", http.StatusFound)
}
