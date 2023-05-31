package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

const (
	url      = "https://www.gstatic.com/ipranges/goog.txt"
	filename = "goog.txt"
)

func main() {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error retrieving content: %s\n", err)
		return
	}
	defer resp.Body.Close()

	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf("Error creating file: %s\n", err)
		return
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		fmt.Printf("Error saving content to file: %s\n", err)
		return
	}

	fmt.Println("Content saved to", filename)
}
