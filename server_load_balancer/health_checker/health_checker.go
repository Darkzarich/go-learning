package health_checker

import (
	"fmt"
	"net/http"
	"server_load_balancer/pool"
	"server_load_balancer/types"
	"time"

	"github.com/fatih/color"
)

type HealthChecker struct {
	interval time.Duration
	pool     *pool.ServerPool
}

func NewHealthChecker(interval time.Duration, pool *pool.ServerPool) *HealthChecker {
	return &HealthChecker{
		interval: interval,
		pool:     pool,
	}
}

func (hc *HealthChecker) CheckHealthPeriodically() {
	servers := hc.pool.GetAllServers()

	fmt.Printf("Checking health of %d servers\n", len(servers))

	for {
		color.Yellow("\nChecking health...\n")

		for _, server := range servers {
			isHealthy := hc.CheckHealth(server)

			hc.UpdateServerStatus(server, isHealthy)
		}

		for _, server := range servers {
			fmt.Printf("Server %s status: ", server.ID)

			if !server.IsHealthy {
				color.Red("NOT")
			} else {
				color.Green("OK")
			}
		}

		time.Sleep(hc.interval)
	}
}

func (hc *HealthChecker) CheckHealth(server *types.Server) bool {
	resp, err := http.Get(server.URL)

	if err != nil {
		return false
	}

	defer resp.Body.Close()

	return resp.StatusCode == http.StatusOK
}

func (hc *HealthChecker) UpdateServerStatus(server *types.Server, isHealthy bool) {
	server.IsHealthy = isHealthy
}
