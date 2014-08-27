package main

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Entry struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	Timestamp time.Time
	Name      string
	Message   string
}

func NewEntry() *Entry {
	return &Entry{
		Timestamp: time.Now(),
	}
}
