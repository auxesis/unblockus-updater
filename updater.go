package main

import (
	"log"
	"net/http"
	"strings"
	"os"
	"fmt"
)

func GetCredentials() (string, string) {
	if len(os.Args) != 3 {
		fmt.Println("Usage: unblockus-updater <username> <password>")
		os.Exit(1)
	}

	username := os.Args[1]
	password := os.Args[2]

	return username, password
}

func GetUrl(username string, password string) string {
	//base := "http://localhost:9292/login"
	base := "https://api.unblock-us.com/login"
	parts := []string{base, "?", username, ":", password}
	url := strings.Join(parts, "")

	return url
}

func main() {
	username, password := GetCredentials()
	url := GetUrl(username, password)

	log.Println(url)

	_, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
}
