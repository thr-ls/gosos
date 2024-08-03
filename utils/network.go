package utils

import (
	"net/http"
	"time"
)

func IsUp(url string) bool {
	resp, err := http.Get(url)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	return resp.StatusCode >= 200 && resp.StatusCode < 300
}

type StatusUpdate struct {
	URL  string
	IsUp bool
}

func MonitorStatus(url string, stop <-chan struct{}, status chan<- StatusUpdate) {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	checkAndSend(url, status)

	for {
		select {
		case <-stop:
			return
		case <-ticker.C:
			checkAndSend(url, status)
		}
	}
}

func checkAndSend(url string, status chan<- StatusUpdate) {
	status <- StatusUpdate{URL: url, IsUp: IsUp(url)}
}
