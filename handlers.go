package main

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

// handleProcessReceipt processes a new receipt, calculates points, stores them, and returns a UUID.
func handleProcessReceipt(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var receipt Receipt
	if err := json.NewDecoder(r.Body).Decode(&receipt); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	id := uuid.NewString()
	points := CalculatePoints(receipt)
	savePoints(id, points)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"id": id})
}

// handleGetPoints retrieves points associated with a receipt ID passed in the path: /receipts/{id}/points
func handleGetPoints(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Expected path: /receipts/{id}/points
	pathParts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	if len(pathParts) != 3 || pathParts[0] != "receipts" || pathParts[2] != "points" {
		http.Error(w, "Bad Request: URL should be /receipts/{id}/points", http.StatusBadRequest)
		return
	}

	id := pathParts[1]
	points, found := getPoints(id)
	if !found {
		http.Error(w, "Receipt ID not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int{"points": points})
}
