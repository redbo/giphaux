package main

import (
	"database/sql"
	"fmt"
	"math/rand"
)

type dataStore interface {
	GifById(id string) (GIF, error)
	GifsById(ids []string) []GIF
	Search(query string, limit int, offset int, rating string, lang string) ([]GIF, int)
	Trending(limit int, off int, rating string) []GIF
	RandomSearch(q string, weirdness int) (GIF, error)
	RandomByTag(q string, rating string) (GIF, error)
	RandomID() string
	AddGif(username string, path string, tags []string, sourceURL string)
}

type sqlDataStore struct {
	db *sql.DB
}

func (s *sqlDataStore) RandomID() string {
	return fmt.Sprintf("%016x%016x", rand.Int63(), rand.Int63())
}
