package main

import (
    "fmt"
    //"strconv"
)

func print(logLevel string, message string){

	// define the color code here:
	lightRed := "\033[38;5;9m"
	red := "\033[38;5;1m"
	green := "\033[38;5;2m"
	yellow := "\033[38;5;3m"
	cyan := "\033[38;5;14m"
	//darkBlue := "\033[38;5;25m"
	normal := "\033[39;49m"

	var colorCode string
	
	switch logLevel {
	case "INFO":
		colorCode = green
	case "WARN":
		colorCode = yellow
	case "ERROR":
		colorCode = lightRed
	case "FATAL":
		colorCode = red
	case "DEBUG":
		colorCode = cyan
	default:
		colorCode = normal
	}
	fmt.Printf("%s[%s] %s", colorCode, logLevel, message)
}

func main(){
    print("INFO","This is the messgae with log level [INFO]\n")
    print("DEBUG","This is the messgae with log level [DEBUG]\n")
    print("WARN","This is the messgae with log level [WARN]\n")
    print("ERROR","This is the messgae with log level [ERROR]\n")
    print("FATAL","This is the messgae with log level [FATAL]\n")
	print("UNKNOWN","This is the messgae with log level [UNKNOWN]\n")
}

