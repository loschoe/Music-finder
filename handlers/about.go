package handlers

import (
	"html/template"
	"net/http"
	"path/filepath"
)

func About(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(filepath.Join("templates", "AboutUS.html")))
	tmpl.Execute(w, nil)
}
