package main

import (
	"fmt"
	"net/http"
)

func main() {
	server := http.NewServeMux()
	server.Handle("GET /checks", http.HandlerFunc(PlayerServer))
	server.Handle("POST /checks", http.HandlerFunc(PlayerServer))
	// server.HandleFunc("GET /checks", p.PlayerServer)
	http.ListenAndServe(":8080", server)

	// handler := http.HandlerFunc(PlayerServer)
	// http.Handle("/checks", handler)
	// if err := http.ListenAndServe(":8080", handler); err != nil {
	// 	log.Fatalf("Failed to start server: %v", err)
	// }

	// http.HandleFunc("/check", PlayerServer)
	// if err := http.ListenAndServe(":8080", nil); err != nil {
	// 	log.Fatalf("Failed to start server: %v", err)
	// }
}

func PlayerServer(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	// Handle the request here
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Player Server is running"))
}
