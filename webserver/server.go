package server

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	"github.com/avicienna99/reserverings_app_fonteyn_new/db"
)

type PageData struct {
	Houses []db.House
}

type Reservation struct {
	HouseID   int    `json:"house_id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	houses, err := db.GetHouses()
	if err != nil {
		http.Error(w, "Error fetching house data", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}

	data := PageData{Houses: houses}
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}

func reserveHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var res Reservation
	err := json.NewDecoder(r.Body).Decode(&res)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = db.AddReservation(db.Reservation{
		HouseID:   res.HouseID,
		Name:      res.Name,
		Email:     res.Email,
		StartDate: res.StartDate,
		EndDate:   res.EndDate,
	})
	if err != nil {
		http.Error(w, "Failed to create reservation", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func Start() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/reserve", reserveHandler)

	log.Println("Starting server on http://localhost:8080")
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
