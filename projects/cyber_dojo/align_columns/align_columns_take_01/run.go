package main

import (
	"fmt"
	"strings"
)

func main() {
	appName := "Align Columns"
	fmt.Println("Start app", appName)

	inputTextOrig := "Given$a$text$file$of$many$lines,$where$fields$within$a$line$\n" +
		"are$delineated$by$a$single$'dollar'$character,$write$a$program\n" +
		"that$aligns$each$column$of$fields$by$ensuring$that$words$in$each$\n" +
		"column$are$separated$by$at$least$one$space"

	// Replace $ with ColSep constant
	inputText := strings.ReplaceAll(inputTextOrig, "$", ColSep)

	fmt.Println("\nLeft-aligned text:")
	fmt.Println(AlignColumns(inputText, LeftAlignment))

	fmt.Println("\nCenter-aligned text:")
	fmt.Println(AlignColumns(inputText, CenterAlignment))

	fmt.Println("\nRight-aligned text:")
	fmt.Println(AlignColumns(inputText, RightAlignment))

	fmt.Println("\nEnd app", appName)
}
