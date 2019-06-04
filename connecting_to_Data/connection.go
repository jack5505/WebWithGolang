package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const (
	DBHost  = "127.0.0.1"
	DBPort  = ":3306"
	DBUser  = "root"
	DBPass  = ""
	DBDbase = "db"
	PORT    = ":8080"
)

var database *sql.DB

type Page struct {
	Title   string
	Content string
	Date    string
}

func ServePage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pageId := vars["id"]
	thisPage := Page{}
	fmt.Println(pageId)
	err := database.QueryRow("select page_t,content,date from pages where id = ?", pageId).Scan(&thisPage.Title,
		&thisPage.Content, &thisPage.Date)
	if err != nil {
		log.Println("Couldn't get page" + pageId)
		log.Println(err.Error())
	}
	html := `<html><head><title>` + thisPage.Title + `</title></head><body><h1>` + thisPage.Title + `</h1><div>` + thisPage.Content + `</div></body></html>`
	fmt.Fprintln(w, html)
}
func main() {
	//before to run this code make sure you have page table and db database
	dbConn := fmt.Sprintf("%s:%s@/%s", DBUser, DBPass, DBDbase)
	fmt.Println(dbConn)
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		log.Println("Couldn't connect to" + DBDbase)
		log.Println(err.Error)
	}
	database = db

	routes := mux.NewRouter()
	routes.HandleFunc("/page/{id:[0-9]+}", ServePage)
	http.Handle("/", routes)
	http.ListenAndServe(PORT, nil)

}
