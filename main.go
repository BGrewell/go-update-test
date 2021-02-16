package main

import "fmt"

var (
	version = "1.0.0"
)

func main() {
	fmt.Println("[+] This is a simple program to test out auto-updating go applications")
	fmt.Println("[+] it will check for an update and if one is available it will apply it")
	fmt.Printf("[+] Current Version: %s\n", version)
	fmt.Println("============================================================================")
}
