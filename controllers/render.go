package controllers

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"text/template"
)

// RenderHandler renders static web page
func RenderHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	body := "<h1>Oh oh</h1><h2>We couldn't find that page</h2>"
	title := "¯\\_(ツ)_/¯"

	file, err := os.Open("pages" + r.URL.Path + ".md")
	md, err := ioutil.ReadAll(file)
	if err == nil {
		body = ""
		arr := strings.Split(string(md), "\n")
		for _, element := range arr {
			if strings.HasPrefix(element, "# ") {
				body += strings.Replace(element, "# ", "<h1>", 1) + "</h1>"
				title = strings.Replace(element, "# ", "", 1)
			} else if strings.HasPrefix(element, "## ") {
				body += strings.Replace(element, "## ", "<h2>", 1) + "</h2>"
			} else {
				body += "<p>" + element + "</p>"
			}

		}
	}

	t, err := template.ParseFiles("static/index.html")

	if err != nil {
		http.Error(w, "Something went wrong.", http.StatusInternalServerError)
		return
	}

	var tpl bytes.Buffer
	data := struct {
		Body  string
		Title string
	}{
		Title: title,
		Body:  body,
	}
	if err := t.Execute(&tpl, data); err != nil {
		http.Error(w, "Something went wrong.", http.StatusInternalServerError)
		return
	}
	result := tpl.String()

	fmt.Fprintf(w, result)
}
