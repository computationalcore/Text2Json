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
		input_file, output_file  string
	)

	flag.StringVar(&input_file, "input", "input.txt", "Name/Path of the input file  (default is input.txt).")
	flag.StringVar(&output_file, "output", "output.json", "Name/Path of the output json file (default is output.json).")

	txtContent := make(map[string]map[string]string)
	f, _ := os.Open(input_file)
	r := bufio.NewReader(f)
	var currentPage map[string]string
	pageNum := 0
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				log.Println("Error in parsing :", err)
			}
			break
		}
	}
	f, err := os.Create(output_file)
	if err != nil {
		log.Println("Error :", err)
		return
	}
	defer f.Close()
	bout, _ := json.Marshal(txtContent)
	f.Write(bout)
}
