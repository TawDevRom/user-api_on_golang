package router

import (
	"net/http"
	"user-api/handlers"
	"user-api/middleware"
)

func SetupRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/register", handlers.RegisterUser)
	mux.HandleFunc("/login", handlers.Login)

	mux.HandleFunc("/users", middleware.AuthMiddleware(handlers.GetAllUsers))
	return mux
}

// JWT (JSON Web Token)
