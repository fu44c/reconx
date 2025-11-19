````markdown
# ReconX
![ReconX Logo](./assets/logo.png)

**ReconX** is an advanced **Reconnaissance & Bug Bounty** tool designed to automate domain collection, vulnerability scanning, and security assessments for bug bounty hunters and penetration testers.

The tool integrates various techniques and third-party APIs to efficiently perform reconnaissance on websites and services, identifying subdomains, scanning for live hosts, and checking for common security vulnerabilities like SQLi, XSS, and LFI.

---

## Features

- **Domain Collection**: Gather subdomains from multiple sources.
- **HTTP/HTTPS Scanning**: Identify live hosts and check protocol responses.
- **Vulnerability Scanning**: Detect common vulnerabilities like SQLi, XSS, LFI, and more.
- **Automation**: Run the entire workflow (collection, scanning, vulnerability check) automatically.
- **Customizable Reports**: Output results in **JSON** or **HTML** format for easy analysis.

---

## Requirements

- **Go** >= 1.20 (for the Go version)
- **Python** 3.10+ (for the Bash version)
- **Dependencies**:
  - `httpx`, `subfinder`, `nmap`, `ghauri`
  - Install these dependencies using:
    ```bash
    go get github.com/projectdiscovery/httpx/cmd/httpx@latest
    go get github.com/projectdiscovery/subfinder/v2/cmd/subfinder@latest
    pip3 install ghauri
    ```

---

## Installation

Clone the repository:
```bash
git clone https://github.com/fu44c/reconx.git
cd reconx
````

### Install Dependencies

```bash
# Update system packages
sudo apt update && sudo apt upgrade -y

# Install Go and Python
sudo apt install go python3 python3-pip -y

# Install dependencies
go mod tidy
go get github.com/spf13/cobra
```

---

## Usage

### Commands:

1. **Collect Subdomains**:
   Collect subdomains for the target domain:

   ```bash
   go run main.go collect -d example.com
   ```

2. **Scan HTTP/HTTPS**:
   Scan for live HTTP/HTTPS endpoints:

   ```bash
   go run main.go scan -d example.com
   ```

3. **Check Vulnerabilities**:
   Scan for vulnerabilities (e.g., SQLi, XSS, LFI):

   ```bash
   go run main.go vuln -d example.com
   ```

4. **Full Automation**:
   Automatically run the entire workflow (collect, scan, vuln):

   ```bash
   go run main.go --auto -d example.com --output report.json
   ```

---

## Example Output

Once the scan is completed, you can expect the following output:

1. **Subdomain Collection**:

   ```bash
   [*] Collecting subdomains for example.com...
   [+] Found: api.example.com
   [+] Found: dev.example.com
   ```

2. **HTTP/HTTPS Scan**:

   ```bash
   [*] Scanning HTTP/HTTPS for example.com...
   [LIVE] https://api.example.com
   [DEAD] http://dev.example.com
   ```

3. **Vulnerability Scan**:

   ```bash
   [*] Checking vulnerabilities for example.com...
   [+] No SQLi detected
   [+] Possible XSS endpoint at /search
   ```

---

## Reports

You can output the results of your scans in **JSON** or **HTML** format using the `--output` flag.

Example:

```bash
go run main.go --auto -d example.com --output report.json
```

---

## Project Structure

```
ReconX/
├── cmd/                     # Commands (collect, scan, vuln)
│   ├── root.go              # Main command
│   ├── collect.go           # Subdomain collection command
│   ├── scan.go              # HTTP/HTTPS scanning command
│   └── vuln.go              # Vulnerability scanning command
├── internal/
│   ├── collector/           # Subdomain collection logic
│   ├── scanner/             # HTTP/HTTPS scanning logic
│   ├── vuln/                # Vulnerability scanning logic
│   └── report/              # Report generation logic (JSON/HTML)
├── assets/
│   └── logo.png             # Logo for the tool
├── go.mod                   # Go module file
├── go.sum                   # Go sum file for dependencies
├── README.md                # Project documentation
└── LICENSE                  # MIT License
```

---

## Contributing

We welcome contributions to improve the project. To contribute:

1. **Fork the repository.**
2. **Create a new branch**:

   ```bash
   git checkout -b feature-name
   ```
3. **Commit your changes**:

   ```bash
   git commit -m "Add new feature"
   ```
4. **Push your branch**:

   ```bash
   git push origin feature-name
   ```
5. **Create a pull request** on GitHub.

---

## License

This project is licensed under the **MIT License**.
You are free to use, modify, and distribute it for educational or Bug Bounty purposes — with proper credit to the developer.

---

## Support

For questions, suggestions, or feature requests, contact **Ahmed Ibrahim** via [GitHub](https://github.com/fu44c).

```
