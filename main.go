package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", HelloHandler)
	http.HandleFunc("/test", TestWorkFlow)
	http.HandleFunc("/add", AddRoute)
	fmt.Printf("Serveur running at port 8090")
	http.ListenAndServe(":8090", nil)
}
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello word  from Dockers")
}
func TestWorkFlow(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "je test le workflow")
}
func AddRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Encore une route a tester")
}
