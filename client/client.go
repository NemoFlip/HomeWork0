package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	resp, err := http.Get("http://127.0.0.1:8080")
	if err != nil {
		fmt.Printf("error sending get request: %s", err.Error())
		return
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("error reading response: %s", err.Error())
		return
	}
	if string(data) == "Ok" {
		fmt.Println("GET request was successful!")
	} else {
		fmt.Println("Unknown response word.")
	}
}
