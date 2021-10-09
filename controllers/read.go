package controllers

import (
	"crud/models"
	"encoding/json"
	"net/http"

	"github.com/kamva/mgm/v3"
)

func Read(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "No id provided", http.StatusBadRequest)
		return
	}
	user := &models.User{}
	coll := mgm.Coll(user)
	_ = coll.FindByID(id, user)
	if user == nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
