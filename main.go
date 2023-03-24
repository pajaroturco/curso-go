package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"pajaro.com/curso-go/handlers"
	"pajaro.com/curso-go/models"
	"pajaro.com/curso-go/sistema"
)

func Index(w http.ResponseWriter, r *http.Request) {
	template, error := template.ParseFiles("templates/index.html")
	if error != nil {
		panic(error)
	} else {
		template.Execute(w, nil)
	}
}

func main() {
	// cargamos puerto del archivo .env
	port := sistema.GoDotEnvVariable("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	models.MigrarUser()

	mux := mux.NewRouter()

	mux.HandleFunc("/", Index).Methods("GET")
	mux.HandleFunc("/api/users", handlers.GetUsers).Methods("GET")
	mux.HandleFunc("/api/user/{id:[0-9+]}", handlers.GetUser).Methods("GET")
	mux.HandleFunc("/api/users", handlers.CreateUser).Methods("POST")
	mux.HandleFunc("/api/user/{id:[0-9+]}", handlers.UpdateUser).Methods("PUT")
	mux.HandleFunc("/api/user/{id:[0-9+]}", handlers.DeleteUser).Methods("DELETE")
	

	server := &http.Server{	Addr: ":" + port, Handler: mux}
	fmt.Println("El servidor esta corriendo en el puerto " + port)
	log.Fatal(server.ListenAndServe())
}