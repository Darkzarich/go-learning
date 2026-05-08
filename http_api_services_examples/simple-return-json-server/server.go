package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func main() {
	http.HandleFunc("/products", productsHandler)
	
	log.Println("Starting HTTP server on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServeTLS: ", err)
	}
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	products := []Product{
		{ID: 1, Name: "Laptop", Price: 999.99},
		{ID: 2, Name: "Phone", Price: 699.99},
		{ID: 3, Name: "Tablet", Price: 399.99},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}