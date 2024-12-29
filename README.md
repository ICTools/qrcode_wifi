# Wi-Fi QR Code Generator in Go

This project is a Go-based utility to generate QR codes for Wi-Fi networks. Users can provide Wi-Fi details (SSID, password, authentication type, etc.) as command-line arguments, and the script generates a QR code image that can be scanned to connect to the network.

---

## Features
- Generates a QR code for Wi-Fi networks.
- Supports various authentication types (WPA, WEP, or none).
- Handles hidden and visible networks.
- Easy to use with command-line arguments.

---

## Prerequisites

- **Go**: Ensure Go is installed on your system. You can download it from [go.dev](https://go.dev/).
- **qrcode Library**: The script uses the `github.com/skip2/go-qrcode` library for generating QR codes.

---

## Installation

1. Clone the repository or save the script in a file (e.g., `main.go`).

2. Initialize a Go module (if not already done):
   ```bash
   go mod init wifi_qr_generator
   ```

3. Install the `qrcode` library:
   ```bash
   go get github.com/skip2/go-qrcode
   ```

---

## Usage

### Run the Script

Compile and run the script with the following command:

```bash
go run main.go --ssid="MyWiFi" --password="12345678" --auth="WPA" --hidden=false
```

### Command-Line Arguments

| Argument      | Description                                                   | Required | Default |
|---------------|---------------------------------------------------------------|----------|---------|
| `--ssid`      | The name of the Wi-Fi network.                                | Yes      | None    |
| `--password`  | The password for the Wi-Fi network.                           | No       | None    |
| `--auth`      | Authentication type: `WPA`, `WEP`, or empty for no encryption.| No       | `WPA`   |
| `--hidden`    | Whether the Wi-Fi network is hidden (`true` or `false`).      | No       | `false` |

### Output

The script generates a file named `wifi_qr_code.png` in the current directory. This file contains the QR code with the provided Wi-Fi details.

---

## Example

To generate a QR code for a Wi-Fi network:

- **SSID**: `HomeNetwork`
- **Password**: `mypassword`
- **Authentication**: `WPA`
- **Hidden**: `false`

Run:
```bash
go run main.go --ssid="HomeNetwork" --password="mypassword" --auth="WPA" --hidden=false
```

The output file `wifi_qr_code.png` will be created in the current directory.

---

## Troubleshooting

1. **Missing Module Error**: If you encounter an error like `go.mod file not found`, initialize the Go module using `go mod init wifi_qr_generator`.

2. **Library Not Found**: If the library is missing, ensure you install it using:
   ```bash
   go get github.com/skip2/go-qrcode
   ```

3. **Permission Issues**: Ensure you have write permissions for the directory where the script is executed.

---

## License
This project is licensed under the MIT License. Feel free to use and modify it as needed.
