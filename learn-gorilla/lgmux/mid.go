package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/micro/micro/monitor/handler"
)

// loggingMiddleware ...
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

// handler ...
func handler(w http.ResponseWriter, r *http.Request) {
	return
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handler)
	r.Use(loggingMiddleware)
	amw := authenticationMiddleware{}
	amw.Populate()
	r.Use(amw.Middleware)

	// CORS
	r.Use(mux.CORSMethodMiddleware(r))

	r.HandleFunc("/health", HealthCheckHandlern)
}

type authenticationMiddleware struct {
	tokenUsers map[string]string
}

// Populate
func (amw *authenticationMiddleware) Populate() {
	amw.tokenUsers["000000000"] = "user0"
	amw.tokenUsers["aaaaaaaaa"] = "userA"
	amw.tokenUsers["05f"] = "randomUser"
	amw.tokenUsers["deadbeef"] = "user0"
}

// Middleware ...
func (amw *authenticationMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("X-Session-Token")

		if user, found := amw.tokenUsers[token]; found {
			log.Printf("Authenticated user %s\n", user)
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "Forbidden", http.StatusForbidden)
		}
	})
}

// HealthCheckHandlern
func HealthCheckHandlern(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	io.WriteString(w, `{"alive": true}`)
}
