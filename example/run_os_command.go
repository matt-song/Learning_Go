package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

var (
	enableDEBUG = true
)

func main() {

	cmdGood := "uptime"
	runCommand(cmdGood, false)

	cmdBad := "touch /1.txt"
	runCommand(cmdBad, false)

	cmdFATAL := "mv /tmp/1.txt /tmp/2.txt"
	runCommand(cmdFATAL, true)
}

func runCommand(cmd string, errorOut bool) (output string) {

	plog("DEBUG", "Execute command ["+cmd+"]...")

	out, err := exec.Command(cmd).Output()
	outputFinal := strings.TrimSpace(string(out)) // remove the new line at the end
	plog("DEBUG", "The output is: ["+string(outputFinal)+"]")

	if err != nil {
		if errorOut == false {
			plog("ERROR", "Failed to exeute command ["+cmd+"]")
			plog("ERROR", "The error message is ["+err.Error()+"]")
		} else {
			plog("FATAL", "Failed to exeute command ["+cmd+"]")
			plog("FATAL", "The error message is ["+err.Error()+"]")
		}
	}
	return string(outputFinal)
}

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
