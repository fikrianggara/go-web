package main

import (
	"encoding/json"
	f "fmt"
	"net/http"
)

func handlerIndex(w http.ResponseWriter, r *http.Request) {

	var fikri Developer = Person{"Fikri", 22, []string{"coding", "sports", "music"}}
	f.Println(fikri.Code() + "->" + fikri.Build() + "->" + fikri.Test() + "->" + fikri.Deploy())
	jsonInBytes, err := json.Marshal(fikri.(Person))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonInBytes)
}

func main() {
	// routing
	http.HandleFunc("/", handlerIndex)
	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		var message = "hello from '/index' endpoint"
		w.Write([]byte(message))
	})

	var address = ":9000"
	server := new(http.Server)
	server.Addr = address
	f.Printf("server started at %s", address)
	var err = server.ListenAndServe()

	if err != nil {
		f.Println("error occured", err.Error())
	}
}
