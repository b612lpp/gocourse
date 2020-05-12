//Here is an example from internet. It generates a file and fils two string with divider.
package main

import (
	"encoding/csv"
	"log"
	"os"
)

var data = [][]string{{"Line1", "Hello"}, {"Line2", "world"}}

func main() {
	file, err := os.Create("result.csv")
	checkError("Cannot create file", err)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, value := range data {
		err := writer.Write(value)
		checkError("Cannot write to file", err)
	}
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}
