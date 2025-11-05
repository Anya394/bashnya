package main

import (
	"flag"
	"fmt"
	"os"
)

type Options struct {
	Count      bool
	Repeated   bool
	Unique     bool
	IgnoreCase bool
	NumFields  int
	NumChars   int
}

func main() {
	options, inputFileName, outputFileName, err := parseArgs()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка: %v\n", err)
		flag.Usage()
		os.Exit(1)
	}

	lines := readInput(inputFileName)
	result := processUsingOptions(lines, options)
	writeOutput(result, outputFileName)
}

func parseArgs() (Options, string, string, error) {
	var options Options

	flag.BoolVar(&options.Count, "c", false, "подсчитать количество встречаний строки")
	flag.BoolVar(&options.Repeated, "d", false, "вывести только повторяющиеся строки")
	flag.BoolVar(&options.Unique, "u", false, "вывести только уникальные строки")
	flag.BoolVar(&options.IgnoreCase, "i", false, "не учитывать регистр")
	flag.IntVar(&options.NumFields, "f", 0, "не учитывать первые num_fields полей в строке")
	flag.IntVar(&options.NumChars, "s", 0, "не учитывать первые num_chars символов в строке")

	flag.Parse()

	if err := validateOptions(options); err != nil {
		return options, "", "", err
	}

	var inputFileName string
	var outputFileName string
	args := flag.Args()
	switch len(args) {
	case 0:
		inputFileName = ""
		outputFileName = ""
	case 1:
		inputFileName = args[0]
		outputFileName = ""
	case 2:
		inputFileName = args[0]
		outputFileName = args[1]
	}

	return options, inputFileName, outputFileName, nil
}

func validateOptions(options Options) error {
	modeCount := 0
	if options.Count {
		modeCount++
	}
	if options.Repeated {
		modeCount++
	}
	if options.Unique {
		modeCount++
	}

	if modeCount > 1 {
		return fmt.Errorf("параметры -c, -d, -u не могут использоваться вместе")
	}

	if options.NumFields < 0 {
		return fmt.Errorf("количество полей не может быть отрицательным")
	}

	if options.NumChars < 0 {
		return fmt.Errorf("количество символов не может быть отрицательным")
	}

	return nil
}
