package main

import (
	"calhoun/controllers"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	_, _ = fmt.Fprint(w, "<h1>Page not found</h1>")
}

func main() {
	staticControllers := controllers.NewStatic()
	usersController := controllers.NewUsers()

	fmt.Println("Live on http://localhost:3000/")

	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(notFound)
	r.Handle("/", staticControllers.Home).Methods("GET")
	r.Handle("/contact", staticControllers.Contact).Methods("GET")
	r.HandleFunc("/signup", usersController.New).Methods("GET")
	r.HandleFunc("/signup", usersController.Create).Methods("POST")
	_ = http.ListenAndServe(":3000", r)
}
