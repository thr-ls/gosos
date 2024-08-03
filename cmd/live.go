package cmd

import (
	"bufio"
	"git.thrls.net/thrls/gosos/output"
	"git.thrls.net/thrls/gosos/storage"
	"git.thrls.net/thrls/gosos/utils"
	"os"
	"time"
)

func Live(interval int) {
	urlList, err := storage.LoadURLs()
	if err != nil {
		output.PrintError("Error: " + err.Error())
		return
	}

	err = output.InitLiveList(urlList.URLs)
	if err != nil {
		output.PrintError("Error initializing live display: " + err.Error())
		return
	}
	defer output.StopLiveList()

	stopChan := make(chan struct{})
	statusChan := make(chan utils.StatusUpdate)

	for _, url := range urlList.URLs {
		go utils.MonitorStatus(url, stopChan, statusChan)
	}

	inputChan := make(chan struct{})
	go func() {
		bufio.NewReader(os.Stdin).ReadBytes('\n')
		close(inputChan)
	}()

	urlIndexMap := make(map[string]int)
	for i, url := range urlList.URLs {
		urlIndexMap[url] = i
	}

	monitoringActive := true
	for monitoringActive {
		select {
		case status := <-statusChan:
			index, exists := urlIndexMap[status.URL]
			if exists {
				output.UpdateURLStatus(index, status.URL, status.IsUp)
			}
		case <-inputChan:
			close(stopChan)
			output.PrintWarning("Monitoring stopped. Closing all connections.")
			monitoringActive = false
		case <-time.After(100 * time.Millisecond):
		}
	}

	time.Sleep(time.Second)
	close(statusChan)
}
