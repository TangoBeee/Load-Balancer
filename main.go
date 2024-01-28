package main

import (
	"loadbalancer/loadbalancer"
	"loadbalancer/servers"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	// Run the server in a goroutine
	go func() {
		defer wg.Done()
		servers.RunServers(5)
	}()

	// Run the load balancer in a goroutine
	go func() {
		defer wg.Done()
		loadbalancer.MakeLoadBalancer(5)
	}()

	// Wait for both goroutines to finish
	wg.Wait()
}
