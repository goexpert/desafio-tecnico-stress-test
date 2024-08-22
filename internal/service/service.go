package service

import (
	"fmt"
	"net/http"
	"time"
)

func Request(url string, status chan int, duration chan time.Duration) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	duration <- time.Since(start)
	status <- resp.StatusCode

}

func ConcurrentRequests(url string, concurrency, requests int) ([]int, []time.Duration) {
	status := make(chan int, concurrency)
	duration := make(chan time.Duration, concurrency)

	for range concurrency {
		go func() {
			Request(url, status, duration)
		}()
	}

	var statusesCodes []int
	var durations []time.Duration
	for range concurrency {
		statusesCodes = append(statusesCodes, <-status)
		durations = append(durations, <-duration)
	}
	close(status)
	close(duration)

	return statusesCodes, durations

}
