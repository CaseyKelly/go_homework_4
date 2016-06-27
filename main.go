package main

import (
	"flag"
)

func main() {
	// Capture and parse flags from command line
	serial := flag.String("serial", "0", "employee serial number")
	filename := flag.String("filename", "0", "filename")
	flag.Parse()

	// Call function to fetch employee record
	readFile(*filename, *serial)
}

func readFile(filename string, serial string) {

}
