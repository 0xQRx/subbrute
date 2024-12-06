
# Subdomain Bruteforce Tool

## Overview
This tool is a highly efficient and recursive subdomain brute-forcing program written in Go. It supports multiple levels of recursion to discover nested subdomains, customizable concurrency, and the use of multiple DNS servers for resolution.

### Features:
- Recursive subdomain brute-forcing with configurable depth.
- Supports multiple DNS servers with round-robin selection.
- Timeout-controlled DNS resolution for efficiency.
- Handles large wordlists and high concurrency.
- Ensures program termination after processing all subdomains.

---

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/0xQRx/subbrute.git
   cd subbrute
   ```

2. Install the tool:
   ```bash
   go install ./cmd/subbrute
   ```

3. Ensure the binary is available in your `$PATH` (typically `~/go/bin`).

---

## Usage

### Basic Command
```bash
subbrute -d <domain> -w <wordlist> [-t <goroutines>] [-ns <dns servers>] [--depth <depth>]
```

### Parameters
- `-d`: **Required**. The domain to start brute-forcing (e.g., `example.com`).
- `-w`: **Required**. Path to a wordlist file (e.g., `dns-names.txt`).
- `-t`: Number of concurrent goroutines (default: `10`).
- `-ns`: Comma-separated DNS servers to use (e.g., `8.8.8.8,1.1.1.1`). Defaults to the system DNS if not specified.
- `--depth`: Depth of recursive brute-forcing. Determines how many levels of subdomains will be generated and checked (default: `1`).

---

## Example

### Single-Level Brute Forcing
Discover first-level subdomains for `example.com`:
```bash
subbrute -d example.com -w dns-names.txt -t 20 -ns 8.8.8.8
```

### Multi-Level Recursive Brute Forcing
Discover subdomains up to 3 levels deep:
```bash
subbrute -d example.com -w dns-names.txt -t 20 -ns 8.8.8.8,1.1.1.1 --depth 3
```

### Output Example
For a domain `example.com` with a depth of 2 and a wordlist containing 13 entries, the output might look like:
```plaintext
Found: www.example.com
Found: mail.example.com
Found: api.example.com
Found: admin.www.example.com
Found: staging.mail.example.com
Found: v2.api.example.com
...
```

---

## How It Works
1. **Initial Domain:** The tool starts by resolving the given domain.
2. **Wordlist Application:** For each domain, it appends entries from the wordlist to generate subdomains.
3. **Recursive Resolution:** At each level, it resolves subdomains and generates further subdomains for unresolved ones until the specified depth is reached.
4. **DNS Resolution:** Queries are sent to DNS servers with a timeout of 3 seconds to ensure efficiency.
5. **Concurrency Management:** Uses a goroutine pool (`-t` parameter) to limit simultaneous DNS queries.
6. **Program Termination:** Ensures all subdomains are processed and resolves gracefully.

---

## Performance Details
### Depth and Wordlist
- The total number of subdomains checked is exponential based on the depth and wordlist size.
- Example:
  - **Depth:** 3
  - **Wordlist Size:** 13
  - **Total Subdomains Checked:** 2380

### Efficiency Enhancements
- DNS queries are managed with timeouts to avoid hanging.
- Round-robin DNS server selection distributes load evenly.
- Channels and wait groups ensure controlled concurrency and program termination.

---

## Requirements
- **Go**: Version 1.19+
- A valid wordlist file for subdomain names.

---

## Troubleshooting

### Program Hangs
- Ensure your DNS servers are responsive and not rate-limited.
- Reduce the depth or concurrency (`-t`) if processing too many subdomains overwhelms the network or system.

### Infinite Subdomain Generation
- Ensure your wordlist does not contain entries that create infinite loops (e.g., redundant entries that result in cyclic subdomains).

---

## License
This tool is released under the MIT License. Feel free to use, modify, and distribute it.
