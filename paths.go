package main

import (
	"os"
	"os/user"
	"path"
)

func getHomeDir() string {
	user, err := user.Current()
	if err != nil {
		writelnStderr("could not determine home directory for cache")
		os.Exit(1)
	}
	return user.HomeDir
}

func getCacheDir() string {
	home := getHomeDir()
	return path.Join(home, ".cache")
}

func getLocalOuiPath() string {
	cache := getCacheDir()
	return path.Join(cache, "oui.txt")
}
