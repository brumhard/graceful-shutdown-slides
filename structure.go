// START OMIT
package main

import "fmt"

func main() {
	printStructure() // HL
}

// END OMIT

func printStructure() {
	structure := `structure:
- basics
	- context package
	- signals
	- select statement
- demo
	- basic run function
	- container shutdown
	- http server`

	fmt.Print(structure)
}
