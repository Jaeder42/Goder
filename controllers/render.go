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

func formatLink(link string) string {
	return strings.Replace(strings.ToLower(link), " ", "-", 1)
}

// RenderHandlerHOF returns a function that renders static web page
func RenderHandlerHOF(t *template.Template) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "GET" {
			http.Error(w, "Method is not supported.", http.StatusNotFound)
			return
		}

		body := "<h1>Oh oh</h1><h2>We couldn't find that page</h2>"
		title := "¯\\_(ツ)_/¯"
		path := strings.ToLower(r.URL.Path)
		if path == "/" {
			path = "/index"
		}
		file, err := os.Open("pages" + path + ".md")
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
				} else if strings.HasPrefix(element, "### ") {
					body += strings.Replace(element, "### ", "<h3>", 1) + "</h3>"
				} else if strings.HasPrefix(element, "~ ") {
					body += strings.Replace(element, "~ ", "<a class='bloglink' href='"+"/devLog/"+formatLink(strings.Replace(element, "~ ", "", 1))+"'>", 1) + "</a>"
				} else if len(element) > 0 {
					body += "<p>" + element + "</p>"
				}

			}
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
}
