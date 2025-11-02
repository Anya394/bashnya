package main

import (
	"testing"
)

func TestConvertThreeDigit(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected string
	}{
		{"Zero", 0, ""},
		{"One", 1, "один"},
		{"Nineteen", 19, "девятнадцать"},
		{"Forty", 40, "сорок"},
		{"NinetyNine", 99, "девяносто девять"},
		{"OneHundredTwentyFive", 125, "сто двадцать пять"},
		{"ThreeHundred", 300, "триста"},
		{"FiveHundredSixtySeven", 567, "пятьсот шестьдесят семь"},
		{"NineHundredNinetyNine", 999, "девятьсот девяносто девять"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := convertThreeDigit(tt.input)
			if result != tt.expected {
				t.Errorf("convertThreeDigit(%d) = %q, expected %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestNumberToString(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected string
	}{
		{"FiveThousand", 5000, "пять тысяч"},
		{"SevenThousandFiveHundredSixtySeven", 7567, "семь тысяч пятьсот шестьдесят семь"},
		{"TwelveThousandThreeHundred", 12300, "двенадцать тысяч триста"},
		{"NineHundredNinetyNineThousand", 999000, "девятьсот девяносто девять тысяч"},
		{"NineHundredNinetyNineThousandNineHundredNinetyNine", 999999, "девятьсот девяносто девять тысяч девятьсот девяносто девять"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := numberToString(tt.input)
			if result != tt.expected {
				t.Errorf("numberToString(%d) = %q, expected %q", tt.input, result, tt.expected)
			}
		})
	}
}
