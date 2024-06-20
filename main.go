package main

import (
	"log"
	"net/http"
	"os"

	"go-file-explorer/handlers"
)

func main() {
	filesPath := os.Getenv("FILES_PATH")
	if filesPath == "" {
		log.Fatal("FILES_PATH environment variable is not set")
	}

	http.HandleFunc("/", handlers.FileExplorerHandler(filesPath))
	http.HandleFunc("/files/", handlers.FileHandler(filesPath))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Println("Server started at :8282")
	log.Fatal(http.ListenAndServe(":8282", nil))
}
