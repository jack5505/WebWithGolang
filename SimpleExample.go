package main

import (
	"fmt"
	"net/http"
	"time"
)

// declare constant variables here
const (
	port = ":8080"
)

//requests here and action after each request
func serverDynamic(w http.ResponseWriter, r *http.Request) {
	response := "this time is " + time.Now().String()
	fmt.Fprintln(w, response)
}

//requests here and action after each request
func serverStatic(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static.html")
}
func main() {
	//using http package and HandleFun method you can declare path and requests function here
	// this second parameter type of function should be executed after first url requests
	http.HandleFunc("/static", serverStatic)
	http.HandleFunc("/dynamic", serverDynamic)
	// here is running server with following port :8080
	http.ListenAndServe(port, nil)
}
