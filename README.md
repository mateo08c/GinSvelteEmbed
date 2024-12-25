# 🚀 SvelteKit + Go: The Ultimate Full-Stack Boilerplate

Supercharge your web development with this powerful combination of SvelteKit for the frontend and Go (with Gin) for the
backend. This boilerplate provides a seamless integration between these two cutting-edge technologies, allowing you to
build scalable and performant web applications with ease.

## ✨ Features

- 🎨 **SvelteKit**: A powerful frontend framework for building blazing-fast web apps
- 🚄 **Go + Gin**: High-performance backend with one of the fastest web frameworks for Go
- 🔗 **Seamless Integration**: Pre-configured setup for smooth communication between frontend and backend
- 🛠 **Development Ready**: Includes proxy configuration for easy local development
- 📦 **Production Optimized**: Static adapter configuration for SvelteKit, ready for deployment

## 🖥 Frontend Configuration (SvelteKit)

### Static Adapter

Update `svelte.config.js`:

```js
import adapter from '@sveltejs/adapter-static';
import {vitePreprocess} from '@sveltejs/vite-plugin-svelte';

/** @type {import('@sveltejs/kit').Config} */
const config = {
    preprocess: vitePreprocess(),
    kit: {
        adapter: adapter({
            pages: 'build',
            assets: 'build',
            fallback: 'index.html',
            precompress: false
        }),
        prerender: {
            handleHttpError: 'ignore'
        }
    }
};

export default config;
```

### Disable SSR

In `src/routes/+layout.(js|ts)`:

```js
export const ssr = false;
export const prerender = false;
```

### Go Integration

Create `web.go` in the root:

```go
package web

import "embed"

//go:embed build/*
var Fs embed.FS
```

## 🖥 Backend Configuration (Go)

### Install Gin

First, install the Gin web framework:

```sh
go get -u github.com/gin-gonic/gin
```

### Example Usage

Check out `main.go` in the backend directory for an example of how to set up a basic Gin server and integrate it with
the SvelteKit frontend. Here's a snippet of what you might find:

```go
package main

import (
	"github.com/gin-gonic/gin"
	"io/fs"
	"log"
	"net/http"
	"package-name/web"
	"strings"
)

func main() {
	f := getFileSystem()

	// Create a Gin engine
	r := gin.Default()

	// Define an API route
	api := r.Group("/api")
	api.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, world!"})
	})

	// Serve static files and fallback for SPA
	r.Use(spaFileHandler(f))

	// Run the server on port 8080
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

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

```

## 🛠 Development

For local development, update `vite.config.js` to proxy API requests:

```js
import {sveltekit} from '@sveltejs/kit/vite';
import {defineConfig} from 'vite';

export default defineConfig({
    plugins: [sveltekit()],
    server: {
        proxy: {
            "/api": {
                target: "http://127.0.0.1:8080",
                changeOrigin: true,
                secure: false,
                ws: true,
            },
        },
    },
});
```

## ⚠️ Important Note

Before running the Go server, make sure to build the SvelteKit frontend or create empty directories for the static files:

```sh
npm run build
```

## 📚 Learn More

- [SvelteKit Documentation](https://kit.svelte.dev/docs)
- [Go Documentation](https://golang.org/doc/)
- [Gin Web Framework](https://gin-gonic.com/docs/)

## 🤝 Contributing

Contributions, issues, and feature requests are welcome! Feel free to check [issues page](link-to-issues).

---

Happy coding! 🎉