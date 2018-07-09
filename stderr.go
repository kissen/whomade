package main

import (
	"fmt"
	"os"
)

func writelnStderr(s string) {
	msg := fmt.Sprintf("%s: %s\n", os.Args[0], s)
	os.Stderr.WriteString(msg)
}
