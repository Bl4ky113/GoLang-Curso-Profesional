package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	timeStart := time.Now()

	apis := []string{
		"https://management.azure.com",
		"https://dev.azure.com",
		"https://api.github.com",
		"https://outlook.office.com",
		"https://api.notfound",
		"https://graph.microsoft.com",
	}

	channel := make(chan string)
	for _, api := range apis {
		go checkAPI(api, channel)
	}

	for i := 0; i < len(apis); i++ {
		fmt.Println(<-channel)
	}

	elapsed := time.Since(timeStart)
	fmt.Printf("Execution Time: %s", elapsed)
}

func checkAPI(api string, ch chan string) {
	if _, err := http.Get(api); err != nil {
		ch <- fmt.Sprintf("Error API %s is down!", api)
		return
	}

	ch <- fmt.Sprintf("API %s is Working!", api)
}
