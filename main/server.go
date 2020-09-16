package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/api/thumbnail", thumbnailHandler)

	fs := http.FileServer(http.Dir("./frontend/dist"))
	http.Handle("/", fs)

	fmt.Println("Server Listening on port 3000")
	log.Panic(
		http.ListenAndServe(":3000", nil),
	)
}

type screenshotAPIRequest struct {
	Token          string `json:"token"`
	Url            string `json:"url"`
	Output         string `json:"output"`
	Width          int    `json:"width"`
	Height         int    `json:"height"`
	ThumbnailWidth int    `json:"thumbnail_width`
}

type screenshotAPIResponse struct {
	Screenshot string `json"screenshot"`
}

type thumbnailRequest struct {
	Url string `json:"url"`
}

func thumbnailHandler(w http.ResponseWriter, r *http.Request) {
	var decoded thumbnailRequest

	// Try to decode the request into the thumbnailRequest
	if err := json.NewDecoder(r.Body).Decode(&decoded); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	apiRequest := screenshotAPIRequest{
		Token:          "MVZJWIAQWGKBTSXXRNU7CFT9TBSRGI12",
		Url:            decoded.Url,
		Output:         "json",
		Width:          1920,
		Height:         1080,
		ThumbnailWidth: 300,
	}

	// convert the struct to json
	jsonString, err := json.Marshal(apiRequest)
	checkError(err)

	// Create HTTP Request
	req, err := http.NewRequest("POST", "https://screenshotapi.net/api/v1/screenshot", bytes.NewBuffer(jsonString))
	req.Header.Set("Content-Type", "application/json")

	// Execute the HTTP request
	client := &http.Client{}
	response, err := client.Do(req)
	checkError(err)

	// Close the response
	defer response.Body.Close()

	var apiResponse screenshotAPIResponse
	err = json.NewDecoder(response.Body).Decode(&apiResponse)
	checkError(err)

	_, err = fmt.Fprintf(w, `{"screenshot": "%s"}`, apiResponse.Screenshot)
	checkError(err)
}

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hello World")
	checkError(err)

}

func checkError(err error) {
	if err != nil {
		log.Panic(err)
	}
}
