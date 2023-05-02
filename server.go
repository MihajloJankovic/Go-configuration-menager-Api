package main

import (
	"errors"
	"log"
	"mime"
	"net/http"

	"github.com/gorilla/mux"
)

type postServer struct {
	data map[string]*Config
}

func (ts *postServer) createPostHandler(w http.ResponseWriter, req *http.Request) {
	contentType := req.Header.Get("Content-Type")
	mediatype, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if mediatype != "application/json" {
		err := errors.New("Expect application/json Content-Type")
		http.Error(w, err.Error(), http.StatusUnsupportedMediaType)
		return
	}

	rt, err := decodeBody(req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusAccepted)
		return
	}

	id := createId()
	rt.Id = id
	ts.data[id] = rt
	renderJSON(w, rt)
}

func (ts *postServer) getAllHandler(w http.ResponseWriter, req *http.Request) {
	allCons := []*Config{}
	for _, v := range ts.data {
		allCons = append(allCons, v)
	}

	renderJSON(w, allCons)
}

func (ts *postServer) getPostHandler(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	config, ok := ts.data[id]
	if !ok {
		err := errors.New("key not found")
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	renderJSON(w, config)
}

func (ts *postServer) delPostHandler(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	log.Println(id)

	delete(ts.data, id)
	renderJSON(w, ts.data[0])

}
