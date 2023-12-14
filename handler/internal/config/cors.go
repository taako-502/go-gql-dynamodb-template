package config

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	rscors "github.com/rs/cors"
)

// CORS
func SettingCors(hosts []string) gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins: hosts,
		AllowMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodPut,
			http.MethodOptions,
		},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	})
}

// CORS(For Local Server)
func SettingCrosForLocalServer() *rscors.Cors {
	return rscors.New(rscors.Options{
		AllowedOrigins: []string{
			"http://localhost:3333",
			"http://localhost:8080",
		},
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodPut,
			http.MethodOptions,
		},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
		Debug:            true,
	})
}
