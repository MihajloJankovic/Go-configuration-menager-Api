package main

import (
	"context"
	"errors"
	"fmt"
	ps "github.com/MihajloJankovic/Alati/Dao"
	pss "github.com/MihajloJankovic/Alati/Dao2"
	//tracer "github.com/MihajloJankovic/Alati/tracer"
	"github.com/gorilla/mux"
	opentracing "github.com/opentracing/opentracing-go"
	"io"
	"mime"
	"net/http"
)

type postServer struct {
	keys  map[string]string
	keys2 map[string]string

	Dao    *ps.Dao
	Dao2   *pss.Dao2
	tracer opentracing.Tracer
	closer io.Closer
}

// swagger:route POST /config/ config createConfig
// Add new config
//
// responses:
//
//	415: ErrorResponse
//	400: ErrorResponse
//	201: ResponsePost
func (ts *postServer) createPostHandler(ctx context.Context, w http.ResponseWriter, req *http.Request) {
	contentType := req.Header.Get("Content-Type")
	key := req.Header.Get("key")
	mediatype, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	v := ts.keys[key]
	if v == "" {
		ts.keys[key] = key

	} else {
		http.Error(w, "Already Created", http.StatusAccepted)
		return
	}
	if mediatype != "application/json" {
		err := errors.New("Expect application/json Content-Type")
		http.Error(w, err.Error(), http.StatusUnsupportedMediaType)
		return
	}

	rt, err := decodeBody(ctx, req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusAccepted)
		return
	}
	config, err := ts.Dao.Create(ctx, rt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	renderJSON(ctx, w, config)

}

// swagger:route GET /configs/ config getConfigs
// Get all configs
//
// responses:
//
//	200: []ResponsePost
func (ts *postServer) getAllHandler(ctx context.Context, w http.ResponseWriter, req *http.Request) {
	configs, err := ts.Dao.GetAll(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	renderJSON(ctx, w, configs)
}

// swagger:route GET /config/{id}/{version}/ config getConfigById
// Get config by ID
//
// responses:
//
//	404: ErrorResponse
//	200: ResponsePost
func (ts *postServer) getPostHandler(ctx context.Context, w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	version := mux.Vars(req)["version"]
	config, err := ts.Dao.Get(ctx, id, version)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	renderJSON(ctx, w, config)
}

// swagger:route DELETE /config/{id}/{version}/ config deleteConfig
// Delete config
//
// responses:
//
//	404: ErrorResponse
//	204: NoContentResponse
func (ts *postServer) delPostHandler(ctx context.Context, w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	version := mux.Vars(req)["version"]
	msg, err := ts.Dao.Delete(ctx, id, version)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	renderJSON(ctx, w, msg)

}
func (ts *postServer) getPostByLabel(ctx context.Context, w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	version := mux.Vars(req)["version"]
	labels := mux.Vars(req)["labels"]

	task, err := ts.Dao.GetPostsByLabels(ctx, id, version, labels)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	renderJSON(ctx, w, task)
}

// swagger:route POST /configGroup/ configGroup createConfigGroup
// Add new config group
//
// responses:
//
//	415: ErrorResponse
//	400: ErrorResponse
//	201: ResponsePost
func (ts *postServer) createConfigGroupHandler(ctx context.Context, w http.ResponseWriter, req *http.Request) {
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
	key := req.Header.Get("key")
	v := ts.keys2[key]
	if v == "" {
		ts.keys2[key] = key

	} else {
		http.Error(w, "Already Created", http.StatusAccepted)
		return
	}
	rt, err := decodeGroupBody(ctx, req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	group, err := ts.Dao2.CreateGroup(ctx, rt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	renderJSON(ctx, w, group)
}

// swagger:route GET /configGroups/ configGroup getConfigGroups
// Get all config groups
//
// responses:
//
//	200: []ResponsePost
func (ts *postServer) getAllConfigGroupHandlers(ctx context.Context, w http.ResponseWriter, req *http.Request) {
	allTasks, err := ts.Dao2.GetAllGroups(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	renderJSON(ctx, w, allTasks)
}

// swagger:route GET /configGroup/{id}/ configGroup getConfigGroupById
// Get config group by ID
//
// responses:
//
//	404: ErrorResponse
//	200: ResponsePost
func (ts *postServer) getConfigGroupHandler(ctx context.Context, w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	task, err := ts.Dao2.GetGroup(ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	renderJSON(ctx, w, task)
}

// swagger:route DELETE /configGroup/{id}/ configGroup deleteConfigGroup
// Delete config group
//
// responses:
//
//	404: ErrorResponse
//	204: NoContentResponse
func (ts *postServer) delConfigGroupHandler(ctx context.Context, w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	version := mux.Vars(req)["version"]
	msg, err := ts.Dao2.DeleteGroup(ctx, id, version)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	renderJSON(ctx, w, msg)
}

// swagger:route GET /configGroup/{id}/{version}/{configGroup}/ config, configGroup addConfigInConfigGroup
// Add config in config group
//
// responses:
//
//	415: ErrorResponse
//	400: ErrorResponse
//	201: ResponsePost
func (ts *postServer) addConfigInConfigGroup(ctx context.Context, w http.ResponseWriter, req *http.Request) {

	id := mux.Vars(req)["id"]
	version := mux.Vars(req)["version"]
	configGroup := mux.Vars(req)["configGroup"]

	config, err := ts.Dao.Get(ctx, id, version)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	task, err := ts.Dao2.GetGroup(ctx, configGroup)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	task.Configs = append(task.Configs, config)
	fmt.Println(task.Configs)
	grupas, err := ts.Dao2.SaveGroup(ctx, task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	renderJSON(ctx, w, grupas)

}

func (ts *postServer) swaggerHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./swagger.yaml")
}
