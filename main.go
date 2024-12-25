package main

import (
	"GinSvelteEmbed/internal/counter"
	"GinSvelteEmbed/internal/debug"
	"GinSvelteEmbed/web"
	"github.com/gin-gonic/gin"
	"io/fs"
	"log"
	"net/http"
	"strings"
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

// setupAPIRoutes configures the API endpoints
func setupAPIRoutes(r *gin.Engine, cnt *counter.Counter) {
	api := r.Group("/api")
	{
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
