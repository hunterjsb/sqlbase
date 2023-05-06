package main

import (
	"fmt"
	"github.com/hunterjsb/sqlbase/internal/redis"

	// "github.com/hunterjsb/sqlbase/internal/database"
	// "github.com/hunterjsb/sqlbase/internal/redis"
	// "github.com/hunterjsb/sqlbase/pkg/config"
	"github.com/hunterjsb/sqlbase/router"
	"log"
	"net/http"
)

func main() {
	// Load configuration
	//config.Load()

	// Initialize database and Redis connections
	//database.InitDB()
	//defer database.CloseDB()

	// Initialize Redis client
	redis.InitRedisClient("localhost:6379", "", 0)

	// Set up routes for the web pages and API
	webMux := router.SetupWebRoutes()
	apiMux := router.SetupAPIRoutes()

	// Combine both routers into a single http.ServeMux
	mux := http.NewServeMux()
	mux.Handle("/", webMux)
	mux.Handle("/api/", apiMux)

	// Start the web server
	// port := config.Get().Port
	port := "8080"
	fmt.Printf("Starting server on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
