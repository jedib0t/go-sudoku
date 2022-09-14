package main

import (
	"fmt"
	"os"
	"strings"
)

func logErrorAndExit(msg string, a ...interface{}) {
	_, _ = fmt.Fprintf(os.Stderr, "ERROR: "+strings.TrimSpace(msg)+"\n", a...)
	cleanup()
	os.Exit(-1)
}
