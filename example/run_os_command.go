package main

import (
    "fmt"
//       "net/http"
    "os"
    "os/exec"

    "strings"
)

func main() {


	ping := "ping -c4 sgp-ping.vultr.com | tail -1 | awk -F'/' '{print $5}' "
	runCommand(ping)


}

func runCommand(cmd string) (output string){
        
	printDEBUG("Execute command ["+cmd+"]...")

	out, err := exec.Command(cmd).Output()
	outputFinal := strings.TrimSpace(string(out))		// remove the new line at the end
	printDEBUG("The output is: ["+ string(outputFinal)+"]")

	if err != nil {
		printERROR(err, "Failed to exeute command ["+cmd+"]")
	}
	return string(outputFinal)

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
	DEBUG := 1
	if DEBUG == 1 {
		printColor("DEBUG", message + "\n")
	}
	
}
func printERROR(err error, message string){
	printColor("ERROR", message + "\n")
	printColor("ERROR", "The error is [" + err.Error() + "]\n")
	os.Exit(1)
}
func printFATAL(err error, message string){
	printColor("FATAL", message + "\n")
	printColor("FATAL", "The error is [" + err.Error() + "]\n")
	os.Exit(1)
}
func printINFO(message string) {
	printColor("INFO", message + "\n")
}