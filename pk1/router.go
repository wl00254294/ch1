package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

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
}

func readConfig() []CustomRoute {

	data, err := ioutil.ReadFile("./config.json")
	if err != nil {
		fmt.Print(err)
	}

	var obj []CustomRoute

	// unmarshall it
	err = json.Unmarshal(data, &obj)
	if err != nil {
		fmt.Println("error:", err)
	}

	return obj
}

func ServerOn() {

	var obj = readConfig()

	srv := mux.NewRouter()

	for _, dt := range obj {

		if functions[dt.HandleFunc] == nil {
			fmt.Printf("Error: %s not found\n", dt.HandleFunc)
		} else {
			srv.HandleFunc(dt.Path, functions[dt.HandleFunc]).Methods(dt.Method)
			fmt.Printf("Success: %s add Ok!\n", dt.HandleFunc)
		}

	}

	http.ListenAndServe(":8080", srv)

}
