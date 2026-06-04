package middleware

import (
	"os"
	"strings"

	"github.com/QuantumNous/new-api/common"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// CORS returns a CORS middleware that supports configuration via the CORS_ORIGINS
// environment variable. When CORS_ORIGINS is unset or empty, a safe default list
// of origins is used (including the production frontend and local dev server).
// When set, only the specified origins are permitted. This enables secure
// cross-origin access for separated frontend deployments.
func CORS() gin.HandlerFunc {
	config := cors.DefaultConfig()

	corsOrigins := os.Getenv("CORS_ORIGINS")
	var origins []string
	if corsOrigins != "" {
		// Parse comma-separated origin list, trim whitespace
		for _, o := range strings.Split(corsOrigins, ",") {
			o = strings.TrimSpace(o)
			if o != "" {
				origins = append(origins, o)
			}
		}
	} else {
		// Default origins for separated frontend deployment
		origins = []string{
			"https://app.cinatoken.com",
			"http://localhost:5173",
		}
	}

	if len(origins) == 0 {
		config.AllowAllOrigins = true
	} else {
		config.AllowOrigins = origins
		config.AllowAllOrigins = false
	}

	config.AllowCredentials = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"}
	config.AllowHeaders = []string{"*"}
	// Expose headers needed for SSE/streaming responses (e.g. Cache-Control,
	// X-Accel-Buffering) so browsers can read them in cross-origin requests.
	config.ExposeHeaders = []string{
		"Content-Length",
		"Cache-Control",
		"X-Accel-Buffering",
		"X-Request-Id",
	}

	return cors.New(config)
}

func PoweredBy() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("X-New-Api-Version", common.Version)
		c.Next()
	}
}
