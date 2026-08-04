package sentinel

import (
	"sync"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type RootLB struct {
	hostToBalancer map[string]middleware.ProxyBalancer
	mutex          sync.RWMutex
}

func NewRootLB() *RootLB {
	return &RootLB{
		hostToBalancer: make(map[string]middleware.ProxyBalancer),
	}
}

func (h *RootLB) AddLB(lb middleware.ProxyBalancer, allowedHosts []string) {
	h.mutex.Lock()
	defer h.mutex.Unlock()

	for _, host := range allowedHosts {
		h.hostToBalancer[host] = lb
	}
}

// no-op to satisfy lb interface
func (h *RootLB) AddTarget(target *middleware.ProxyTarget) bool {
	return true
}

// no-op to satisfy lb interface
func (h *RootLB) RemoveTarget(target string) bool {
	return true
}

func (h *RootLB) Next(c echo.Context) *middleware.ProxyTarget {
	h.mutex.RLock()
	defer h.mutex.RUnlock()

	host := c.Request().Host

	// O(1) direct map lookup
	balancer, exists := h.hostToBalancer[host]
	if !exists {
		return nil // Triggers Echo's 502 Bad Gateway
	}

	return balancer.Next(c)
}
