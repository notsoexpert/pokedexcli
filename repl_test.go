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
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    " ApPlE ChErRy PEACH banana                ",
			expected: []string{"apple", "cherry", "peach", "banana"},
		},
		{
			input:    "                                    TEST l",
			expected: []string{"test", "l"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Output slice length != expected slice length")
			return
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("Output word doesn't match expected word")
				return
			}
		}
	}
}
