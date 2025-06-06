package main

import (
	"reflect"
	"testing"
)

var filenameRangeTests = map[string][]int{
	"src/Hiker_spec.re":              {4, 9},
	"test/hiker_test.exs":            {5, 10},
	"wibble/test/hiker_spec.rb":      {12, 17},
	"hiker_steps.rb":                 {0, 5},
	"hiker_spec.rb":                  {0, 5},
	"test_hiker.rb":                  {5, 10},
	"test_hiker.py":                  {5, 10},
	"test_hiker.sh":                  {5, 10},
	"tests_hiker.sh":                 {6, 11},
	"test_hiker.coffee":              {5, 10},
	"hiker_spec.coffee":              {0, 5},
	"hikerTest.chpl":                 {0, 5},
	"hiker.tests.c":                  {0, 5},
	"hiker_tests.c":                  {0, 5},
	"hiker_test.c":                   {0, 5},
	"hiker_Test.c":                   {0, 5},
	"HikerTests.cpp":                 {0, 5},
	"hikerTests.cpp":                 {0, 5},
	"HikerTest.cs":                   {0, 5},
	"HikerTest.java":                 {0, 5},
	"DiamondTest.kt":                 {0, 7},
	"HikerTest.php":                  {0, 5},
	"hikerTest.js":                   {0, 5},
	"hiker-test.js":                  {0, 5},
	"hiker-spec.js":                  {0, 5},
	"hiker.test.js":                  {0, 5},
	"hiker.tests.ts":                 {0, 5},
	"hiker_tests.erl":                {0, 5},
	"hiker_test.clj":                 {0, 5},
	"fizzBuzz_test.d":                {0, 8},
	"hiker_test.go":                  {0, 5},
	"hiker.tests.R":                  {0, 5},
	"hikertests.swift":               {0, 5},
	"HikerSpec.groovy":               {0, 5},
	"hikerSpec.feature":              {0, 5},
	"hiker.feature":                  {0, 5},
	"hiker.fun":                      {0, 5},
	"hiker.t":                        {0, 5},
	"hiker.plt":                      {0, 5},
	"prj/src/Hiker_bike_spec.re.bak": {8, 18},
	"src\\Hiker_spec.re":             {4, 9},
	"hiker":                          {0, 5},
	"":                               {},
}

func TestFilenameRange(t *testing.T) {
	for filename, expected := range filenameRangeTests {
		t.Run(filename, func(t *testing.T) {
			result := FilenameRange(filename)
			if !reflect.DeepEqual(result, expected) {
				t.Errorf("FilenameRange(%q) = %v; want %v",
					filename, result, expected)
			}
		})
	}
}
