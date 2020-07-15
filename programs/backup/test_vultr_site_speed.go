package main

import (
	"fmt"
    "strconv"
	"os"
	"os/exec"
	"errors"
	//"github.com/davecgh/go-spew/spew"
	"strings"
	"regexp"


)
const (
	DEBUG  = 0		// enable debug mode	
)

type report struct {
	latency string
	speed string 
}

func main() {

	/* Create the work folder */
	var workFolder string = fmt.Sprintf("/tmp/.testSpeed.%d",os.Getpid()); 
	createWorkFolder(workFolder)

	/* Get the list of site*/ 
	cmd := "curl -s https://www.vultr.com/resources/faq/#downloadspeedtests | grep 100MB | awk -F 'href=' '{print $2}' | grep -v ipv6 | grep https | awk '{print $1}' | sed 's/\\\"//g'"
	curlOutput := runCommand(cmd)
	if len(curlOutput) == 0 {
		printFATAL(errors.New("curl command failed"), "Unable to get site list from vultr.com, please check the URL!")
	}
	var allSite []string = strings.Split(curlOutput, "\n")

	/* Test the speed and generate the report */
	testingReport := runSpeedTest(allSite)

	/* Generate the report */
	generateReport(testingReport)
}

func generateReport(rawReport map[string]report){

	cyan := "\033[38;5;14m"
	normal := "\033[39;49m"

	/* generate header */
	fmt.Printf("%s\n============================================== Vultr site speed test report ==============================================%s\n\n", cyan,normal)
	for site, report := range rawReport {
		fmt.Printf("%sHost: [%s]  \tlatency: [%s]\t  speed: [%s] %s\n", cyan,site,report.latency,report.speed,normal)
	}
	fmt.Printf("\n%s======================================================== End  ===========================================================%s\n", cyan,normal)
}

func runSpeedTest(allSite []string)  map[string]report{
    
    testDuration := 5 					// how long will be used for download the file 
	fullReport := make(map[string]report)	// create a hash for hold the full report
	var siteReport report

	printINFO("Start the speed test...")
	for siteID := 0; siteID < len(allSite); {

		host := strings.Split(allSite[siteID], "/")[2]
		printDEBUG("Testing speed on site: [" + host +"]")

		/* --------------- latency test ---------------*/
		// step#1: ping the host
		
		printDEBUG("Ping test against host ["+host+"]...")
		pingSummary := runCommand("ping -c3 -q " + host)

		// step#2: get the latency
		latencyValue := strings.Split(strings.Split(strings.Split(pingSummary, "\n")[4],"/")[4],".")[0]
		printDEBUG("The latency to host is: " + string(latencyValue) )
		
		if len(latencyValue) == 0 {
			printWARN("unable to get latency from host [" + host +"]")
			latencyValue = "Unknown"
		}
		
		siteReport.latency = latencyValue + " ms"; 

		/* --------------- download test ---------------*/
		
		// download a 100MB file for 10 seconds
		printINFO("Testing download speed from url [" +allSite[siteID]+ "]")
		var fileSize = 100 * 1024 * 1024			// 100 MB
		var downloadedSize int
		getFileCMD := "curl " +allSite[siteID]+ " -o /dev/null -m " + strconv.Itoa(testDuration) +" 2>&1 | grep 'Operation timed out'"
		downloadSummary := runCommand(getFileCMD)
		
		if len(downloadSummary) == 0 {				// download finished within 10 sec
			siteReport.speed = string(fileSize/testDuration/1024) + " KB/s"
		}else{
			downloadedSizeString := strings.Split(downloadSummary, " ")[9]
			
			// check if the output is number
			match, _ := regexp.MatchString("([0-9]+)", downloadedSizeString)
			if match {
				downloadedSize, _ = strconv.Atoi(downloadedSizeString)
				siteReport.speed = strconv.Itoa(downloadedSize/testDuration/1024) + " KB/s"
			} else {
				siteReport.speed = "Unknown"
			}
		}
		printDEBUG("Download speed is ["+ siteReport.speed +"]")
		fullReport[allSite[siteID]] = siteReport

		//spew.Dump(siteReport)
		//panic("test")

		siteID++
	}
	//spew.Dump(fullReport)
	return fullReport
}

func runCommand(cmd string) (output string){
        
	printDEBUG("Execute command ["+cmd+"]...")
	out, err := exec.Command("bash","-c",cmd).Output()
	outputFinal := strings.TrimSpace(string(out))		// remove the new line at the end
	printDEBUG("The output is: ["+ string(outputFinal)+"]")

	if err != nil {
		printERROR(err, "Failed to exeute command ["+cmd+"]")
	}
	return string(outputFinal)
}

func createWorkFolder(folder string) {
	
	//var workFolder string = fmt.Sprintf("/proc/.testSpeed.%d",os.Getpid()); 
	printDEBUG("The work folder is " + folder)    
	
	err := os.MkdirAll(folder, 0755) 
	if err != nil { printFATAL(err, "Failed to create folder [" + folder + "], exit!") }

	printINFO("workFolder "+folder+" created")
	defer os.RemoveAll(folder)   // clean the work folder if abnormally exit
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
func printWARN(message string) {
	printColor("WARN", message + "\n")
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