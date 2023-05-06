package handlers

import (
	"html/template"
	"net/http"
)

func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("web/templates/dashboard.html"))
	tmpl.Execute(w, nil)
}
