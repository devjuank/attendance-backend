package middleware

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/juank/attendance-backend/config"
)

func CORSMiddleware(cfg *config.Config) gin.HandlerFunc {
	config := cors.DefaultConfig()

	if len(cfg.CORS.AllowedOrigins) > 0 {
		config.AllowOrigins = cfg.CORS.AllowedOrigins
	} else {
		config.AllowAllOrigins = true
	}

	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	config.AllowCredentials = true
	config.MaxAge = 12 * time.Hour

	return cors.New(config)
}
