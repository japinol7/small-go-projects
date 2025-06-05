package main

import (
	"fmt"
)

func main() {
	appName := "LCD Digits"
	fmt.Println("Start app", appName)

	inputNumber := 1234567890
	lcd := LcdDigits{}
	result, err := lcd.Generate(inputNumber)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(
			"LCD Digits for " + fmt.Sprint(inputNumber) + "\n" + result,
		)
	}

	fmt.Println("End app", appName)
}
