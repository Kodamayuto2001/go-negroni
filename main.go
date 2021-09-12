package main

import (
	"fmt"
	"net/http"

	"github.com/urfave/negroni"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello, middleware!")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("This is middleware test!")
}

func middleware1(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	fmt.Println("[START] middleware1")
	next(w, r)
	fmt.Println("[END] middleware1")
}

func middleware2(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	fmt.Println("[START] middleware2")
	next(w, r)
	fmt.Println("[END] middleware2")
}

func middleware3(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	fmt.Println("[START] middleware3")
	next(w, r)
	fmt.Println("[END] middleware3")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/about", aboutHandler)

	n := negroni.New()
	n.Use(negroni.HandlerFunc(middleware1))
	n.Use(negroni.HandlerFunc(middleware2))
	n.Use(negroni.HandlerFunc(middleware3))
	n.UseHandler(mux)
	http.ListenAndServe(":3000", n)
}