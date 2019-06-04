package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func pageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pageId := vars["id"]
	fmt.Print(pageId)
	fileName := "files/" + pageId + ".html"
	http.ServeFile(w, r, fileName)
}
func pageHandlerWith404(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pageId := vars["id"]
	fileName := "Serving_and_Routing/files/" + pageId + ".html"
	//here we chech if this file exist or not
	if _, err := os.Stat(fileName); err != nil {
		if os.IsNotExist(err) {
			fileName = "Serving_and_Routing/files/404.html"
		}
	}
	http.ServeFile(w, r, fileName)
}
func main() {
	rtr := mux.NewRouter()
	//First example with custom 404 not found errors
	//rtr.HandleFunc("/pages/{id:[0-9]+}",pageHandler)

	//Hint
	//If you pay attention we have only four files inside of files folder if we go further out of this four index's files
	// to search we got 404 error not found file (404 pagenotfound exception) this is default not found exception
	// and we could create own 404 not found exception
	rtr.HandleFunc("/pages/{id:[0-9]+}", pageHandlerWith404)
	http.Handle("/", rtr)
	http.ListenAndServe(":8080", nil)

}
