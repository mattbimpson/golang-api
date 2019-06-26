package main

import (
	"log"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MoviesDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "movies"
)

// GetMovies returns all movies
func GetMovies() string {
	return "all movies"
}

// Connect creates a connection with the db
func (m *MoviesDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}

	db = session.DB(m.Database)
}

// FindAll gets all movies from the db
func (m *MoviesDAO) FindAll() ([]Movie, error) {
	var movies []Movie
	err := db.C(COLLECTION).Find(bson.M{}).All(&movies)
	return movies, err
}

// Insert inserts a single movie into the db
func (m *MoviesDAO) Insert() {
	var movie Movie
	movie.Description = "test"
	movie.ID = "99"
	movie.Name = "a test movie"
	//verr := db.C(COLLECTION).Insert(&movie)
	//vreturn err
	db.C(COLLECTION).Insert(&movie)
}
