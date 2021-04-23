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
	- select statement
	- context package
	- signals
- graceful shutdown
	- basic example
	- different signals
	- http server
	
Have fun.`

	fmt.Print(structure)
}
