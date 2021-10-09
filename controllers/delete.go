package controllers

import (
	"crud/models"
	"fmt"
	"net/http"

	"github.com/kamva/mgm/v3"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	// Get the id from the URL
	id := r.URL.Query().Get("id")
	user := &models.User{}
	coll := mgm.Coll(user)
	_ = coll.FindByID(id, user)
	err := mgm.Coll(user).Delete(user)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
