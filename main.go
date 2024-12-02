package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var db *sql.DB

func main() {
	var err error
	db, err = ConnectDB()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	defer db.Close()

	r := mux.NewRouter()
	r.HandleFunc("/current-time", CurrentTimeHandler).Methods("GET")

	log.Println("Starting server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func CurrentTimeHandler(w http.ResponseWriter, r *http.Request) {
	location, err := time.LoadLocation("America/Toronto")
	if err != nil {
		http.Error(w, "Failed to load timezone", http.StatusInternalServerError)
		return
	}

	currentTime := time.Now().In(location)

	// Log to database
	_, err = db.Exec("INSERT INTO time_log (timestamp) VALUES (?)", currentTime)
	if err != nil {
		http.Error(w, "Failed to log time to database", http.StatusInternalServerError)
		return
	}

	response := map[string]string{
		"current_time": currentTime.Format("2006-01-02 15:04:05"),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
