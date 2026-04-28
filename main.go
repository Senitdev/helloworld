package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/user", handler)
	fmt.Printf("Serveur running at port 8090")
	http.ListenAndServe(":8090", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"ok":true}`))
}
