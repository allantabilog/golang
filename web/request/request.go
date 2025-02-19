package main

import (
	"fmt"
	"net/http"
)

func main() {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "http://www.google.com", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	fmt.Println(resp.Status)
	fmt.Println("Headers:")
	for key, value := range resp.Header {
		fmt.Println(key, ":", value)
	}
}	

func mainOriginal() {
	resp, err := http.Get("http://www.google.com/")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp.Status)
}