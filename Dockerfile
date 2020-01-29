FROM golang
ADD . /go/src/github.com/golang/redbo/giphaux
RUN go install github.com/redbo/giphaux/...
ENTRYPOINT giphaux
