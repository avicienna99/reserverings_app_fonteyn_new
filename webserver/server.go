package server

import (
	"html/template"
	"log"
	"net/http"

	"github.com/avicienna99/reserverings_app_fonteyn_new/db"
)

type PageData struct {
	Houses []db.House
}

// handler serves the main HTML page with a list of houses
func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request received: %s %s", r.Method, r.URL.Path)

	// Fetch houses from the database
	houses, err := db.GetHouses()
	if err != nil {
		http.Error(w, "Error fetching house data", http.StatusInternalServerError)
		return
	}

	// Parse the HTML template
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}

	// Render the template with house data
	data := PageData{Houses: houses}
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}

// Start initializes and starts the HTTP server
func Start() {
	http.HandleFunc("/", handler)

	log.Println("Starting server on http://localhost:8080")
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
