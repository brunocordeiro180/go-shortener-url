package main

import (
	"github.com/brunocordeiro180/go-url-shortener/db"
	"github.com/brunocordeiro180/go-url-shortener/repository"
	_ "github.com/mattn/go-sqlite3"
	"html/template"
	"log"
	"net/http"
	"sync"
)

type PageData struct {
	ShortenedURL string
	Host         string
}

var mutex = sync.RWMutex{}

func main() {

	myDb, err := db.Connect()
	defer myDb.Close()

	if err != nil {
		panic(err)
	}

	err = db.MigrateTables(myDb)
	if err != nil {
		panic(err)
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			inputURL := r.FormValue("url")
			shortURL, err := repository.GetOrCreateUrl(myDb, inputURL)
			if err != nil {
				http.Error(w, "Failed to shorten URL", http.StatusInternalServerError)
				return
			}
			myHost := "http://localhost:8080"
			data := PageData{ShortenedURL: shortURL, Host: myHost}
			if err := tmpl.Execute(w, data); err != nil {
				log.Printf("Error executing template: %v", err)
			}
		} else {
			key := r.URL.Path[1:] // Remove leading "/"
			if key == "" {
				if err := tmpl.Execute(w, nil); err != nil {
					log.Printf("Error executing template: %v", err)
				}
			} else {
				mutex.RLock()
				originalURL, exists := repository.GetUrlFromKey(myDb, key)
				mutex.RUnlock()

				if !exists {
					http.NotFound(w, r)
					return
				}

				// Redirect to the original URL
				http.Redirect(w, r, originalURL, http.StatusFound)
			}
		}
	})

	// Start the server
	log.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
