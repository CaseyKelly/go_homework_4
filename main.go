package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
)

func main() {
	// Capture and parse flags from command line
	serial := flag.String("serial", "0", "employee serial number")
	filename := flag.String("filename", "0", "filename")
	flag.Parse()
	// Call function to fetch employee record
	readFile(*filename, *serial)
}

// define employee struct
type Employee struct {
	Name    string
	Serial  string
	Age     int
	Address string
}

func readFile(filename string, serial string) {
	// read json file and save it to data variable
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	// take slice of data and unmarshal it into slice of employee structs
	jsonBlob := []byte(data)
	var employees []Employee
	err = json.Unmarshal(jsonBlob, &employees)
	if err != nil {
		fmt.Println("error:", err)
	}
	// check each employee struct for the one matching serial from command line flag
	for i := 0; i < (len(employees)); i++ {
		if employees[i].Serial == serial {
			fmt.Printf("%+v\n", employees[i])
		}
	}
}
