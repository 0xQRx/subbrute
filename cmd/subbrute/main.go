package main

import (
	"flag"
	"fmt"
	"github.com/0xQRx/subbrute/internal/bruteforce"
	"os"
)

func main() {
	// Command-line arguments
	var t int
	var domain, wordlist, ns string
	var depth int

	flag.IntVar(&t, "t", 10, "Number of goroutines")
	flag.StringVar(&domain, "d", "", "Domain to start brute-forcing (e.g., google.com)")
	flag.StringVar(&wordlist, "w", "", "Wordlist file (e.g., dns-names.txt)")
	flag.StringVar(&ns, "ns", "", "Comma-separated DNS servers (e.g., 8.8.8.8,1.1.1.1)")
	flag.IntVar(&depth, "depth", 1, "Depth of recursive brute-forcing")

	flag.Parse()

	// Validate input
	if domain == "" || wordlist == "" {
		fmt.Println("Usage: subbrute -d <domain> -w <wordlist> [-t <goroutines>] [-ns <dns servers>] [--depth <depth>]")
		os.Exit(1)
	}

	// Start the brute-force process
	bruteforce.Run(t, domain, wordlist, ns, depth)
}
