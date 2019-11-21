package main

import (
	"fmt"
	"os"
	"time"
)

const (
	enableDEBUG = true
)

func plog(logLevel string, message string) {

	// define the color code here:
	lightRed := "\033[38;5;9m"
	red := "\033[38;5;1m"
	green := "\033[38;5;2m"
	yellow := "\033[38;5;3m"
	cyan := "\033[38;5;14m"
	//darkBlue := "\033[38;5;25m"
	normal := "\033[39;49m"

	var colorCode string
	var errorOut = false

	switch logLevel {
	case "INFO":
		colorCode = green
	case "WARN":
		colorCode = yellow
	case "ERROR":
		colorCode = lightRed
	case "FATAL":
		colorCode = red
		errorOut = true
	case "DEBUG":
		if enableDEBUG == true {
			colorCode = cyan
		} else {
			return
		}
	default:
		colorCode = normal
	}
	curTime := time.Now()
	fmt.Printf("%s"+curTime.Format("2006-01-02 15:04:05")+" [%s] %s\n", colorCode, logLevel, message)
	if errorOut == true {
		os.Exit(1)
	}
}

func main() {

	var msg = "This is the messgae"
	plog("INFO", msg)
	plog("DEBUG", msg)
	plog("WARN", msg)
	plog("ERROR", msg)
	plog("UNKNOWN", msg)
	plog("FATAL", msg)
}
