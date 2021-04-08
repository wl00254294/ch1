package service

import (
	"fmt"
	"net/http"
	"strings"

	jwtctrl "eric.com/go/ch1/pk3"
)

func getToken(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("Addr", r.RemoteAddr)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	token := jwtctrl.GenJWToken(36000, 1, "eric", r.RemoteAddr)

	fmt.Fprintf(w, "tokem:"+token)
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	token := ""
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))

		if k == "token" {
			token = strings.Join(v, "")
		}
	}

	if token == "" {
		fmt.Fprintf(w, "token parameter should add!")
		return
	} else {
		fmt.Println(token)
	}
	ret, err := jwtctrl.ParseToken(token, r.RemoteAddr)
	if err != nil {
		fmt.Fprintf(w, "token is invaild!")
		return
	}

	fmt.Fprintf(w, "Hello !"+ret.Issuer)
}
