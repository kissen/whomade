package main

import (
	"github.com/klauspost/oui"
	"io"
	"net/http"
	"os"
)

// Ensure that oui.txt exists at the expected path.
func updateOuiCache() {
	const OUI_URL = "http://standards-oui.ieee.org/oui.txt"

	writelnStderr("building MAC database cache...")

	// Fetch data
	reply, err := http.Get(OUI_URL)
	if err != nil {
		writelnStderr("could not fetch MAC database from IEEE")
		os.Exit(1)
	}
	defer reply.Body.Close()

	// Create file for writing
	f, err := os.Create(getLocalOuiPath())
	if err != nil {
		writelnStderr("could not create MAC database cache")
		os.Exit(1)
	}
	defer f.Close()

	// Write cache
	_, err = io.Copy(f, reply.Body)
	if err != nil {
		writelnStderr("could not write MAC database cache")
		os.Exit(1)
	}
}

func createDatabase() oui.OuiDB {
	// First, try opening from cache
	db, err := oui.OpenStaticFile(getLocalOuiPath())

	if err != nil {
		// Cache does not exist or is corrupt, try re-downloading
		updateOuiCache()
		db, err = oui.OpenStaticFile(getLocalOuiPath())

		if err != nil {
			writelnStderr("could not load MAC database")
			os.Exit(1)
		}
	}

	return db
}
