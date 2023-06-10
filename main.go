package main

import (
	"encoding/json"
	f "fmt"
	"html/template"
	"net/http"
	"net/url"
	"strconv"
)

func handlerIndex(w http.ResponseWriter, r *http.Request) {

	var fikri Developer = Person{"Fikri", 22, []string{"coding", "sports", "music"}}
	// f.Println(fikri.Code() + "->" + fikri.Build() + "->" + fikri.Test() + "->" + fikri.Deploy())
	jsonInBytes, err := json.Marshal(fikri.(Person))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonInBytes)
}

func handlerIntroduction(w http.ResponseWriter, r *http.Request) {
	var fikri Developer = Person{"Fikri", 22, []string{"coding", "sports", "music"}}

	var t, err = template.ParseFiles("introduction.html")
	if err != nil {
		f.Println("error occured", err.Error())
		return
	}

	t.Execute(w, fikri)
}

func handlerIntroWithParams(w http.ResponseWriter, r *http.Request) {

	var urlString = r.URL.RequestURI()
	var u, err = url.Parse(urlString)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	type DataRes struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	name := u.Query()["name"][0]
	var age string = u.Query()["age"][0]
	ageInt, err := strconv.Atoi(age)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := DataRes{name, ageInt}
	f.Println(name, age)
	jsonInBytes, err := json.Marshal(data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonInBytes)
}

func main() {
	// routing
	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		var message = "hello from '/index' endpoint"
		w.Write([]byte(message))
	})
	http.HandleFunc("/introduction", handlerIntroduction)
	http.HandleFunc("/named-introduction", handlerIntroWithParams)
	http.HandleFunc("/", handlerIndex)

	// initiate server
	var address = ":9000"
	server := new(http.Server)
	server.Addr = address
	f.Printf("server started at %s", address)
	var err = server.ListenAndServe()

	if err != nil {
		f.Println("error occured", err.Error())
	}
}
