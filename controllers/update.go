package controllers

import (
	"crud/models"
	"crud/utils"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kamva/mgm/v3"
)

func Update(w http.ResponseWriter, r *http.Request) {
	var data utils.CreateRequest
	id := r.URL.Query().Get("id")
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil || id == "" {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
	}
	user := &models.User{}
	coll := mgm.Coll(user)
	_ = coll.FindByID(id, user)
	if data.Address != "" {
		user.Address = data.Address
	}
	if data.Description != "" {
		user.Description = data.Description
	}
	err = coll.Update(user)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
