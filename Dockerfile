FROM golang
ADD . /go/src/github.com/redbo/giphaux
RUN go get -d -v 'github.com/redbo/giphaux/...'
RUN go install 'github.com/redbo/giphaux/...'
ENTRYPOINT giphaux
