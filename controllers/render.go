package controllers

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"
)

func formatLink(link string) string {
	return strings.ReplaceAll(strings.ToLower(link), " ", "-")
}

// RenderHandlerHOF returns a function that renders static web page
func RenderHandlerHOF(t *template.Template, style string, darktheme string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "GET" {
			http.Error(w, "Method is not supported.", http.StatusNotFound)
			return
		}

		styleString := style

		themeCookie, err := r.Cookie("theme")
		if err != nil {
			log.Println(err)
		} else if themeCookie.Value == "dark" {
			fmt.Println(themeCookie.Value, themeCookie.Value == "dark")
			styleString = darktheme + style
		} else {
			styleString = style
		}

		body := "<h1>Oh oh</h1><h2>We couldn't find that page</h2>"
		title := "¯\\_(ツ)_/¯"
		description := "Jæder42"
		image := "preview.png"
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
					// Refactor to use count for <hx>
					body += strings.Replace(element, "# ", "<h1>", 1) + "</h1>"
					title = strings.Replace(element, "# ", "", 1)
				} else if strings.HasPrefix(element, "## ") {
					body += strings.Replace(element, "## ", "<h2>", 1) + "</h2>"
					description = strings.Replace(element, "## ", "", 1)
				} else if strings.HasPrefix(element, "### ") {
					body += strings.Replace(element, "### ", "<h3>", 1) + "</h3>"
				} else if strings.HasPrefix(element, "% ") {
					body += strings.Replace(element, "% ", "<img src='/static/images/", 1) + "'/>"
					image = strings.Replace(element, "% ", "", 1)
				} else if strings.HasPrefix(element, "~ ") {
					body += strings.Replace(element, "~ ", "<div class='bloglink-container'><a class='bloglink' href='"+"/devLog/"+formatLink(strings.Replace(element, "~ ", "", 1))+"'>", 1) + "</a></div>"
				} else if len(element) > 0 {
					body += "<p>" + element + "</p>"
				}

			}
		}

		var tpl bytes.Buffer
		data := struct {
			Body        string
			Title       string
			Description string
			Image       string
			URL         string
			Style       string
		}{
			Title:       title,
			Body:        body,
			Description: description,
			Image:       image,
			URL:         "https://jaeder42.tech" + path,
			Style:       styleString,
		}
		if err := t.Execute(&tpl, data); err != nil {
			http.Error(w, "Something went wrong.", http.StatusInternalServerError)
			return
		}
		result := tpl.String()
		log.Println("Rendering " + path)
		fmt.Fprintf(w, result)
	}
}
