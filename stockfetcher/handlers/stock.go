package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"stockfetcher/db"
)

func GetStocksHandler(w http.ResponseWriter, r *http.Request) {
	ticker := r.URL.Query().Get("ticker")
	company := r.URL.Query().Get("company")
	sort := r.URL.Query().Get("sort") 

	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")
	monthStr := r.URL.Query().Get("month")

	// Valores por defecto
	page := 1
	limit := 10
	month := 0 // 0 = no filtrar por mes

	if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
		page = p
	}
	if l, err := strconv.Atoi(limitStr); err == nil && l > 0 {
		limit = l
	}
	if m, err := strconv.Atoi(monthStr); err == nil && m >= 1 && m <= 12 {
		month = m
	}

	stocks, err := db.GetFilteredStocks(ticker, company, sort, page, limit, month)
	if err != nil {
		http.Error(w, "Error retrieving stocks", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stocks)
}
