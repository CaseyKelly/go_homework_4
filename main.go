package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
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

// define show method for employee struct type
func (employee *Employee) show() {
	fmt.Println("Name: " + employee.Name)
	fmt.Println("Serial: " + employee.Serial)
	fmt.Println("Age: " + strconv.Itoa(employee.Age))
	fmt.Println("Address: " + employee.Address)
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
	// check each employee struct for the one matching serial from command line flag and show it
	for i := 0; i < (len(employees)); i++ {
		if employees[i].Serial == serial {
			employees[i].show()
		}
	}
}
