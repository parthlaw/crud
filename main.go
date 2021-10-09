package main

import (
	"crud/db"
	"crud/routes"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/rs/cors"
)

func main() {
	db.Init()
	fmt.Println("Hello, World!")
	mux := http.NewServeMux()
	mux.Handle("/", routes.Handlers())
	c := cors.Default()
	handler := c.Handler(mux)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println("Listening on port " + port)
	e := http.ListenAndServe(":"+port, handler)
	if e != nil {
		log.Fatal(e)
	}
}
