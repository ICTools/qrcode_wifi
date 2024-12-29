package main

import (
	"flag"
	"fmt"
	"log"

	qrcode "github.com/skip2/go-qrcode"
)

func main() {
	ssid := flag.String("ssid", "", "Wi-Fi network name (SSID)")
	password := flag.String("password", "", "Wi-Fi network password")
	authType := flag.String("auth", "WPA", "Authentication type (WPA, WEP, or leave empty for no encryption)")
	hidden := flag.Bool("hidden", false, "Whether the network is hidden (true or false)")

	flag.Parse()

	if *ssid == "" {
		log.Fatalf("The --ssid parameter is required.")
	}

	hiddenStr := "false"
	if *hidden {
		hiddenStr = "true"
	}
	wifiData := fmt.Sprintf("WIFI:T:%s;S:%s;P:%s;H:%s;;", *authType, *ssid, *password, hiddenStr)

	err := qrcode.WriteFile(wifiData, qrcode.Medium, 256, "wifi_qr_code.png")
	if err != nil {
		log.Fatalf("Error generating QR code: %v", err)
	}

	fmt.Println("QR code successfully generated: wifi_qr_code.png")
}
