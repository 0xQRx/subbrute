package dns

import (
	"context"
	"fmt"
	"net"
	"strings"
	"sync"
	"time"
)

var (
	nsList  []string
	nsIndex int
	mutex   sync.Mutex
)

// Resolve checks if a subdomain resolves within a controlled timeout.
func Resolve(subdomain, ns string) bool {
	resolver := roundRobinResolver(ns)

	// Create a context with timeout for the lookup
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := resolver.LookupHost(ctx, subdomain)
	return err == nil
}

// roundRobinResolver provides a round-robin DNS resolver with a timeout and custom DNS servers.
func roundRobinResolver(ns string) *net.Resolver {
	if ns == "" {
		return net.DefaultResolver
	}

	mutex.Lock()
	if len(nsList) == 0 {
		nsList = strings.Split(ns, ",")
	}
	dnsServer := nsList[nsIndex%len(nsList)]
	nsIndex++
	mutex.Unlock()

	// Using a shorter dial timeout for efficiency
	dialer := &net.Dialer{Timeout: 3 * time.Second}

	return &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			return dialer.DialContext(ctx, "udp", fmt.Sprintf("%s:53", dnsServer))
		},
	}
}
