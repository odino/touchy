<img align="right" width="250px" src="https://raw.githubusercontent.com/odino/touchy/master/touchy.gif" />

# Touchy

> Remote control for your laptop / desktop, over HTTP.

Touchy provides a remote keyboard for your workstation:
It is perfect for presentations / conferencing, to be able
to control your laptop directly from your phone, without
having to install any additional software.

## Usage

[Download the binary](https://github.com/odino/touchy/releases) and launch it, you should see something like:

```
~/projects/go/src/github.com/odino/touchy (master ✘)✹ ᐅ ./touchy
Aye, here we are: connect to http://192.168.0.112:8080 and enjoy!
```

Now, assuming that your phone is connected to the same network as your
laptop, open `http://192.168.0.112:8080` and have fun!

By default, `touchy` will run on port `8080`, but you can customize that
with the `HTTP_PORT` environment variable.

## Authentication

You can even protect the server with basic HTTP authentication by
setting the environment variables `HTTP_USER` and `HTTP_PASSWD`.

## License

[MIT](https://opensource.org/licenses/MIT)

> Basically: do as you wish, you're on your own.

## Contributing

> Prerequisites:
>
> * go 1.4 or above
> * babel installed globally
> * webpack installed globally

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

For now I just mapped 4 buttons (up / down / left / right) and the interface
looks pretty crappy, but I'm open to adding /  changing stuff, if requested:
pull requests are always welcome!

## Why?

I bought one of those Logitech presenters and it seemed very cool -- until
my second conference using the controller, as I forgot it was turned on
and the batteries got drained. Thus, I started looking for a solution
that would involve less stuff to remember (laziness is a virtue).

At the same time, I wanted to refresh my rusty Go skills, and wanted
to try react + babel + webpack on one of my OS projects.

So, with a weekend in front of you and a patient wife, what would
you do? :-)

## Tests & known issues

I want to add some very simple tests for the Go server: there isn't much
I wrote myself, so I'll be honest and admit that tests aren't the #1 priority
as of now.

There's currently an issue that prevents me from publishing [some binaries](https://github.com/micmonay/keybd_event/issues/2).
