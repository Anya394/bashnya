package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInput(filename string) []string {
	var lines []string

	if filename == "" {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}
		return lines
	} else {
		file, err := os.Open(filename)
		if err != nil {
			return nil
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}
		return lines
	}
}

func writeOutput(lines []string, filename string) {
	if filename == "" {
		for _, line := range lines {
			fmt.Println(line)
		}
	} else {
		output, err := os.Create(filename)
		if err != nil {
			return
		}

		for _, line := range lines {
			_, err := fmt.Fprintln(output, line)
			if err != nil {
				return
			}
		}
	}
}
