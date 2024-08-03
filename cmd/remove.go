package cmd

import (
	"flag"
	"git.thrls.net/thrls/gosos/output"
	"git.thrls.net/thrls/gosos/storage"
	"git.thrls.net/thrls/gosos/utils"

	"golang.org/x/exp/slices"
)

func Remove(args []string) {
	rmCmd := flag.NewFlagSet("remove", flag.ExitOnError)
	rmCmd.Parse(args)

	if rmCmd.NArg() < 1 {
		output.PrintInfo("Usage: gosos remove <url>")
		return
	}

	url := rmCmd.Arg(0)

	urlList, err := storage.LoadURLs()
	if err != nil {
		output.PrintError("Error: " + err.Error())
		return
	}

	ok := slices.Contains(urlList.URLs, url)
	if !ok {
		output.PrintError("Error: URL does not exist")
		return
	}

	urlList.URLs = utils.RemoveElement(urlList.URLs, url)

	err = storage.SaveURLs(urlList)
	if err != nil {
		output.PrintError("Error: " + err.Error())
		return
	}

	output.PrintSuccess("URL removed from list successfully")
}
