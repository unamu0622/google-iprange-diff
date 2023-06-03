package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

const (
	url      = "https://www.gstatic.com/ipranges/goog.txt"
	fileName = "goog.txt"
)

func main() {
	err := downloadFile(url, fileName)
	if err != nil {
		log.Fatal(err)
	}
}

func downloadFile(url string, fileName string) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error retrieving content: %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error: Response status code is %d", resp.StatusCode)
	}

	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("error creating file: %s", err)
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return fmt.Errorf("error saving content to file: %s", err)
	}

	fmt.Println("Content saved to", fileName)
	return nil
}
