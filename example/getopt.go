package main

import (
	"fmt"
	"os"

	"github.com/pborman/getopt/v2"
)

var (
	filePath = "/tmp" // default vaule set here
)

func main() {

	// StringLong: func StringLong(name string, short rune, value string, helpvalue ...string) *string {
	//	..return &value   << return address. call the value by *xxx
	theName := getopt.StringLong("name", 'n', "", "Your name")               // --name=xxx or -n xxx, no default value
	theCommand := getopt.StringLong("command", 'c', "uptime", "the command") // --command=xxx or -c xxx, default value 'uptime'

	// BoolLong: func BoolLong(name string, short rune, helpvalue ...string) *bool {
	optHelp := getopt.BoolLong("help", 0, "Help") // --help or -h

	//Bool: func Bool(name rune, helpvalue ...string) *bool {
	// return &b  << call result via *xxxx
	enableDebug := getopt.Bool('D', "display debug message") // -D

	// flag: func Flag(v interface{}, short rune, helpvalue ...string) Option {
	getopt.Flag(&filePath, 'f', "The path of the file") // -f xxxx, change the value of filePath

	getopt.Parse()

	if *optHelp { // need dereference!!!
		getopt.Usage()
		os.Exit(0)
	}

	fmt.Printf("summary: \n    name: [%s] \n    command: [%s]\n    filePath: [%s] \n    enableDebug: [%t] \n", *theName, *theCommand, filePath, *enableDebug)
	if *enableDebug == true {
		fmt.Printf("this is some debug message since -D has enabled")
	}

}
