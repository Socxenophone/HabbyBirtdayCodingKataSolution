// package cli implements solution
package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

var name string

type Person struct {
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Birth     time.Time `json:"birth"`
	Email     string    `json:"email"`
}

func (p *Person) wish() {
	fmt.Printf("Subject: Happy birthday!\nHappy birthday, dear %s\n!", p.FirstName)
}
func (p *Person) update(lastName string, firstName string, birth time.Time, email string) {
	p.LastName = lastName
	p.FirstName = firstName
	p.Birth = birth
	p.Email = email
}

func main() {
	var names string
	fmt.Println("Enter the file containing names:")
	fmt.Scanf(names)
	values, err := readFile(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(values)
}

// function readFile() reads from a text file and returns a string slice and error value
func readFile(filename string) ([]string, error) {
	var values []string

	// Check if the filename already has ".txt" or ".json" extension
	if !strings.HasSuffix(filename, ".txt") {
		filename += ".txt"
	} else if strings.HasSuffix(filename, ".json") {
		readJson(filename)
	}

	//Open the file
	f, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("Error opening file %v", err)

	}
	defer f.Close()

	//Read and split to slices
	reader := bufio.NewReader(f)
	content, err := reader.ReadString('\n')
	if err != nil {
		return nil, fmt.Errorf("Error reading file: %v", err)

	}
	if values == nil {
		return nil, errors.New("File is empty")
	}

	values = strings.Split(content, ",")
	return values, nil
}

// function readJson() reads from a json file and returns a string slice and error value
func readJson(filename string) (s []string, err error) {
	var values []string
	// Open the file
	f, err := os.Open(filename)
	if err != nil {
		return nil, errors.New("Failed opening .json file")
	}

	//Declare need decoder, decode the json into values slice
	decoder := json.NewDecoder(f)
	err = decoder.Decode(&values)
	if err != nil {
		return nil, errors.New("Failed to decode .json file")
	}
	return values, nil
}
