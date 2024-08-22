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
	reqsByConcurrency := requests / concurrency
	modulus := requests % concurrency
	var listaReqs []int
	for range concurrency {
		incremento := 0
		if modulus > 0 {
			incremento = 1
			modulus--
		}
		listaReqs = append(listaReqs, reqsByConcurrency+incremento)
	}

	for k := range concurrency {
		go func() {
			for range listaReqs[k] {
				Request(url, status, duration)
			}
		}()
	}

	var statusesCodes []int
	var durations []time.Duration
	for range requests {
		statusesCodes = append(statusesCodes, <-status)
		durations = append(durations, <-duration)
	}
	close(status)
	close(duration)

	return statusesCodes, durations

}
