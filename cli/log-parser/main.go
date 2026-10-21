package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	filename := "cli/log-parser/sample.log"

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	counts := map[string]int{
		"INFO":  0,
		"WARN":  0,
		"ERROR": 0,
	}

	lineNumber := 0

	for scanner.Scan() {
		lineNumber++

		line := scanner.Text()

		fmt.Printf("%d: %s\n", lineNumber, line)

		switch {
		case strings.HasPrefix(line, "INFO"):
			counts["INFO"]++
		case strings.HasPrefix(line, "WARN"):
			counts["WARN"]++
		case strings.HasPrefix(line, "ERROR"):
			counts["ERROR"]++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println("\n--- Log Summary ---")
	fmt.Println("INFO:", counts["INFO"])
	fmt.Println("WARN:", counts["WARN"])
	fmt.Println("ERROR:", counts["ERROR"])
}
