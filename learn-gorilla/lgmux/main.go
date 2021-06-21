package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

// HomeHandler ...
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Home: %v\n", vars["category"])
}

func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Products: %v\n", vars["category"])
}

func ArticlesHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Articles: %v\n", vars["category"])
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/products", ProductsHandler)
	r.HandleFunc("/articles", ArticlesHandler)
	http.Handle("/", r)

	r := mux.NewRouter()
	r.Host("www.example.com")
	r.Host("{subdomain:[a-z]+}.example.com")
	r.PathPrefix("/products/")
	r.Methods("GET", "POST")
	r.Schemes("https")
	r.Headers("X-Requested-With", "XMLHttpRequest")
	r.Queries("key", "value")
	r.MatcherFunc(func(r *http.Request, rm *RouteMatch) bool {
		return r.ProtoMajor == 0
	})

	r.HandleFunc("/products", ProductsHandler).Host("www.example.com").Methods("GET").Schemes("http")

	r := mux.NewRouter()
	r.HandleFunc("/specific", ProductsHandler)
	r.PathPrefix("/").Handler(catchAllHandler)

	r := mux.NewRouter()
	s := r.Host("www.example.com").Subrouter()
	s.HandleFunc("/products/", ProductsHandler)
	s.HandleFunc("/products/{key}", ProductsHandler)
	s.HandleFunc("/articles/{category}/{id:[0-9]+}", ArticlesHandler)

	r := mux.NewRouter()
	s := r.PathPrefix("/products").Subrouter()
	s.HandleFunc("/", ProductsHandler)
	s.HandleFunc("/{key}/", ProductsHandler)

	// static files
	var dir string
	flag.StringVar(&dir, "dir", ".", "the directory to serve files from. Defaults to the current dir")
	flag.Parse()
	r := mux.NewRouter()

	r.PathPrefix("/static").Handler(http.StripPrefix("/static", http.FileServer(http.Dir(dir))))
	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
