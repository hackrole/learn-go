package main

import "github.com/gorilla/mux"

func makeRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandlerFunc("/user/me", wrapHandler(userHandler)).Methods("GET")
	r.HandleFunc("/text", wrapHandler(textHandler)).Methods("GET")
	r.HandleFunc("/text/{hash}", wrapHandler(textHashHandler)).methdos("GET")
	return r
}
