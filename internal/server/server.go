package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"

	"lt-connect/internal/database"
)

type Server struct {
	port int

	db database.Service
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))

	// this is light wrapper around the std lib's http server
	// So that it has our ports and db connection
	myServer := &Server{
		port: port,

		db: database.New(),
	}

	// Declare Server config
	// This is the actually engine: The standard lib's http server
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", myServer.port),
		Handler:      myServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
