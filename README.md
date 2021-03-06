# Todos with go, oto, sqlite and svelte

> Oto is a go driven rpc code generation tool, we are using this to generate the go endpoints and the javascript to use those endpoints.

## Install && Run

1. Create a database:
   ```
   touch todos.db
   ```
1. Install the frontend (svelte):
   ```
   cd frontend
   npm install && npm run build
   ```
1. Run it:
   ```
   go run main.go
   ```
1. Navigate in your browser to [localhost:8080](http://localhost:8080)

## Generating code

```
# backend
oto -template ./templates/server.go.plush \
    -out ./generated/oto.gen.go \
    -ignore Ignorer \
    -pkg generated \
    ./definitions/definitions.go &&
gofmt -w ./generated/oto.gen.go ./generated/oto.gen.go

# frontend
oto -template ./templates/client.js.plush \
    -out ./frontend/src/oto.gen.js \
    -ignore Ignorer \
    ./definitions/definitions.go
```

## Dependencies

- [go](https://golang.org)
- [oto](https://github.com/pacedotdev/oto)
