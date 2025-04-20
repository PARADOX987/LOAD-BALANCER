package main

import (
	"encoding/json"
	"net/http"
	"sync"
	"time"
)

// Server struct represents a backend server
type Server struct {
	URL            string `json:"url"`
	Active         bool   `json:"active"`
	Weight         int
	CurrentWeight  int
}

// LoadBalancer struct manages backend servers
type LoadBalancer struct {
	Servers []*Server
	Mutex   sync.Mutex
}

// Create a new Load Balancer with weighted servers
func NewLoadBalancer(servers map[string]int) *LoadBalancer {
	lb := &LoadBalancer{}
	for url, weight := range servers {
		lb.Servers = append(lb.Servers, &Server{URL: url, Active: true, Weight: weight, CurrentWeight: weight})
	}
	go lb.healthCheck() // Start health checks
	return lb
}

// Selects the next available server using Weighted Round Robin
func (lb *LoadBalancer) NextServer() *Server {
	lb.Mutex.Lock()
	defer lb.Mutex.Unlock()

	var best *Server
	for _, server := range lb.Servers {
		if server.Active {
			server.CurrentWeight += server.Weight
			if best == nil || server.CurrentWeight > best.CurrentWeight {
				best = server
			}
		}
	}

	if best != nil {
		best.CurrentWeight -= sumWeights(lb.Servers)
	}

	return best
}

// Sum of all active servers' weights
func sumWeights(servers []*Server) int {
	sum := 0
	for _, server := range servers {
		sum += server.Weight
	}
	return sum
}

// Periodically checks server health
func (lb *LoadBalancer) healthCheck() {
	for {
		for _, server := range lb.Servers {
			resp, err := http.Get(server.URL)
			if err != nil || resp.StatusCode != http.StatusOK {
				server.Active = false
			} else {
				server.Active = true
			}
		}
		time.Sleep(5 * time.Second)
	}
}

// API to return server statuses for the frontend
func (lb *LoadBalancer) GetServersStatus(w http.ResponseWriter, r *http.Request) {
	lb.Mutex.Lock()
	defer lb.Mutex.Unlock()

	status := []Server{}
	for _, server := range lb.Servers {
		status = append(status, Server{
			URL:    server.URL,
			Active: server.Active,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(status)
}
