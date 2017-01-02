package main

import (
	"fmt"
	"time"
)

// for debug output
var Verbose bool

// get timestamp for logging purposes
func ts() string {
	return time.Now().Format("2006-01-02 15:04:05.999 -0700 MST")
}

// time-stamped printf
func TSPrintf(format string, a ...interface{}) {
	fmt.Printf("%s ", ts())
	fmt.Printf(format, a...)
}

// print debug/status conditionally on having Verbose on
func VPrintf(format string, a ...interface{}) {
	if Verbose {
		TSPrintf(format, a...)
	}
}

func p(format string, a ...interface{}) {
	TSPrintf(format, a...)
}
