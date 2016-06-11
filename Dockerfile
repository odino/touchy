FROM golang

RUN go get github.com/gorilla/mux
RUN go get github.com/goji/httpauth
RUN go get github.com/micmonay/keybd_event
RUN go get github.com/mitchellh/gox

WORKDIR /go/src/github.com/odino/touchy
CMD gox
