package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		// Test cases follow this format
		input string
		expected []string
	} {
		// Test cases go here
		{
			input: "",
			expected: []string{},
		},
		{
			input: "   hello   world   ",
			expected: []string{"hello", "world"},
		},
		{
			input: "Scorbunny, Xatu, Dragapult",
			expected: []string{"scorbunny,", "xatu,", "dragapult"},
		},
		{
			input: "AAAA BBBB cccc",
			expected: []string{"aaaa", "bbbb", "cccc"},
		},
	}

	// Loop over cases and test
	for _, c := range cases {
		actual := cleanInput(c.input)

		// Test slice lengths
		if len(actual) != len(c.expected) {
			t.Errorf("Error: Slice length mismatch (%v)", c.input)
		}

		// Test formatting
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("Error: Formatting mismatch (%v)", c.input)
			}
		}
	}
}