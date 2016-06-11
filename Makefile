scratch: install build run
install:
	@echo "Installing go packages..."
	go get -u github.com/jteeuwen/go-bindata/...
	go get github.com/codegangsta/gin
	go get github.com/gorilla/mux
	go get github.com/goji/httpauth
	go get github.com/micmonay/keybd_event
	@echo "Installing frontend dependencies..."
	cd static && npm install
build: frontend
	@echo "Compiling static assets..."
	$$GOPATH/bin/go-bindata -pkg http -o http/bindata.go static/bundle.js static/index.html
	@echo "Formatting source code..."
	go fmt ./...
	@echo "Building..."
	go build -o touchy main.go
frontend:
	@echo "Compiling frontend with babel + webpack..."
	cd static && babel index.js > compiled.js && webpack ./compiled.js bundle.js
run:
	@echo "Running the server..."
	@$$GOPATH/bin/gin -a 8080 main.go -t .
release: build
	@echo "Building docker image to cross-compile touchy..."
	docker build -t touchy .
	@echo "Compiling touchy for multiple archs..."
	docker run -ti -v $$(pwd):/go/src/github.com/odino/touchy touchy gox
