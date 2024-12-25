
# For SvelteKit - Frontend
Replace "svelte.config.js" to use **[adapter-static](https://svelte.dev/docs/kit/adapter-static)**.
```js
import adapter from '@sveltejs/adapter-static';
import { vitePreprocess } from '@sveltejs/vite-plugin-svelte';

/** @type {import('@sveltejs/kit').Config} */
const config = {
	preprocess: vitePreprocess(),
	kit: {
		adapter: adapter({
			pages: 'web/build',
			assets: 'web/build',
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

Disable server-side rendering in `src/routes/+layout.(js|ts)`:
```js
export const ssr = false;
export const prerender = false;
```

Create **web.go** in the root directory to embed the web build:
```go
package web

import "embed"

//go:embed build/*
var Fs embed.FS
```

Optionally, you can add /api route to vite.config.js to proxy the backend **in development**:
```js
import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vite';

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

# For Go - Backend

Install gin "github.com/gin-gonic/gin"
```sh
go get -u github.com/gin-gonic/gin
```

