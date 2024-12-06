package bruteforce

import (
	"bufio"
	"fmt"
	"github.com/0xQRx/subbrute/internal/dns"
	"os"
	"sync"
)

type queueItem struct {
	domain string
	depth  int
}

var (
	queue   = make(chan queueItem, 1000)
	wg      sync.WaitGroup
	sem     chan struct{}
	queueWg sync.WaitGroup // Tracks how many domains are left to process
)

func Run(t int, domain, wordlist, ns string, depth int) {
	sem = make(chan struct{}, t)

	// Start workers
	wg.Add(t)
	for i := 0; i < t; i++ {
		go worker(ns, wordlist)
	}

	// Add initial domain to the queue
	queueWg.Add(1)
	queue <- queueItem{domain: domain, depth: depth}

	// Close the queue when all tasks are done
	go func() {
		queueWg.Wait()
		close(queue)
	}()

	// Wait for all workers to finish
	wg.Wait()
}

func processWordlist(baseDomain, wordlist string) []string {
	file, err := os.Open(wordlist)
	if err != nil {
		fmt.Printf("Error opening wordlist: %v\n", err)
		return nil
	}
	defer file.Close()

	var subdomains []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		subdomains = append(subdomains, fmt.Sprintf("%s.%s", scanner.Text(), baseDomain))
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading wordlist: %v\n", err)
	}
	return subdomains
}

func worker(ns, wordlist string) {
	defer wg.Done()

	for item := range queue {
		sem <- struct{}{}
		//fmt.Printf("SEARCHING: %s\n", item.domain)
		// Check if the domain resolves
		if dns.Resolve(item.domain, ns) {
			fmt.Printf("%s\n", item.domain)
		}

		// Regardless of resolution, if depth > 0, generate and enqueue next-level subdomains
		if item.depth > 0 {
			subdomains := processWordlist(item.domain, wordlist)
			for _, s := range subdomains {
				queueWg.Add(1)
				queue <- queueItem{domain: s, depth: item.depth - 1}
			}
		}

		// Mark this domain as processed
		queueWg.Done()

		<-sem
	}
}
