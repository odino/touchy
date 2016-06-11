<img align="right" width="250px" src="https://raw.githubusercontent.com/odino/touchy/master/bin/images/logo.png?token=AAUC5MTk6aOK1OBS04HB9VUkQyvtudRqks5XZH9swA%3D%3D" />

# Touchy

> Remote control for your laptop / desktop, over HTTP.

Touchy provides a remote keyboard for your workstation:
it is perfect for presentations / conferencing to be able
to control your laptop directly from your phone, without
having to install any additional software.

## Usage

TBD

## Authentication

TBD

## License

TBD

## Contributing

As a dev, you might want to run `touchy` locally. Just clone this repo
and then run `make`:

```
~/projects/go/src/github.com/odino/touchy (master ✘)✚ ᐅ make
Installing go packages...
go get -u github.com/jteeuwen/go-bindata/...
go get github.com/gorilla/mux
go get github.com/goji/httpauth
go get github.com/micmonay/keybd_event
Compiling static assets...
$GOPATH/bin/go-bindata -pkg http -o http/bindata.go static/...
Formatting source code...
go fmt ./...
http/bindata.go
Building...
go build -o touchy main.go
Running the server...
Aye, here we are: connect to http://192.168.0.112:8080 and enjoy!
```

## Why?

TBD

## Tests

TBD
