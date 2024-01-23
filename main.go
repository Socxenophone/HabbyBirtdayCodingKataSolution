// You can edit this code!
// Click here and start typing.
package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

var name string

type Person struct {
	LastName  string
	FirstName string
	birth     time.Time
	email     string
}

func (p *Person) wish() {
	fmt.Printf("Subject: Happy birthday!\nHappy birthday, dear %s\n!", p.FirstName)

}
func main() {

	fmt.Println("Hello, 世界")
	fmt.Scanf("Filename:", &name)
	values, err := readFile(name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(values)
}

// function readFile() reads from a text file and returns a string slice
func readFile(filename string) ([]string, error) {
	var values []string

	// Check if the filename already has ".txt" extension
	if !strings.HasSuffix(filename, ".txt") {
		filename += ".txt"
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
