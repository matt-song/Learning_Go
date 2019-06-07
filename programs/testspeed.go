package main

import (
       "fmt"
//       "net/http"
       "os"
)



func main() {


	

    fmt.Println("Start to testing speed for VPS...")
    createWorkFolder()
}

func createWorkFolder(){

	// had to use this way so the output can be string, then we can combine it with folder name
	var workFolder string = fmt.Sprintf("/tmp/.testSpeed.%d",os.Getpid()) 
	fmt.Println("[DEBUG] The work folder is ", workFolder)    
	
	err := os.MkdirAll(workFolder, 0755) 
	fmt.Println("[DEBUG] Error:", err)    
/*	if err != nil {
		log.Fatal(err)
	}
*/
}
