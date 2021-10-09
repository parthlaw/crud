package controllers

import (
	"crud/models"
	"crud/utils"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kamva/mgm/v3"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var requestData utils.CreateRequest
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	user := models.NewUser(requestData.Name, requestData.Dob, requestData.Address, requestData.Description)
	err = mgm.Coll(user).Create(user)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
