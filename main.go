package main

import (
	"log"
	"net/http"
	"text/template"

	"gopkg.in/mgo.v2"
)

type EntriesHandler struct {
	entries []Entry
}

var index = template.Must(template.ParseFiles(
	"templates/_base.html",
	"templates/index.html",
))

func (e *EntriesHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	s := session.Clone()
	defer s.Close()

	coll := s.DB("gostbook").C("entries")
	query := coll.Find(nil).Sort("-timestamp")

	if err := query.All(&e.entries); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println(e.entries)

	if err := index.Execute(w, e.entries); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

var session *mgo.Session

func main() {
	var err error
	session, err = mgo.Dial("localhost")
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/entries", &EntriesHandler{})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
