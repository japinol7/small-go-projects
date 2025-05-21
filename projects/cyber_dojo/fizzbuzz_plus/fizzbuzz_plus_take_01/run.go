package main

import (
	"fmt"
)

func main() {
	appName := "Quicksort for regular integers"
	fmt.Println("Start app", appName)

	res := FizzbuzzRange(15)
	fmt.Println(String(res))

	fmt.Println("\nEnd app", appName)
}
