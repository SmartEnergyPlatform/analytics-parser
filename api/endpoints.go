package api

import (
	"net/http"
	"encoding/json"
	"parsing-service/lib"
	"github.com/gorilla/mux"
	"parsing-service/parser"
	"fmt"
)

type Endpoint struct {
	flowApiService lib.FlowApiService
	flowParser *parser.FlowParser
}

func NewEndpoint(flowApiService lib.FlowApiService, operator_api lib.OperatorApiService) *Endpoint{
	ret := parser.NewFlowParser(flowApiService, operator_api)
	return &Endpoint{flowApiService ,ret}
}

func (e *Endpoint) getRootEndpoint(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(lib.Response{"OK"})
}

func (e *Endpoint) getParseFlow(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	ret := e.flowParser.ParseFlow(vars["id"], e.getUserId(req))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(ret)
}

func (e *Endpoint) getGetInputs(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	ret := e.flowParser.GetInputs(vars["id"], e.getUserId(req))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(ret)
}

func (e *Endpoint) getUserId(req *http.Request) (userId string){
	userId = req.Header.Get("X-UserId")
	if userId == "" {
		userId = "admin"
	}
	fmt.Println("UserID: " + userId)
	return
}