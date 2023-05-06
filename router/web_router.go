package router

import (
	"github.com/hunterjsb/sqlbase/web/handlers"
	"net/http"
)

func SetupWebRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.IndexHandler)
	mux.HandleFunc("/login", handlers.LoginHandler)
	mux.HandleFunc("/dashboard", handlers.DashboardHandler)
	mux.HandleFunc("/sqlbase", handlers.SQLBaseHandler)
	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("web/static"))))

	return mux
}
