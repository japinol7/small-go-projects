package main

import (
	"fmt"
)

func main() {
	appName := "Filename Range"
	fmt.Println("Start app", appName)

	filename := "wibble/test/hiker_spec.rb"
	nameRange := FilenameRange(filename)
	wantedName := filename[nameRange[0]:nameRange[1]]
	fmt.Printf("%s -> %v -> %v\n", filename, nameRange, wantedName)

	fmt.Println("End app", appName)
}
