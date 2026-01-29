# Support for web deployments
To create a new `main.wasm` file, you can run `GOOS=js GOARCH=wasm go build -o fishing/web/main.wasm main.go` in the git root drectory.
To start a simple server to listen on `localhost:8000` use `python3 -m http.server 8000`