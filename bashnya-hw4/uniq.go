package main

import (
	"fmt"
	"strings"
)

func processUsingOptions(lines []string, options Options) []string {
	if options.Repeated {
		return findRepeated(lines, options)
	} else if options.Unique {
		return findNotRepeated(lines, options)
	} else if options.Count {
		return countLines(lines, options)
	} else {
		return findUnique(lines, options)
	}
}

func findUnique(lines []string, options Options) []string {
	var result []string
	checkedLines := make(map[string]bool)

	for _, line := range lines {
		key := processLine(line, options)
		if !checkedLines[key] {
			checkedLines[key] = true
			result = append(result, line)
		}
	}
	return result
}

func findNotRepeated(lines []string, options Options) []string {
	counts := countOccurrences(lines, options)
	var result []string

	for _, line := range lines {
		key := processLine(line, options)
		if counts[key] == 1 {
			result = append(result, line)
		}
	}
	return result
}

func findRepeated(lines []string, options Options) []string {
	counts := countOccurrences(lines, options)
	var result []string
	checkedLines := make(map[string]bool)

	for _, line := range lines {
		key := processLine(line, options)
		if counts[key] > 1 && !checkedLines[key] {
			checkedLines[key] = true
			result = append(result, line)
		}
	}
	return result
}

func countLines(lines []string, options Options) []string {
	counts := countOccurrences(lines, options)
	var result []string
	checkedLines := make(map[string]bool)

	for _, line := range lines {
		key := processLine(line, options)
		if !checkedLines[key] {
			checkedLines[key] = true
			result = append(result, fmt.Sprintf("%d %s", counts[key], line))
		}
	}
	return result
}

func countOccurrences(lines []string, options Options) map[string]int {
	counts := make(map[string]int)
	for _, line := range lines {
		key := processLine(line, options)
		counts[key]++
	}
	return counts
}

// Обработка строки согласно параметрам -f, -s, -i
func processLine(line string, options Options) string {
	processed := line

	if options.NumFields > 0 {
		fields := strings.Fields(processed)
		if len(fields) > options.NumFields {
			processed = strings.Join(fields[options.NumFields:], " ")
		} else {
			processed = ""
		}
	}

	if options.NumChars > 0 {
		if len(processed) > options.NumChars {
			processed = processed[options.NumChars:]
		} else {
			processed = ""
		}
	}

	if options.IgnoreCase {
		processed = strings.ToLower(processed)
	}

	return processed
}
