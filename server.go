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

	fmt.Printf("Starting server at port 443\n")
	if err := http.ListenAndServeTLS(":443", "/etc/letsencrypt/live/jaeder42.tech/fullchain.pem", "/etc/letsencrypt/live/jaeder42.tech/privkey.pem", nil); err != nil {
		log.Fatal(err)
	}
}
