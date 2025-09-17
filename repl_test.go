package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "hello world",
			expected: []string{"hello", "world"},
		},
		{
			input:    "    hello world",
			expected: []string{"hello", "world"},
		},
		{
			input:    "goodbye       moon",
			expected: []string{"goodbye", "moon"},
		},
		{
			input:    "g o o d b y e",
			expected: []string{"g", "o", "o", "d", "b", "y", "e"},
		},
		{
			input:    "test case            ",
			expected: []string{"test", "case"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Incorrect output length")
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Incorrect word")
			}
		}
	}
}
