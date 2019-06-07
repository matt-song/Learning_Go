package main

import (
       "fmt"
//       "net/http"
       "os"
//       "path"
)
const (
	DEBUG  = 1		// enable debug mode
	
)

func main() {

	/* Create the work folder */
	var workFolder string = fmt.Sprintf("/tmp/.testSpeed.%d",os.Getpid()); 
	createWorkFolder(workFolder)

	/* Get the list of site*/


    
}

func createWorkFolder(folder string) {
	
	//var workFolder string = fmt.Sprintf("/proc/.testSpeed.%d",os.Getpid()); 
	printDEBUG("The work folder is " + folder)    
	
	err := os.MkdirAll(folder, 0755) 
	if err != nil { printFATAL(err, "Failed to create folder [" + folder + "], exit!") }

	printINFO("workFolder "+folder+" created")
	defer os.RemoveAll(folder)   // clean the work folder if abnormally exit

	//return folder
}

func printColor(logLevel string, message string){

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
	fmt.Printf("%s[%s] %s%s", colorCode, logLevel, message,normal)
}

func printDEBUG(message string){

	if DEBUG == 1 {
		printColor("DEBUG", message + "\n")
	}
	
}
func printFATAL(err error, message string){
	printColor("FATAL", message + "\n")
	printColor("FATAL", "The error is [" + err.Error() + "]\n")
	os.Exit(1)
}
func printINFO(message string) {
	printColor("INFO", message + "\n")
}