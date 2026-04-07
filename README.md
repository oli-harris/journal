# E2E Encrypted Journal App

## Local development

Run backend in API-only dev mode (Go API on port 3000):

```sh
cd backend
go run . -dev
```

Run frontend (Vite/Svelte on port 5173 with /api proxied to Go backend):

```sh
cd frontend
pnpm dev
```

Open http://localhost:5173.

In dev mode, the backend serves only /api routes. The frontend is served by Vite.

## Build single executable

Build the frontend into backend/web and compile Go:

```sh
cd frontend
pnpm build

cd ../backend
go build -o journal-server .
```

Then run:

```sh
./journal-server
```

Open http://localhost:3000.

In prod mode (without -dev), the backend serves both /api and embedded frontend files.

## Docker

Build image:

```sh
docker build -t journal:latest .
```

Run container:

```sh
docker run --rm -p 3000:3000 journal:latest
```
