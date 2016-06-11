scratch: install build run
install:
	@echo "Installing go packages..."
	go get -u github.com/jteeuwen/go-bindata/...
	go get github.com/gorilla/mux
	go get github.com/goji/httpauth
	go get github.com/micmonay/keybd_event
build:
	@echo "Compiling static assets..."
	$$GOPATH/bin/go-bindata -pkg http -o http/bindata.go static/...
	@echo "Formatting source code..."
	go fmt ./...
	@echo "Building..."
	go build -o touchy main.go
run:
	@echo "Running the server..."
	@./touchy
