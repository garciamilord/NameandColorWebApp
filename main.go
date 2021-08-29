package main

import (
	"html/template"
	"net/http"
)

type GData struct {
	Title string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)
	http.HandleFunc("/process", process)
	http.Handle("/assets/",
		http.StripPrefix("/assets", http.FileServer(http.Dir("./assets"))))

	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	gd := GData{
		Title: "Tell us about yourself",
	}
	tpl.ExecuteTemplate(w, "submit.gohtml", gd)
}

func about(w http.ResponseWriter, r *http.Request) {
	gd := GData{
		Title: "About",
	}
	tpl.ExecuteTemplate(w, "about.gohtml", gd)
}

func contact(w http.ResponseWriter, r *http.Request) {
	gd := GData{
		Title: "Contact",
	}
	tpl.ExecuteTemplate(w, "contact.gohtml", gd)
}

func process(w http.ResponseWriter, r *http.Request) {

	name := r.FormValue("name")
	color := r.FormValue("color")

	d := struct {
		GData
		Name  string
		Color string
	}{
		GData: GData{
			Title: "PROCESS",
		},
		Name:  name,
		Color: color,
	}
	if name != "" && (color != "" && color == "red" || color == "blue" || color == "yellow" || color == "green") {
		tpl.ExecuteTemplate(w, "process.gohtml", d)
	} else {
		index(w, r)
	}

}
