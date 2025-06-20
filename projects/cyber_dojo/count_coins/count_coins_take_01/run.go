package main

import (
	"fmt"
)

func main() {
	appName := "Count coins"
	fmt.Println("Start app", appName)

	cc := &CountCoins{}
	result := cc.Changes(100)

	fmt.Println("How many ways are there to make change for a dollar")
	fmt.Println("using these common coins? (1 dollar = 100 cents)")
	fmt.Printf("Result: %d\n", result)

	fmt.Println("End app", appName)
}
