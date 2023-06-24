package controller

import (
	"encoding/json"
	"go-web/model"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {

	var fikri model.Developer = model.Person{Name: "Fikri", Age: 22, Hobby: []string{"coding", "sports", "music"}}
	// f.Println(fikri.Code() + "->" + fikri.Build() + "->" + fikri.Test() + "->" + fikri.Deploy())
	jsonInBytes, err := json.Marshal(fikri.(model.Person))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonInBytes)
}
