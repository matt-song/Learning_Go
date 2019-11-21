package plog

import (
	"fmt"
	"os"
	"time"
)

func plog(logLevel string, message string, logLevel int, enableDEBUG int) {

	// define the color code here:
	lightRed := "\033[38;5;9m"
	red := "\033[38;5;1m"
	green := "\033[38;5;2m"
	yellow := "\033[38;5;3m"
	cyan := "\033[38;5;14m"
	//darkBlue := "\033[38;5;25m"
	normal := "\033[39;49m"

	var colorCode string
	var errorOut = 0

	switch logLevel {
	case "INFO":
		colorCode = green
	case "WARN":
		colorCode = yellow
	case "ERROR":
		colorCode = lightRed
	case "FATAL":
		colorCode = red
		errorOut = 1
	case "DEBUG":
		if enableDEBUG == 1 {
			colorCode = cyan
		} else {
			return
		}
	default:
		colorCode = normal
	}
	curTime := time.Now()
	fmt.Printf(curTime.Format("2006-01-02 15:04:05")+" %s[%s] %s\n", colorCode, logLevel, message)
	if errorOut == 1 {
		os.Exit(1)
	}
}
