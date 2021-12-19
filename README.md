## http/basicHandler.go

A simple api for getting random numbers using `http` built-in multiplexer

## custom/customMux.go

A simple api for getting random number using a custom mux for multiplexing
HTTP routes (URLs) to different handlers.

## httprouter/execService.go

This api executes command to show go version and files content

## httprouter-static/fileServer.go

A static file server

## gorilla/muxRouter.go

Handle request parameters using gorilla mux router

## gorilla-static/fileServer.go

A static file server with gorilla mux

## gorilla-query/muxQuery.go

Handle query parameters with gorilla mux

## gorilla-short-url/muxRouter.go

API interface:

| URL | REST Verb | Action | Success | Failure |
| --- | --- | --- | --- | --- |
| /api/v1/new | POST | Create a shortened URL | 200 | 404 |
| /api/v1/:url | GET | Given the shortened URL, return the original URL | 200 | 404 |

curl commands to test:
```
curl -X POST http://localhost:8000/api/v1/new -H 'Content-Type: application/json' -d '{"url":"www.yahoo.com"}'
```
```
curl -X GET http://localhost:8000/api/v1/t92Yu82boFWec5u7uD
```