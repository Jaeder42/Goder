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

	fs := http.FileServer(http.Dir("./static"))

	http.Handle("/static/", http.StripPrefix("/static", fs))

	http.HandleFunc("/", controllers.RenderHandlerHOF(t))

	fmt.Println("Starting server at port 443")
	if err := http.ListenAndServeTLS(":443", "/etc/letsencrypt/live/jaeder42.tech/fullchain.pem", "/etc/letsencrypt/live/jaeder42.tech/privkey.pem", nil); err != nil {
		fmt.Println("Server could not start on port 443, attempting 80")
		if err := http.ListenAndServe(":80", nil); err != nil {
			fmt.Println("Server could not start on 80, falling back to 8080")
			if err := http.ListenAndServe(":8080", nil); err != nil {
				log.Fatal(err)
			}

		}
	}
}
