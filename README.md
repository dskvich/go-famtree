# Golang + Svelte/Snowpack simple app

A relatively lightweight simple app:
* Server API using Golang
* UI is built with Svelte + Snowpack

CI: https://go-famtree.herokuapp.com

## Running in Docker

`docker-compose up --build`

or to just start the DB:
`docker-compose up -d db`

This will bind to `127.0.0.1:65432` by default

# Development

After clone:

```
npm install
```

Then:

```
npm run watch
```

To start the server:

```
go run .\cmd\main.go
```