package main

import (
	"fmt"
	"net/http"
)

func handlerIndex(w http.ResponseWriter, r *http.Request) {
	var message = "hello from  '/' endpoint"
	w.Write([]byte(message))
}

func handlerHello(w http.ResponseWriter, r *http.Request) {
	var message = "hello world"
	w.Write([]byte(message))
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var message = "hello from '/' endpoint"
		w.Write([]byte(message))
	})
	http.HandleFunc("/index", handlerIndex)
	http.HandleFunc("/hello", handlerHello)

	var address = "localhost:9000"
	fmt.Printf("server started at %s", address)
	var err = http.ListenAndServe(address, nil)

	if err != nil {
		fmt.Println("error occured", err.Error())
	}
}
