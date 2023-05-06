package handlers

import (
	"html/template"
	"net/http"

	"github.com/hunterjsb/sqlbase/internal/session"
)

// replace the following with your actual user authentication logic
func authenticateUser(username, password string) (userID string, success bool) {
	// Dummy authentication logic
	if username == "admin" && password == "admin" {
		return "1", true
	}
	return "", false
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl := template.Must(template.ParseFiles("web/templates/login.html"))
		tmpl.Execute(w, nil)
	} else if r.Method == "POST" {
		username := r.FormValue("username")
		password := r.FormValue("password")

		userID, authenticated := authenticateUser(username, password)

		if authenticated {
			session.CreateSession(w, userID)
			http.Redirect(w, r, "/sqlbase", http.StatusFound)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			tmpl := template.Must(template.ParseFiles("web/templates/login.html"))
			tmpl.Execute(w, map[string]interface{}{
				"Error": "Invalid username or password",
			})
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
