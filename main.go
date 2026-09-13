package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"time"
)

type Response struct {
	Days int    `json:"days"`
	From string `json:"from"`
	To   string `json:"to"`
	Err  string `json:"error,omitempty"`
}

func NewMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/days", handleDays)
	return mux
}

func handleDays(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		_ = json.NewEncoder(w).Encode(Response{Err: "method not allowed"})
		return
	}

	dateParam := r.URL.Query().Get("date")
	var from time.Time

	if dateParam == "" {
		from = time.Now()
	} else {
		parsed, err := time.Parse("2006-01-02", dateParam)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(w).Encode(Response{Err: "invalid date format, expected YYYY-MM-DD"})
			return
		}
		from = parsed
	}

	days, err := DaysUntilNewYear(from)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(Response{Err: err.Error()})
		return
	}

	resp := Response{
		Days: days,
		From: from.Format("2006-01-02"),
		To:   time.Date(from.Year()+1, time.January, 1, 0, 0, 0, 0, from.Location()).Format("2006-01-02"),
	}

	_ = json.NewEncoder(w).Encode(resp)
}

func main() {
	addr := flag.String("addr", ":8080", "listen address")
	flag.Parse()

	log.Printf("starting server on %s", *addr)
	if err := http.ListenAndServe(*addr, NewMux()); err != nil {
		log.Fatal(err)
	}
}
