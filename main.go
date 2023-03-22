package main

import (
	"fmt"
	"log"
	"net/http"

	"pajaro.com/curso-go/sistema"
)


func main() {
	// cargamos puerto del archivo .env
	port := sistema.GoDotEnvVariable("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}
	// router
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hola Mundo")
	})
	fmt.Println("El servidor esta corriendo en el puerto 3000")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}