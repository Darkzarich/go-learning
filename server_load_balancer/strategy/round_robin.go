package strategy

import (
	"errors"
	"sync"

	"server_load_balancer/pool"
	"server_load_balancer/types"
)

type RoundRobinBalancer struct {
	pool *pool.ServerPool
	mu   sync.Mutex // We need this for locking index update
	idx  int
}

func NewRoundRobinBalancer(pool *pool.ServerPool) *RoundRobinBalancer {
	return &RoundRobinBalancer{
		pool: pool,
		idx:  -1, // No server has been selected yet
	}
}

func (rrb *RoundRobinBalancer) GetNextServer() (*types.Server, error) {
	servers := rrb.pool.GetAllServers()

	if len(servers) == 0 {
		return nil, errors.New("No servers available")
	}

	// We need to lock here to avoid race conditions,
	// when two requests change the index at the same time
	rrb.mu.Lock()
	defer rrb.mu.Unlock()

	rrb.idx++

	if rrb.idx >= len(servers) {
		rrb.idx = 0
	}

	selectedServer := servers[rrb.idx]

	return selectedServer, nil
}
