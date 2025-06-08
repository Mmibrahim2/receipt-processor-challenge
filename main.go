package main

import (
	"log"
	"net/http"
)

func main() {
	// Register endpoint handlers
	http.HandleFunc("/receipts/process", handleProcessReceipt)
	http.HandleFunc("/receipts/", handleGetPoints) // Will match /receipts/{id}/points

	log.Println("Receipt Processor API is running at http://localhost:8080")
	log.Println("Endpoints:")
	log.Println("  POST /receipts/process")
	log.Println("  GET  /receipts/{id}/points")

	// Start HTTP server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
