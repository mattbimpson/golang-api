package main

import (
	"gopkg.in/mgo.v2/bson"
)

type Movie struct {
	ID        	 bson.ObjectId 		`json:"_id,omitempty" bson:"_id,omitempty"`
	Name 		 string             `json:"name,omitempty" bson:"name,omitempty"`
	Description  string             `json:"description,omitempty" bson:"description,omitempty"`
}