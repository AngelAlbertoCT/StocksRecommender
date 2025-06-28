package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"stockfetcher/db"
)

func GetRecommendationsHandler(w http.ResponseWriter, r *http.Request) {
	profile := r.URL.Query().Get("profile")
	if profile == "" {
		http.Error(w, "Missing profile parameter", http.StatusBadRequest)
		return
	}

	page := 1
	pageStr := r.URL.Query().Get("page")
	if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
		page = p
	}

	const limit = 10

	stocks, err := db.GetRecommendedStocks(profile, page, limit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stocks)
}
