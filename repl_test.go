package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "   hello world    ",
			expected: []string{"hello", "world"},
		},
		{
			input:    " HeLlo  WOrld",
			expected: []string{"hello", "world"},
		},
		{
			input:    " what's your name ? obi wan",
			expected: []string{"what's", "your", "name", "?", "obi", "wan"},
		},
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    "                    ",
			expected: []string{},
		},
		{
			input:    "fhakjlafhailhfkjsdahflkasjhfkasdhjfjkadshfkasdjhf",
			expected: []string{"fhakjlafhailhfkjsdahflkasjhfkasdhjfjkadshfkasdjhf"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Return length doesn't match expected length.\nGot: %d\nExpected: %d", len(actual), len(c.expected))
		}
		for i := range actual {
			word := actual[i]
			expected_word := c.expected[i]
			if word != expected_word {
				t.Errorf("Mismatch at index %d: %v != %v", i, word, expected_word)
			}
		}
	}
}
