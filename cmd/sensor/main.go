package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/ra9dev/safe-and-sound/internal/configs"
	"github.com/ra9dev/safe-and-sound/pkg/config"
	"github.com/ra9dev/safe-and-sound/pkg/log"
	"math/rand"
	"net/http"

	stdlog "log"
)

func main() {
	appConfig := config.Parse(new(configs.SensorConfig)).(*configs.SensorConfig)
	log.Setup(appConfig.LogMode)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, r, map[string]interface{}{
			"id":         appConfig.ID,
			"soundLevel": rand.Intn(100),
		})
	})

	stdlog.Printf("Running HTTP server on port %s", appConfig.ListenAddr)
	if err := http.ListenAndServe(appConfig.ListenAddr, r); err != nil {
		stdlog.Printf("[ERROR] HTTP server failed: %v", err)
		return
	}
}
