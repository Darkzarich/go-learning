package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"server_load_balancer/config"
	"server_load_balancer/health_checker"
	"server_load_balancer/load_balancer"
	"server_load_balancer/pool"
	"server_load_balancer/strategy"
)

func main() {
	servers := config.ReadConfig().Servers

	sp := pool.NewServerPool()

	for _, server := range servers {
		sp.AddServer(server)
	}

	rb := strategy.NewRoundRobinBalancer(sp)

	lb := load_balancer.NewLoadBalancer(rb)

	healthChecker := health_checker.NewHealthChecker(10*time.Second, sp)

	go healthChecker.CheckHealthPeriodically()

	distributeRequests(3000, lb)
}

func distributeRequests(port int, lb *load_balancer.LoadBalancer) {
	mux := http.NewServeMux()

	mux.HandleFunc("/", lb.BalanceRequest)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}

	log.Printf("Starting load balancer on port %d", port)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
