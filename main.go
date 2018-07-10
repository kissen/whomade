package main

import (
	"fmt"
	"github.com/klauspost/oui"
	"os"
)

type callConfig struct {
	macs         []string
	printHelp    bool
	printVersion bool
	updateDB     bool
}

func (cc *callConfig) isValid() bool {
	return len(cc.macs) > 0 || cc.printVersion || cc.updateDB
}

func parseArgs() callConfig {
	macs := []string{}
	printHelp := false
	printVersion := false
	updateDB := false

	for _, arg := range os.Args[1:] {
		if arg == "-v" || arg == "--version" {
			printVersion = true
			continue
		}

		if arg == "-h" || arg == "--help" {
			printHelp = true
			continue
		}

		if arg == "-u" || arg == "--update-db" {
			updateDB = true
			continue
		}

		macs = append(macs, arg)
	}

	return callConfig{
		macs:         macs,
		printHelp:    printHelp,
		printVersion: printVersion,
		updateDB:     updateDB,
	}
}

func printHelp() {
	usage := fmt.Sprintf("usage: %s [OPTION] [MAC ADDRESS]...\n", os.Args[0])

	os.Stderr.WriteString(usage)
	os.Stderr.WriteString("\n")
	os.Stderr.WriteString("-u --update-db  update the MAC database\n")
	os.Stderr.WriteString("-h --help       print this usage information\n")
	os.Stderr.WriteString("-v --version    print version information\n")
}

func printVersion() {
	fmt.Println("whomade 0.2")
}

func handleMac(db oui.OuiDB, arg *string) {
	entry, err := db.Query(*arg)

	if err != nil {
		msg := fmt.Sprintf("%s\tnot assigned or invalid\n", *arg)
		os.Stderr.WriteString(msg)
	} else {
		addr := entry.Prefix
		org := entry.Manufacturer

		fmt.Printf("%s\t%s\n", addr, org)
	}
}

func main() {
	conf := parseArgs()

	if !conf.isValid() {
		printHelp()
		os.Exit(1)
	}

	// printing help and version

	if conf.printHelp {
		printHelp()
		os.Exit(0)
	}

	if conf.printVersion {
		printVersion()
		os.Exit(0)
	}

	if conf.updateDB {
		updateOuiCache()
	}

	// actual lookup

	db := createDatabase()

	for _, arg := range conf.macs {
		handleMac(db, &arg)
	}
}
