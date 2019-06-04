package Serving_and_Routing

import (
	"net/http"
)

func main() {
	http.ListenAndServe(":8080", http.FileServer(http.Dir("/var/file")))
}
