package main

import (
	"reflect"
	"testing"
)

func TestProcessUsingOptions_Basic(t *testing.T) {
	tests := []struct {
		name     string
		lines    []string
		options  Options
		expected []string
	}{
		{
			name:     "unique",
			lines:    []string{"a", "b", "a", "c"},
			options:  Options{},
			expected: []string{"a", "b", "c"},
		},
		{
			name:     "count",
			lines:    []string{"a", "b", "a", "c"},
			options:  Options{Count: true},
			expected: []string{"2 a", "1 b", "1 c"},
		},
		{
			name:     "repeated",
			lines:    []string{"a", "b", "a", "c"},
			options:  Options{Repeated: true},
			expected: []string{"a"},
		},
		{
			name:     "not repeated",
			lines:    []string{"a", "b", "a", "c"},
			options:  Options{Unique: true},
			expected: []string{"b", "c"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := processUsingOptions(tt.lines, tt.options)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("%s: expected %v, got %v", tt.name, tt.expected, result)
			}
		})
	}
}

func TestProcessUsingOptions_WithFilters(t *testing.T) {
	tests := []struct {
		name     string
		lines    []string
		options  Options
		expected []string
	}{
		{
			name:     "ignore case",
			lines:    []string{"Apple", "apple", "Banana"},
			options:  Options{IgnoreCase: true},
			expected: []string{"Apple", "Banana"},
		},
		{
			name:     "ignore fields",
			lines:    []string{"1 apple", "2 apple", "3 banana"},
			options:  Options{NumFields: 1},
			expected: []string{"1 apple", "3 banana"},
		},
		{
			name:     "ignore chars",
			lines:    []string{"abc", "abd", "abc"},
			options:  Options{NumChars: 2},
			expected: []string{"abc", "abd"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := processUsingOptions(tt.lines, tt.options)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("%s: expected %v, got %v", tt.name, tt.expected, result)
			}
		})
	}
}
