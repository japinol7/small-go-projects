package main

import (
	"strings"
	"testing"
)

// Original text with $ as a column separator
const inputTextOrig = `Given$a$text$file$of$many$lines,$where$fields$within$a$line$
are$delineated$by$a$single$'dollar'$character,$write$a$program
that$aligns$each$column$of$fields$by$ensuring$that$words$in$each$
column$are$separated$by$at$least$one$space.`

// inputText with $ replaced by ColSep
var inputText string

func init() {
	// Initialize inputText by replacing $ with ColSep
	inputText = strings.ReplaceAll(inputTextOrig, "$", ColSep)
}

func TestAlignColumnsLeft(t *testing.T) {
	// Original string with spaces and $ as separators
	ExpectedOrig := "Given  a          text      file   of     many     lines,     where    fields within  a  line\n" +
		"are    delineated by        a      single 'dollar' character, write    a      program\n" +
		"that   aligns     each      column of     fields   by         ensuring that   words   in each\n" +
		"column are        separated by     at     least    one        space.  "

	// Replace spaces with FillCh
	expected := strings.ReplaceAll(ExpectedOrig, " ", FillCh)

	result := AlignColumns(inputText, LeftAlignment)
	if result != expected {
		t.Errorf("Expected:\n%s\nGot:\n%s", expected, result)
	}
}

func TestAlignColumnsRight(t *testing.T) {
	// Original string with spaces and $ as separators
	ExpectedOrig := " Given          a      text   file     of     many     lines,    where fields  within  a line\n" +
		"   are delineated        by      a single 'dollar' character,    write      a program\n" +
		"  that     aligns      each column     of   fields         by ensuring   that   words in each\n" +
		"column        are separated     by     at    least        one   space."

	// Replace spaces with FillCh
	expected := strings.ReplaceAll(ExpectedOrig, " ", FillCh)

	result := AlignColumns(inputText, RightAlignment)
	if result != expected {
		t.Errorf("Expected:\n%s\nGot:\n%s", expected, result)
	}
}

func TestAlignColumnsCenter(t *testing.T) {
	// Original string with spaces and $ as separators
	ExpectedOrig := "Given      a        text     file    of     many     lines,    where   fields within  a  line\n" +
		" are   delineated    by       a    single 'dollar' character,  write     a    program\n" +
		" that    aligns     each    column   of    fields      by     ensuring  that   words  in each\n" +
		"column    are     separated   by     at    least      one      space. "

	// Replace spaces with FillCh
	expected := strings.ReplaceAll(ExpectedOrig, " ", FillCh)

	result := AlignColumns(inputText, CenterAlignment)
	if result != expected {
		t.Errorf("Expected:\n%s\nGot:\n%s", expected, result)
	}
}

func TestAlignColumnsEmpty(t *testing.T) {
	result := AlignColumns("", LeftAlignment)
	expected := ""
	if result != expected {
		t.Errorf("Expected empty string, got: %s", result)
	}
}
