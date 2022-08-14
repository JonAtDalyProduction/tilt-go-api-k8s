package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"net/http"
	"os"
)

type AppConfig struct {
	HttpPort string `env:"HTTP_PORT,required=false"`
}

var config AppConfig

type App struct {
	config *AppConfig
}

var app App

func init() {
	//envloader.ParseEnv(&config, "dev.env")
	config.HttpPort = "3333"

}
func (a *App) getRoot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	res := make(map[string]string)
	res["message"] = "Go API is Working"
	jsonRes, err := json.Marshal(res)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	_, err = w.Write(jsonRes)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

func (a *App) postRoot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	res := make(map[string]string)
	res["message"] = "Got post message from frontend"
	jsonRes, err := json.Marshal(res)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	_, err = w.Write(jsonRes)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

func main() {
	app.config = &config
	r := chi.NewRouter()
	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://localhost", "http://api.localhost", "http://localhost:5173", "http://localhost:3333"},
		//AllowOriginFunc:    nil,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
		//OptionsPassthrough: false,
		//	Debug:              false,
	}))
	//Base route
	r.Get("/", app.getRoot)
	r.Post("/", app.postRoot)
	fmt.Printf("starting api server on port %s\n", config.HttpPort)
	err := http.ListenAndServe(fmt.Sprintf(":%s", config.HttpPort), r)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
