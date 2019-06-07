package main

import (
    "fmt"
    "strconv"
)

func main() {
    for id := 0; id < 256; id++ {

        code := strconv.Itoa(id)
        //fmt.Printf("%s", code)
        
        fullCode := `\033[38;5;` + code + "m"
        colorCode := "\033[38;5;" + code + "m"
        resetColor := "\033[39;49m"

        fmt.Printf("%sThe code is [%s]%s\n", colorCode, fullCode, resetColor)
    }

}