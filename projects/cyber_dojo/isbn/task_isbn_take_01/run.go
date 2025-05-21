package main

import (
	"fmt"
)

func main() {
	appName := "Validate ISBN-13 and ISBN-10"
	fmt.Println("Start app", appName)

	// Check valid isbn-13
	isbn := "978-0-262-13472-9"
	res := ValidateISBN(isbn)
	fmt.Println("Should be true.  Is '"+isbn+"' a valid isbn?: ", res)

	// Check invalid isbn-13
	isbn = "978-0-262-13472-1"
	res = ValidateISBN(isbn)
	fmt.Println("Should be false. Is '"+isbn+"' a valid isbn?: ", res)

	// Check valid isbn-10
	isbn = "0-8044-2957-X"
	res = ValidateISBN(isbn)
	fmt.Println("Should be true.  Is '"+isbn+"' a valid isbn?: ", res)

	// Check valid isbn-10
	isbn = "0-4A0-84525-2"
	res = ValidateISBN(isbn)
	fmt.Println("Should be false. Is '"+isbn+"' a valid isbn?: ", res)

	fmt.Println("End app", appName)
}
