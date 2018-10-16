package api

import (
"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"github.com/rs/cors"
	"parsing-service/lib"
	"parsing-service/flows-api"
	"parsing-service/operator-api"
)

func CreateServer(){
	f := flows_api.NewFlowApi(
		lib.GetEnv("FLOW_API_ENDPOINT", ""),
	)
	o := operator_api.NewOperatorApi(lib.GetEnv("OPERATOR_API_ENDPOINT", ""))
	port := lib.GetEnv("API_PORT", "8000")
	fmt.Print("Starting Server at port " + port + "\n")
	router := mux.NewRouter()

	e := NewEndpoint(f, o)
	router.HandleFunc("/", e.getRootEndpoint).Methods("GET")
	router.HandleFunc("/pipe/{id}", e.getParseFlow).Methods("GET")
	router.HandleFunc("/pipe/getinputs/{id}", e.getGetInputs).Methods("GET")
	handler := cors.New(cors.Options{
		AllowedHeaders: []string{"Content-Type", "Authorization"},
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "DELETE", "OPTIONS"},
	}).Handler(router)
	logger := lib.NewLogger(handler, "CALL")
	log.Fatal(http.ListenAndServe(":"+ port, logger))
}
