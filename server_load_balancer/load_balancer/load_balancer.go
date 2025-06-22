package load_balancer

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"server_load_balancer/types"
)

type LoadBalancerStrategy interface {
	GetNextServer() (*types.Server, error)
}

type LoadBalancer struct {
	strategy LoadBalancerStrategy
}

func NewLoadBalancer(strategy LoadBalancerStrategy) *LoadBalancer {
	return &LoadBalancer{
		strategy: strategy,
	}
}

func (lb *LoadBalancer) BalanceRequest(w http.ResponseWriter, r *http.Request) {
	server, err := lb.strategy.GetNextServer()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	targetURL, err := url.Parse(server.URL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	targetPath := strings.TrimSuffix(targetURL.String(), "/") + r.URL.Path
	if r.URL.RawQuery != "" {
		targetPath += "?" + r.URL.RawQuery // Copy query from original request
	}

	req, err := http.NewRequest(r.Method, targetPath, r.Body) // Copy method and headers from original request
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Copy headers from original request
	for key, values := range r.Header {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}

	req.Header.Set("X-Forwarded-For", r.RemoteAddr)

	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	res, err := client.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	byteRes, err := io.ReadAll(res.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer res.Body.Close()

	for key, values := range res.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}

	w.WriteHeader(res.StatusCode)

	fmt.Fprintf(w, string(byteRes))
}
