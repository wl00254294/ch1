package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	dao "eric.com/go/ch1/dao"
	jwtctrl "eric.com/go/ch1/pk3"
)

func getToken(w http.ResponseWriter, r *http.Request) {

	var p dao.User
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token := jwtctrl.GenJWToken(3600, p)
	js, err := json.Marshal(token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

}

func sayhelloName(w http.ResponseWriter, r *http.Request) {

	auth := r.Header.Get("Authorization")

	token := strings.Split(auth, "Bearer ")[1]
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))

	}

	ret, err := jwtctrl.ParseToken(token)
	if err != nil {
		fmt.Fprintf(w, "token is invaild!")
		return
	}

	fmt.Fprintf(w, "Hello !"+ret.Issuer)
}
