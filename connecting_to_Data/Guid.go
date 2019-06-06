package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var database1 *sql.DB

type Page1 struct {
	Title   string
	Content string
	Date    string
}

const (
	name     = "root"
	password = ""
	host     = "localhost:3306"
	ba       = "db"
)

func handler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid := vars["guid"]
	fmt.Println(guid)
	t := Page1{}
	err := database1.QueryRow("select page_t,content,date from pages where page_t = ?", guid).Scan(&t.Title, &t.Content, &t.Date)
	if err != nil {
		log.Println("error in getting")
	}
	html := `<html><head><title>` + t.Title + `</title></head><body><h1>` + t.Title + `</h1><div>` + t.Content + `</div></body></html>`
	fmt.Fprintln(w, html)

}
func main() {
	//The main reason why we use guid it shows string instead of ids in url
	//dbm := fmt.Sprintf("%s@tcp(%s)/%s",name,host,ba)
	// other way of connection
	//d := fmt.Sprintf("%s:%s@tcp(%s)/%s", dbuser,dbpassword,dbhost,dbname)
	dbconn, err := sql.Open("mysql", "root:@/db")
	if err != nil {
		log.Println("error in connection")
		log.Println(err.Error)
	}
	database1 = dbconn
	routes := mux.NewRouter()
	routes.HandleFunc("/pages/{guid:[0-9a-zA\\-]+}", handler)
	http.Handle("/", routes)
	http.ListenAndServe(":8080", nil)

}
