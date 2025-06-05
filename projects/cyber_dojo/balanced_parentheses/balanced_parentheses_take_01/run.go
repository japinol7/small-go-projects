package main

import (
	"fmt"
)

func main() {
	appName := "Balanced Parentheses"
	fmt.Println("Start app", appName)

	input_str := "((()))"
	res := AreParenthesesBalanced(input_str)
	fmt.Println("Are parentheses balanced for : '"+input_str+"' ? ", res)

	input_str = "(()))"
	res = AreParenthesesBalanced(input_str)
	fmt.Println("Are parentheses balanced for : '"+input_str+"' ? ", res)

	fmt.Println("End app", appName)
}
