package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const layout = "02/01/2006 15:04:05"

func main() {
	if len(os.Args) < 3 {
		printUsage()
		return
	}

	flag := os.Args[1]

	switch flag {
	case "--human", "-h":
		nsStr := os.Args[2]
		ns, err := strconv.ParseInt(nsStr, 10, 64)
		if err != nil {
			fmt.Printf("Invalid nanoseconds: %v\n", err)
			os.Exit(1)
		}

		// Convert to time, set to system's local timezone, and format
		t := time.Unix(0, ns).Local()
		fmt.Println(t.Format(layout))

	case "--nano", "-n":
		// Join args to handle unquoted input
		dateStr := strings.Join(os.Args[2:], " ")

		// Parse the time using the system's local timezone
		t, err := time.ParseInLocation(layout, dateStr, time.Local)
		if err != nil {
			fmt.Println("Invalid date format. Expected: DD/MM/YYYY HH:MM:SS")
			os.Exit(1)
		}

		fmt.Println(t.UnixNano())

	default:
		printUsage()
	}
}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  ns -h, --human <nanoseconds>           Convert nanoseconds to human date")
	fmt.Println("  ns -n, --nano  <DD/MM/YYYY HH:MM:SS>   Convert human date to nanoseconds")
}
