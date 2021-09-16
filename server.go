package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"text/template"

	"goder/controllers"
)

func redirectTLS(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://"+r.Host+r.RequestURI, http.StatusMovedPermanently)
}

func main() {

	go http.ListenAndServe(":80", http.HandlerFunc(redirectTLS))
	mux := http.NewServeMux()

	t, err := template.ParseFiles("./static/index.html")
	if err != nil {
		log.Fatal(err)
	}
	styleFile, err := os.Open("./static/style.goder")
	if err != nil {
		log.Fatal(err)
	}

	darkFile, err := os.Open("./static/dark-theme.goder")
	if err != nil {
		log.Fatal(err)
	}
	dark, err := ioutil.ReadAll(darkFile)
	style, err := ioutil.ReadAll(styleFile)
	if err != nil {
		log.Fatal(err)
	}

	fs := http.FileServer(http.Dir("./static"))

	mux.Handle("/static/", http.StripPrefix("/static", fs))

	mux.HandleFunc("/", controllers.RenderHandlerHOF(t, string(style), string(dark)))

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
	// 	if err := http.ListenAndServe(":80", http.HandlerFunc(redirectTLS)); err != nil {
	// 		log.Fatal(err)
	// }
}
