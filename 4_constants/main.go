package main

import "fmt"

const firstName = "Swapnil"

func main() {

	const name string = "Swapnil"

	fmt.Println(name)

	const (
		port = 5000
		host = "localhost"
	)

	fmt.Println(port, host)
}
