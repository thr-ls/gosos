package cmd

import (
	"git.thrls.net/thrls/gosos/output"
	"git.thrls.net/thrls/gosos/storage"
	"git.thrls.net/thrls/gosos/utils"
	"sync"
)

func Run() {
	urlList, err := storage.LoadURLs()
	if err != nil {
		output.PrintError("Error: " + err.Error())
		return
	}

	var wg sync.WaitGroup
	results := make(chan struct {
		url  string
		isUp bool
	}, len(urlList.URLs))

	for _, url := range urlList.URLs {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			isUp := utils.IsUp(url)
			results <- struct {
				url  string
				isUp bool
			}{url, isUp}
		}(url)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	output.PrintInfo("Checking URLs:")
	for result := range results {
		output.PrintURLStatus(result.url, result.isUp)
	}
}
