/*
Parse lines of a text file separated with line breakers (\n) converting them to array elements of a json file.
*/
package main

import (
	"os"
	"bufio"
	"io"
	"log"
	"encoding/json"
	"flag"
)

func main() {
	var (
		verbose         bool
		input_file, output_file  string
	)

	// Command line arguments parser
	flag.BoolVar(&verbose, "verbose", false, "Print each line from input file while parsing it.")
	flag.StringVar(&input_file, "input", "input.txt", "Name/Path of the input file  (default is input.txt).")
	flag.StringVar(&output_file, "output", "output.json", "Name/Path of the output json file (default is output.json).")
	flag.Parse()

	txtContent := make(map[string]map[string]string)
	file, _ := os.Open(input_file)
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				log.Println("Error in parsing :", err)
			}
			break
		}
		// If verbose is set print each line while parsing it
		if verbose {
			log.Println(line)
		}

	}
	file, err := os.Create(output_file)
	if err != nil {
		log.Println("Error :", err)
		return
	}
	defer file.Close()
	jsonStr, _ := json.Marshal(txtContent)
	file.Write(jsonStr)

	log.Println(output_file, "file created with success")

}
