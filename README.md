# giphaux

Giphaux is a minimal clone of Giphy written in go, including a web interface and a rough implementation of the [Giphy API](https://developers.giphy.com/docs/api#quick-start-guide).

### Docker Start (untested)

```sh
$ docker build -t giphaux .
$ docker run --publish 8080:8080 --name gphx1 --rm giphaux
```

### Local Quick Start

First have a [go development environment](https://golang.org/doc/install) configured.  Once that is set up, you can install giphaux.

```sh
$ go install github.com/redbo/giphaux/...
# ... some time will pass ...
```
Then you can use giphaux to create a `~/.giphaux` directory containing a default configuration file and an empty database.
```sh
$ giphaux init
# ... spammy sql stuff ...
```
Then you can launch the giphaux server, which listens on `localhost:8080` by default.
```sh
$ giphaux
# ... logs ...
```

### Tech

Giphaux uses the following tools and libraries in addition to the go standard lib.

* gorilla mux - an http router
* gorm - an ORM for go
* go-sqlite3 - a sqlite driver for go
* httpsnoop - an http response recorder
* zap - uber's logging library
* go-bindata - to package templates for distribution.

