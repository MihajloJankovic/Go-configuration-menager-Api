package main

import (
	"errors"
	"log"
	"mime"
	"net/http"

	"github.com/gorilla/mux"
)

type postServer struct {
	dataConfig map[string]*Config
	data map[string]*ConfigGroup
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
	ts.dataConfig[id] = rt
	renderJSON(w, rt)
}


func (ts *postServer) getAllHandler(w http.ResponseWriter, req *http.Request) {
	allCons := []*Config{}
	for _, v := range ts.dataConfig {
		allCons = append(allCons, v)
	}

	renderJSON(w, allCons)
}

func (ts *postServer) getPostHandler(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	config, ok := ts.dataConfig[id]
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

	delete(ts.dataConfig, id)
	renderJSON(w, ts.dataConfig[0])

}



func (ts *postServer) createConfigGroupHandler(w http.ResponseWriter, req *http.Request) {
    contentType := req.Header.Get("Content-Type")
    mediatype, _, err := mime.ParseMediaType(contentType)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    if mediatype != "application/json" {
        err := errors.New("Expected application/json Content-type")
        http.Error(w, err.Error(), http.StatusUnsupportedMediaType)
        return
    }


    rt, err := decodeGroupBody(req.Body)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    id := createId()
    rt.Id = id
    ts.data[id] = rt
    renderJSON(w, rt)

    log.Println(rt.Id)
}

func (ts *postServer) getAllConfigGroupHandlers(w http.ResponseWriter, req *http.Request) {
    allTasks := []*Group{}
    for _, v := range ts.data {
        allTasks = append(allTasks, v)
    }
    renderJSON(w, allTasks)
}

func (ts *postServer) getConfigGroupHandler(w http.ResponseWriter, req *http.Request) {
    id := mux.Vars(req)["id"]
    task, ok := ts.data[id]
    if !ok {
        err := errors.New("id not found")
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }
    renderJSON(w, task)
}

func (ts *postServer) delConfigGroupHandler(w http.ResponseWriter, req *http.Request) {
    id := mux.Vars(req)["id"]
    if v, ok := ts.data[id]; ok {
        delete(ts.data, id)
        renderJSON(w, v)
    } else {
        err := errors.New("id not found")
        http.Error(w, err.Error(), http.StatusNotFound)
    }
}

func (ts *postServer) addConfigInConfigGroup(w http.ResponseWriter, req *http.Request){
	id := req.URL.Query().Get("id")
	config := req.URL.Query().Get("config")
	configGroup := req.URL.Query().Get("configGroup")

	configW := ts.dataConfig[config]
	configGroupW := ts.data[configGroup]

	configGroupW.Configs[config] = configW

	renderJSON(w, configGroupW)
}
