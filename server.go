package main

import (
	"errors"
	ps "github.com/MihajloJankovic/Alati/Dao"
	pss "github.com/MihajloJankovic/Alati/Dao2"
	"github.com/gorilla/mux"
	"mime"
	"net/http"
)

type postServer struct {
	Dao  *ps.Dao
	Dao2 *pss.Dao2
}

// swagger:route POST /config/ config createConfig
// Add new config
//
// responses:
//
//	415: ErrorResponse
//	400: ErrorResponse
//	201: ResponsePost
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
	config, err := ts.Dao.Create(rt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	renderJSON(w, config)

}

// swagger:route GET /configs/ config getConfigs
// Get all configs
//
// responses:
//
//	200: []ResponsePost
func (ts *postServer) getAllHandler(w http.ResponseWriter, req *http.Request) {
	configs, err := ts.Dao.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	renderJSON(w, configs)
}

// swagger:route GET /config/{id}/{version}/ config getConfigById
// Get config by ID
//
// responses:
//
//	404: ErrorResponse
//	200: ResponsePost
func (ts *postServer) getPostHandler(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	version := mux.Vars(req)["version"]
	config, err := ts.Dao.Get(id, version)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	renderJSON(w, config)
}

// swagger:route DELETE /config/{id}/{version}/ config deleteConfig
// Delete config
//
// responses:
//
//	404: ErrorResponse
//	204: NoContentResponse
func (ts *postServer) delPostHandler(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	version := mux.Vars(req)["version"]
	msg, err := ts.Dao.Delete(id, version)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	renderJSON(w, msg)

}
func (ts *postServer) getPostByLabel(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	version := mux.Vars(req)["version"]
	labels := mux.Vars(req)["labels"]

	task, err := ts.Dao.GetPostsByLabels(id, version, labels)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	renderJSON(w, task)
}

// swagger:route POST /configGroup/ configGroup createConfigGroup
// Add new config group
//
// responses:
//
//	415: ErrorResponse
//	400: ErrorResponse
//	201: ResponsePost
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

	group, err := ts.Dao2.CreateGroup(rt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	renderJSON(w, group)
}

// swagger:route GET /configGroups/ configGroup getConfigGroups
// Get all config groups
//
// responses:
//
//	200: []ResponsePost
func (ts *postServer) getAllConfigGroupHandlers(w http.ResponseWriter, req *http.Request) {
	allTasks, err := ts.Dao2.GetAllGroups()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	renderJSON(w, allTasks)
}

// swagger:route GET /configGroup/{id}/ configGroup getConfigGroupById
// Get config group by ID
//
// responses:
//
//	404: ErrorResponse
//	200: ResponsePost
func (ts *postServer) getConfigGroupHandler(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	task, err := ts.Dao2.GetGroup(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	renderJSON(w, task)
}

// swagger:route DELETE /configGroup/{id}/ configGroup deleteConfigGroup
// Delete config group
//
// responses:
//
//	404: ErrorResponse
//	204: NoContentResponse
func (ts *postServer) delConfigGroupHandler(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	version := mux.Vars(req)["version"]
	msg, err := ts.Dao2.DeleteGroup(id, version)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	renderJSON(w, msg)
}

// swagger:route GET /configGroup/{id}/{version}/{configGroup}/ config, configGroup addConfigInConfigGroup
// Add config in config group
//
// responses:
//
//	415: ErrorResponse
//	400: ErrorResponse
//	201: ResponsePost
func (ts *postServer) addConfigInConfigGroup(w http.ResponseWriter, req *http.Request) {

	//configGroup := mux.Vars(req)["configGroup"]
	//	id := mux.Vars(req)["id"]
	//	version := mux.Vars(req)["version"]

}

func (ts *postServer) swaggerHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./swagger.yaml")
}
