package main

import (
	"errors"
	"log"
	"mime"
	"net/http"

	"github.com/gorilla/mux"
)

type postServer struct {
	dataConfig map[string]map[string]*Config
	data       map[string]*ConfigGroup
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
	if len(rt.Id) == 0 {
		rt.Id = id
		m := map[string]*Config{}
		m[rt.Version] = rt
		ts.dataConfig[id] = m
		renderJSON(w, rt)

	} else {
		n := ts.dataConfig[rt.Id]
		n[rt.Version] = rt
		renderJSON(w, rt)
	}

}

func (ts *postServer) getAllHandler(w http.ResponseWriter, req *http.Request) {
	allCons := []*Config{}

	for i, _ := range ts.dataConfig {
		for _, b := range ts.dataConfig[i] {
			allCons = append(allCons, b)
		}
	}

	renderJSON(w, allCons)
}

func (ts *postServer) getPostHandler(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	version := mux.Vars(req)["version"]
	mapa, ok := ts.dataConfig[id]
	config, oke := mapa[version]
	if !ok {
		err := errors.New("key not found")
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	if !oke {
		err := errors.New("version not found")
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	renderJSON(w, config)
}

func (ts *postServer) delPostHandler(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	version := mux.Vars(req)["version"]

	mapa := ts.dataConfig[id]

	if len(mapa) > 1 {
		delete(mapa, version)
	} else {
		delete(ts.dataConfig, id)
	}
	renderJSON(w, ts.dataConfig[id])

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
	allTasks := []*ConfigGroup{}
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

func (ts *postServer) addConfigInConfigGroup(w http.ResponseWriter, req *http.Request) {

	configGroup := mux.Vars(req)["configGroup"]
	id := mux.Vars(req)["id"]
	version := mux.Vars(req)["version"]
	mapa, mm := ts.dataConfig[id]
	config, oke := mapa[version]
	if !mm {
		err := errors.New("key not found")
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	if !oke {
		err := errors.New("version not found")
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	if configGroupW, ok := ts.data[configGroup]; ok {
		if configGroupW.Configs == nil {
			configGroupW.Configs = make(map[string]map[string]*Config)

		}

		if configGroupW.Configs[id] == nil {

			m := map[string]*Config{}
			m[version] = config
			configGroupW.Configs[id] = m
		} else {
			z := configGroupW.Configs[id]
			z[version] = config
		}

		renderJSON(w, configGroupW)
	} else {
		err := errors.New("ConfigGroup id not found")
		http.Error(w, err.Error(), http.StatusNotFound)
	}

}
