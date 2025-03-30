package handlers

import (
	"html/template"
	"log/slog"
	"net/http"
)

func renderTemplate(w http.ResponseWriter, tmpl string, data any) {
	tmplPath := "templates/" + tmpl + ".html"
	tmpls, err := template.ParseFiles("templates/base.html", tmplPath)

	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		slog.Debug(err.Error())
		return
	}
	if err := tmpls.ExecuteTemplate(w, "base.html", data); err != nil {
		slog.Debug(err.Error())
	}
}

func RootHandler(w http.ResponseWriter, r *http.Request) {

	data := map[string]any{
		"Title": "dummy",
	}

	slog.Debug("hit")
	renderTemplate(w, "index", data)
}
