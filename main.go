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
}

func parseArgs() callConfig {
	macs := []string{}
	printHelp := false
	printVersion := false

	for _, arg := range os.Args[1:] {
		if arg == "-v" || arg == "--version" {
			printVersion = true
			continue
		}

		if arg == "-h" || arg == "--help" {
			printHelp = true
			continue
		}

		macs = append(macs, arg)
	}

	if len(macs) == 0 {
		printHelp = true
	}

	return callConfig{
		macs:         macs,
		printHelp:    printHelp,
		printVersion: printVersion,
	}
}

func printHelp() {
	usage := fmt.Sprintf("usage: %s [OPTION] [MAC ADDRESS]...\n", os.Args[0])

	os.Stderr.WriteString(usage)
	os.Stderr.WriteString("\n")
	os.Stderr.WriteString("-h --help      print this usage information\n")
	os.Stderr.WriteString("-v --version   print version information\n")
}

func printVersion() {
	fmt.Println("whomade 0.1")
}

func handleMac(db oui.OuiDB, arg *string) {
	entry, err := db.Query(*arg)

	if err != nil {
		msg := fmt.Sprintf("%s\tnot assigned or invalid'\n", *arg)
		os.Stderr.WriteString(msg)
	} else {
		addr := entry.Prefix
		org := entry.Manufacturer

		fmt.Printf("%s\t%s\n", addr, org)
	}
}

func main() {
	conf := parseArgs()

	// printing help and version

	if conf.printHelp {
		printHelp()
		os.Exit(0)
	}

	if conf.printVersion {
		printVersion()
		os.Exit(0)
	}

	// actual lookup

	db := createDatabase()

	for _, arg := range conf.macs {
		handleMac(db, &arg)
	}
}
