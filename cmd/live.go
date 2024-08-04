package cmd

import (
	"bufio"
	"git.thrls.net/thrls/gosos/output"
	"git.thrls.net/thrls/gosos/storage"
	"git.thrls.net/thrls/gosos/utils"
	"os"
	"time"
)

const (
	updateInterval = 100 * time.Millisecond
	shutdownDelay  = time.Second
)

func Live(interval int) {
	urlList, err := loadURLs()
	if err != nil {
		return
	}

	if err := initializeLiveDisplay(urlList.URLs); err != nil {
		return
	}
	defer output.StopLiveList()

	stopChan := make(chan struct{})
	statusChan := make(chan utils.StatusUpdate, len(urlList.URLs))

	launchMonitors(urlList.URLs, stopChan, statusChan)

	inputChan := listenForUserInput()

	urlIndexMap := createURLIndexMap(urlList.URLs)

	monitorLoop(urlIndexMap, statusChan, inputChan, stopChan)

	shutdown(statusChan)
}

func loadURLs() (*storage.URLList, error) {
	urlList, err := storage.LoadURLs()
	if err != nil {
		output.PrintError("Error loading URLs: " + err.Error())
		return &storage.URLList{}, err
	}
	return urlList, nil
}

func initializeLiveDisplay(urls []string) error {
	if err := output.InitLiveList(urls); err != nil {
		output.PrintError("Error initializing live display: " + err.Error())
		return err
	}
	return nil
}

func launchMonitors(urls []string, stopChan <-chan struct{}, statusChan chan<- utils.StatusUpdate) {
	for _, url := range urls {
		go utils.MonitorStatus(url, stopChan, statusChan)
	}
}

func listenForUserInput() <-chan struct{} {
	inputChan := make(chan struct{})
	go func() {
		bufio.NewReader(os.Stdin).ReadBytes('\n')
		close(inputChan)
	}()
	return inputChan
}

func createURLIndexMap(urls []string) map[string]int {
	urlIndexMap := make(map[string]int, len(urls))
	for i, url := range urls {
		urlIndexMap[url] = i
	}
	return urlIndexMap
}

func monitorLoop(urlIndexMap map[string]int, statusChan <-chan utils.StatusUpdate, inputChan <-chan struct{}, stopChan chan<- struct{}) {
	for {
		select {
		case status := <-statusChan:
			if index, exists := urlIndexMap[status.URL]; exists {
				output.UpdateURLStatus(index, status.URL, status.IsUp)
			}
		case <-inputChan:
			close(stopChan)
			output.PrintWarning("Monitoring stopped. Closing all connections.")
			return
		case <-time.After(updateInterval):
			// This case prevents the select from blocking
		}
	}
}

func shutdown(statusChan chan utils.StatusUpdate) {
	time.Sleep(shutdownDelay)
	close(statusChan)
}
