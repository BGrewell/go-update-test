package main

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"errors"
	"fmt"
	vercheck "github.com/mcuadros/go-version"
	"github.com/inconshreveable/go-update"
	"io"
	"log"
	"net/http"
	"path/filepath"
	"strings"
)

var (
	version = "DEBUG"
	tagCheckUrl = "https://github.com/BGrewell/go-update-test/releases/latest"
	updateUrlStub = "https://github.com/BGrewell/go-update-test/releases/download/%VER1%/go-update-test_%VER2%_Linux_x86_64.tar.gz"
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

func getChecksums() string {
	return ""
}

func getPackage(version string) io.Reader {
	url := strings.ReplaceAll(updateUrlStub, "%VER1%", version)
	url = strings.ReplaceAll(url, "%VER2%", strings.Replace(version, "v", "", 1))
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("failed to get update package: %v\n", err)
	}

	defer resp.Body.Close()
	gzr, err := gzip.NewReader(resp.Body)
	if err != nil {
		log.Fatalf("failed to remove gzip encoding: %v\n", err)
	}
	tr := tar.NewReader(gzr)
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("failed to extract binary from package: %v\n", err)
		}
		fmt.Printf("file: %s\n", hdr.Name)
		if !strings.Contains(hdr.Name, ".md") {
			var filebuf bytes.Buffer
			io.TeeReader(tr, &filebuf)
			return &filebuf
		}
	}

	return nil
}

func main() {
	fmt.Println("[+] This is a simple program to test out auto-updating go applications")
	fmt.Println("[+] it will check for an update and if one is available it will apply it")
	fmt.Printf("[+] Current Version: %s\n", version)
	fmt.Println("============================================================================")
	githubVersion := latestTaggedVersion()
	fmt.Printf("[!] Latest tagged version on GitHub: %s\n", githubVersion)
	updateAvailable := isUpdateAvailable(githubVersion)
	updateStatus := "Update Available"
	if !updateAvailable {
		updateStatus = "Running latest version"
	}
	fmt.Printf("[!] %s\n", updateStatus)
	if updateAvailable {
		binary := getPackage(githubVersion)
		options := update.Options{}
		err := update.Apply(binary, options)
		if err != nil {
			log.Fatalf("failed to apply update: %v", err)
		}
		fmt.Println("[!] Applied update")
	}
}
