package main

import (
	"errors"
	"fmt"
	vercheck "github.com/mcuadros/go-version"
	"log"
	"net/http"
	"path/filepath"
)

var (
	version = "DEBUG"
	tagCheckUrl = "https://github.com/BGrewell/go-update-test/releases/latest"
	updateUrlStub = "https://github.com/BGrewell/go-update-test/releases/download/%VER%/go-update-test_%VER%_Linux_x86_64.tar.gz"
)

func latestTaggedVersion() string {
	client := new(http.Client)
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return errors.New("Redirect")
	}
	req, err := http.NewRequest("GET", tagCheckUrl, nil)
	if err != nil {
		log.Fatalf("failed to get latest version. failed to build request: %v", err)
	}
	response, err := client.Do(req)
	if err != nil && response.StatusCode != http.StatusFound {
		log.Fatalf("failed to get latest version. request failed: %v", err)
	}
	location, err := response.Location()
	if err != nil {
		log.Fatalf("failed to get response location: %v", err)
	}

	return filepath.Base(location.String())
}

func isUpdateAvailable(onlineVersion string) bool {
	return vercheck.Compare(onlineVersion, version, ">")
}

func main() {
	fmt.Println("[+] This is a simple program to test out auto-updating go applications")
	fmt.Println("[+] it will check for an update and if one is available it will apply it")
	fmt.Printf("[+] Current Version: %s\n", version)
	fmt.Println("============================================================================")
	githubVersion := latestTaggedVersion()
	fmt.Printf("[!] Latest tagged version on GitHub: %s\n", githubVersion)
	updateAvailable := isUpdateAvailable(githubVersion)
	update := "Update Available"
	if !updateAvailable {
		update = "Running latest version"
	}
	fmt.Printf("[!] %s\n", update)
}
