package main

import (
	"fmt"
	"testing"
)

func TestPrintDiamondEmpty(t *testing.T) {
	testCases := []struct {
		letter   string
		expected string
	}{
		{"", ""},
		{" ", ""},
		{"ñ", ""},
		{"FG", ""},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Diamond(%s)", tc.letter), func(t *testing.T) {
			result := NewDiamond(tc.letter).String()
			if result != tc.expected {
				t.Errorf("Expected: %q, Got: %q", tc.expected, result)
			}
		})
	}
}

func TestPrintDiamondA(t *testing.T) {
	result := NewDiamond("A").String()
	expected := "A"
	if result != expected {
		t.Errorf("Expected: %q, Got: %q", expected, result)
	}
}

func TestPrintDiamondB(t *testing.T) {
	result := NewDiamond("B").String()
	expected := " A \n" +
		"B B\n" +
		" A "
	if result != expected {
		t.Errorf("Expected: %q, Got: %q", expected, result)
	}
}

func TestPrintDiamondC(t *testing.T) {
	result := NewDiamond("C").String()
	expected := "  A  \n" +
		" B B \n" +
		"C   C\n" +
		" B B \n" +
		"  A  "
	if result != expected {
		t.Errorf("Expected: %q, Got: %q", expected, result)
	}
}

func TestPrintDiamondD(t *testing.T) {
	result := NewDiamond("D").String()
	expected := "   A   \n" +
		"  B B  \n" +
		" C   C \n" +
		"D     D\n" +
		" C   C \n" +
		"  B B  \n" +
		"   A   "
	if result != expected {
		t.Errorf("Expected: %q, Got: %q", expected, result)
	}
}

func TestPrintDiamondF(t *testing.T) {
	result := NewDiamond("F").String()
	expected := "     A     \n" +
		"    B B    \n" +
		"   C   C   \n" +
		"  D     D  \n" +
		" E       E \n" +
		"F         F\n" +
		" E       E \n" +
		"  D     D  \n" +
		"   C   C   \n" +
		"    B B    \n" +
		"     A     "
	if result != expected {
		t.Errorf("Expected: %q, Got: %q", expected, result)
	}
}

func TestPrintDiamondEAndOutput(t *testing.T) {
	diamond := NewDiamond("E")
	result := diamond.String()
	expected := "    A    \n" +
		"   B B   \n" +
		"  C   C  \n" +
		" D     D \n" +
		"E       E\n" +
		" D     D \n" +
		"  C   C  \n" +
		"   B B   \n" +
		"    A    "
	if result != expected {
		t.Errorf("Expected: %q, Got: %q", expected, result)
	}

	// Print output for visual inspection
	fmt.Printf("\nOutput:\n%s\n", result)
}
