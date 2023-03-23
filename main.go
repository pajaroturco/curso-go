package main

import (
	"fmt"
	"log"
	"net/http"

	"pajaro.com/curso-go/sistema"
)

func paginaNF(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Pagina no encontrada")
	http.NotFound(w, r)
}

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
	http.HandleFunc("/page", paginaNF)
	fmt.Println("El servidor esta corriendo en el puerto " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}