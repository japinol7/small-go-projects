package main

import "testing"

func TestAreParenthesesBalanced(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"empty string", "", true},
		{"simple parentheses", "()", true},
		{"simple braces", "{}", true},
		{"nested braces and parentheses", "{()}", true},
		{"complex nested", "{[()]}", true},
		{"alternative nesting", "[({})]", true},
		{"multiple pairs", "{}([])", true},
		{"complex valid", "{()}[[{}]]", true},
		{"unbalanced close", "[]]", false},
		{"incorrectly nested", "{{)(}}", false},
		{"mismatched", "({)}", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AreParenthesesBalanced(tt.input); got != tt.expected {
				t.Errorf("AreParenthesesBalanced(%q) = %v, want %v",
					tt.input, got, tt.expected)
			}
		})
	}
}
