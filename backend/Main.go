package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Define backend servers with weights
	servers := map[string]int{
		"http://localhost:8081": 5,
		"http://localhost:8082": 3,
		"http://localhost:8083": 2,
	}

	// Initialize Load Balancer
	lb := NewLoadBalancer(servers)

	// API for forwarding requests
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		server := lb.NextServer()
		if server == nil {
			http.Error(w, "No available servers", http.StatusServiceUnavailable)
			return
		}
		fmt.Fprintf(w, "Forwarding request to: %s", server.URL)
	})

	// API for fetching server statuses (for frontend)
	http.HandleFunc("/servers", lb.GetServersStatus)

	fmt.Println("Load Balancer running on port 8080")
	http.ListenAndServe(":8080", nil)
}
