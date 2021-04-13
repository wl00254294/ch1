package service

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	message "eric.com/go/ch1/msg"
	"github.com/gorilla/mux"
)

type CustomRoute struct {
	Name       string `json:"name"`
	Method     string `json:"method"`
	Path       string `json:"path"`
	HandleFunc string `json:"handlefunc"`
}

//set handler map
var functions = map[string]http.HandlerFunc{
	"sayhelloName": sayhelloName,
	"getToken":     getToken,
}

func readConfig() []CustomRoute {

	data, err := ioutil.ReadFile("./config.json")
	if err != nil {

		log.Printf("Error:%s, %s", message.Content["configreaderror"].Code, err)
	}

	var obj []CustomRoute

	// unmarshall it
	err = json.Unmarshal(data, &obj)
	if err != nil {
		log.Printf("Error:%s, %s", message.Content["jsonconverror"].Code, err)
	}

	return obj
}

func ServerOn() {

	var obj = readConfig()

	srv := mux.NewRouter()

	for _, dt := range obj {

		if functions[dt.HandleFunc] == nil {
			log.Printf("Error:%s, %s", message.Content["configdataerror"].Code, dt.HandleFunc)
		} else {
			srv.HandleFunc(dt.Path, functions[dt.HandleFunc]).Methods(dt.Method)
			log.Printf("Success: %s add Ok!\n", dt.HandleFunc)
		}

	}

	fs := http.FileServer(http.Dir("./swaggerui"))
	srv.PathPrefix("/swaggerui/").Handler(http.StripPrefix("/swaggerui/", fs))
	http.ListenAndServe(":8080", srv)

}
