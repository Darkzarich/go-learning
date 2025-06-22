package types

import "time"

type Server struct {
	ID        string
	Name      string
	URL       string
	IsHealthy bool
	LastCheck time.Time
}
