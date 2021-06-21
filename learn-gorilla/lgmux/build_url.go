package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

// hello ...
func hello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello world")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/articles/{category}/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Hello world")
	}).Name("article")

	url, err := r.Get("article").URL("category", "technology", "id", "42")
	fmt.Printf("%+v %T %+v\n", url, url, err)

	r = mux.NewRouter()
	r.Host("{subdomain}.example.com").Path("/articles/{category}/{id:[0-9]+}").Queries("filter", "{filter}").HandlerFunc(hello).Name("hello")

	url, err = r.Get("hello").URL("subdomain", "news", "category", "technology", "id", "42", "filter", "gorilla")
	fmt.Printf("%+v %T %+v\n", url, url, err)

	// r.HeadersRegexp("Content-Type", "application/(text|json)")

	host, err := r.Get("hello").URLHost("subdomain", "news")
	path, err := r.Get("hello").URLPath("category", "technology", "id", "42")
	fmt.Printf("%+v %+v %+v", host, path, err)
}
