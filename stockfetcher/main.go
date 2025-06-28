package main

import (
	"log"
	"net/http"
	"stockfetcher/client"
	"stockfetcher/db"
	"stockfetcher/handlers"
	"stockfetcher/middleware" 

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	// Cargar variables de entorno
	err := godotenv.Load()
	if err != nil {
		log.Fatal("❌ Error loading .env")
	}

	// Conectar a CockroachDB
	db.Connect()
	defer db.Close()

	// Obtener datos de la API
	stocks, err := client.FetchAllStocks()
	if err != nil {
		log.Fatalf("❌ Error fetching stocks: %v", err)
	}

	log.Printf("✅ Fetched %d stocks from API", len(stocks))

	// Insertarlos en la base de datos
	if err := db.InsertStocks(stocks); err != nil {
		log.Fatalf("❌ Error inserting stocks: %v", err)
	}

	log.Println("✅ Stocks inserted into database")

	r := mux.NewRouter()

	// Endpoints
	r.HandleFunc("/api/stocks", handlers.GetStocksHandler).Methods("GET")
	r.HandleFunc("/api/recommendations", handlers.GetRecommendationsHandler).Methods("GET")

	handler := middleware.EnableCORS(r)

	log.Println("🚀 Server running on http://localhost:8080")
	http.ListenAndServe(":8080", handler) 
}
