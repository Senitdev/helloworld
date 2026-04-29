package main

import (
	"fmt"
	"helloword/database"
	"helloword/setup"
	"net/http"
	"os"
)

func main() {
	database.Connect()
	r := setup.SetupRoutes(database.DB)
	/*
		http.HandleFunc("/user", handler)
		fmt.Printf("Serveur running at port 8090")
		http.ListenAndServe(":8090", nil)
	*/
	port := os.Getenv("API_PORT")
	if port == "" {
		port = "8090"
	}
	r.Run(":" + port)
	fmt.Printf("Server Running at port %s", port)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"ok":true}`))
}
