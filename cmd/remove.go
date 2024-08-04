package cmd

import (
	"flag"
	"fmt"
	"git.thrls.net/thrls/gosos/output"
	"git.thrls.net/thrls/gosos/storage"
	"git.thrls.net/thrls/gosos/utils"

	"golang.org/x/exp/slices"
)

func Remove(args []string) {
	url, err := parseRemoveArgs(args)
	if err != nil {
		output.PrintError(err.Error())
		return
	}

	urlList, err := storage.LoadURLs()
	if err != nil {
		output.PrintError("Error loading URLs: " + err.Error())
		return
	}

	if err := removeURLFromList(urlList, url); err != nil {
		output.PrintError(err.Error())
		return
	}

	if err := storage.SaveURLs(urlList); err != nil {
		output.PrintError("Error saving URL list: " + err.Error())
		return
	}

	output.PrintSuccess("URL removed from list successfully")
}

func parseRemoveArgs(args []string) (string, error) {
	rmCmd := flag.NewFlagSet("remove", flag.ExitOnError)
	rmCmd.Parse(args)

	if rmCmd.NArg() < 1 {
		return "", fmt.Errorf("insufficient arguments\nUsage: gosos remove <url>")
	}

	return rmCmd.Arg(0), nil
}

func removeURLFromList(urlList *storage.URLList, url string) error {
	if !slices.Contains(urlList.URLs, url) {
		return fmt.Errorf("URL does not exist in the list")
	}

	urlList.URLs = utils.RemoveElement(urlList.URLs, url)
	return nil
}
