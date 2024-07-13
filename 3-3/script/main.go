package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

const baseURL = "http://localhost:8080/articles"

func main() {
	// 1. Create a new article
	fmt.Println("1. Create a new article")
	postData := "title=New Article&content=Content of the new article&author=Alice"
	resp, err := http.Post(baseURL, "application/x-www-form-urlencoded", bytes.NewBufferString(postData))
	handleResponse(resp, err)

	// 2. Read the created article
	fmt.Println("2. Read the created article")
	resp, err = http.Get(baseURL + "/1")
	handleResponse(resp, err)

	// 3. Read a non-existing article (should return not found)
	fmt.Println("3. Read a non-existing article (should return not found)")
	resp, err = http.Get(baseURL + "/999")
	handleResponse(resp, err)

	// 4. Create an article with empty title
	fmt.Println("4. Create an article with empty title (should return error)")
	postData = "title=&content=Some content&author=Alice"
	resp, err = http.Post(baseURL, "application/x-www-form-urlencoded", bytes.NewBufferString(postData))
	handleResponse(resp, err)

	// 5. Create an article with empty content (should return error)
	fmt.Println("5. Create an article with empty content (should return error)")
	postData = "title=Some Title&content=&author=Alice"
	resp, err = http.Post(baseURL, "application/x-www-form-urlencoded", bytes.NewBufferString(postData))
	handleResponse(resp, err)

	// 6. Create an article with empty author (should return error)
	fmt.Println("6. Create an article with empty author (should return error)")
	postData = "title=Some Title&content=Some content&author="
	resp, err = http.Post(baseURL, "application/x-www-form-urlencoded", bytes.NewBufferString(postData))
	handleResponse(resp, err)

	// 7. Read all articles to check the final state
	fmt.Println("7. Read all articles to check the final state")
	resp, err = http.Get(baseURL + "/1")
	handleResponse(resp, err)
	resp, err = http.Get(baseURL + "/2")
	handleResponse(resp, err)
	resp, err = http.Get(baseURL + "/3")
	handleResponse(resp, err)
}

func handleResponse(resp *http.Response, err error) {
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	fmt.Println("Response:", string(body))
	fmt.Println()
}
