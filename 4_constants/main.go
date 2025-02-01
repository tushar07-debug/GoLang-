package main

import "fmt"

const age = 20

// name:="go"
var name string = "tushar"

func main() {
	const name = "golang"
	// name="js"
	// const age = 30

	const(
		port = 5000
		host = "localhost"
	)
	// port = 5500
	fmt.Println(port , host)
}