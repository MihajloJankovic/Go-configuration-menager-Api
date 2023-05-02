package main

import (
    "errors"
    "github.com/gorilla/mux"
    "log"
    "mime"
    "net/http"
)

type postServerGroup struct {
    data map[string]*ConfigGroup
}

func (ts *postServerGroup) createGroupHandler(w http.ResponseWriter, req *http.Request) {
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

func (ts *postServerGroup) getAllGroupHandler(w http.ResponseWriter, req *http.Request) {
    allTasks := []*Group{}
    for _, v := range ts.data {
        allTasks = append(allTasks, v)
    }
    renderJSON(w, allTasks)
}

func (ts *postServerGroup) getGroupHandler(w http.ResponseWriter, req *http.Request) {
    id := mux.Vars(req)["id"]
    task, ok := ts.data[id]
    if !ok {
        err := errors.New("id not found")
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }
    renderJSON(w, task)
}

func (ts *postServerGroup) delGroupHandler(w http.ResponseWriter, req *http.Request) {
    id := mux.Vars(req)["id"]
    if v, ok := ts.data[id]; ok {
        delete(ts.data, id)
        renderJSON(w, v)
    } else {
        err := errors.New("id not found")
        http.Error(w, err.Error(), http.StatusNotFound)
    }
}