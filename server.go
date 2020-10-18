package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"./controllers"
)

func main() {
	t, err := template.ParseFiles("./static/index.html")

	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", controllers.RenderHandlerHOF(t))

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
