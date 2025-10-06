package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", HelloHandler)
	fmt.Printf("Serveur running at port 8090")
	http.ListenAndServe(":8090", nil)
}
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello word  from Dockers")
}
