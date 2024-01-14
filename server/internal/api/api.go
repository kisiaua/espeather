package api

import (
	"ESPeather/internal/db"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

func StartServer() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
	}))

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})

	router.Get("/readings/{tableName}", func(w http.ResponseWriter, r *http.Request) {
		params := r.URL.Query()

		tableName := chi.URLParam(r, "tableName")
		readingsCount := params.Get("count")
		if readingsCount == "" {
			readingsCount = "20"
		}

		readings := db.ReadDB(tableName, readingsCount)
		jsonData, err := json.Marshal(readings)
		if err != nil {
			log.Println("Error encoding JSON:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
