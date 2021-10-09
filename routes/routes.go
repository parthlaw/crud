package routes

import (
	"net/http"

	"crud/controllers"

	"github.com/gorilla/mux"
)

func Handlers() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(http.StatusAccepted)
		rw.Write([]byte("Hello World"))
	}).Methods("GET")
	r.HandleFunc("/create", controllers.Create).Methods("POST")
	r.HandleFunc("/read", controllers.Read).Methods("GET")
	r.HandleFunc("/update", controllers.Update).Methods("PUT")
	r.HandleFunc("/delete", controllers.Delete).Methods("DELETE")
	return r
}
