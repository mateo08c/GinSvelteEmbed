package main

import (
	"GinSvelteEmbed/internal/counter"
	"GinSvelteEmbed/internal/debug"
	"GinSvelteEmbed/web"
	ratelimit "github.com/JGLTechnologies/gin-rate-limit"
	"github.com/gin-gonic/gin"
	"io/fs"
	"log"
	"net/http"
	"strings"
	"time"
)

func main() {
	// Initialize the counter with a persistent file
	cnt, err := counter.NewCounter("counter.txt")
	if err != nil {
		log.Fatal(err)
	}

	f := getFileSystem()

	//print fs content for debugging
	debug.PrintFSContent(f, 1)

	// Create a Gin engine
	r := gin.Default()

	// Set up API routes
	setupAPIRoutes(r, cnt)

	// Serve static files and fallback for SPA
	r.Use(spaFileHandler(f))

	// Run the server on port 8080
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

func keyFunc(c *gin.Context) string {
	return c.ClientIP()
}

func errorHandler(c *gin.Context, info ratelimit.Info) {
	c.JSON(http.StatusTooManyRequests, gin.H{"error": "Rate limit exceeded, reset in " + time.Until(info.ResetTime).String()})
}

// setupAPIRoutes configures the API endpoints
func setupAPIRoutes(r *gin.Engine, cnt *counter.Counter) {
	// Middleware to limit the rate of requests

	//je veux max 10request par minute et reset après 30 secondes
	store := ratelimit.InMemoryStore(&ratelimit.InMemoryOptions{
		Rate:  1 * time.Minute, // 1 minute
		Limit: 10,              // max 10 requests per minute
	})

	mw := ratelimit.RateLimiter(store, &ratelimit.Options{
		ErrorHandler: errorHandler,
		KeyFunc:      keyFunc,
	})

	api := r.Group("/api")
	{
		api.Use(mw)
		api.GET("/count", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"count": cnt.Value()})
		})

		api.POST("/increment", func(c *gin.Context) {
			if err := cnt.Increment(); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{"count": cnt.Value()})
		})

		api.POST("/decrement", func(c *gin.Context) {
			if err := cnt.Decrement(); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{"count": cnt.Value()})
		})
	}
}

// spaFileHandler serves static files and handles fallback for SPA routing
func spaFileHandler(httpFS http.FileSystem) gin.HandlerFunc {
	fileServer := http.FileServer(httpFS)
	return func(c *gin.Context) {
		path := c.Request.URL.Path

		// Bypass the handler for API requests
		if strings.HasPrefix(path, "/api") {
			c.Next()
			return
		}

		// Try to serve the exact file
		if _, err := httpFS.Open(path); err == nil {
			fileServer.ServeHTTP(c.Writer, c.Request)
			return
		}

		// Fallback to serving index.html for SPA
		c.Request.URL.Path = "/"
		fileServer.ServeHTTP(c.Writer, c.Request)
		c.Request.URL.Path = path // Restore the original path
	}
}

// getFileSystem retrieves the embedded filesystem for serving static files
func getFileSystem() http.FileSystem {
	fSys, err := fs.Sub(web.Fs, "build") // Extract the "build" directory from the embedded filesystem
	if err != nil {
		panic(err)
	}
	return http.FS(fSys)
}
