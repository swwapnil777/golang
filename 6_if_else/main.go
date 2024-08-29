package main

import "fmt"

func main() {

	// check odd and even

	n := 11

	var isPrime bool = true

	if n == 1 {
		fmt.Println("Not prime")
		return
	}

	for i := 2; i < n/2; i++ {

		if n%i == 0 {

			isPrime = false
			break

		}
	}

	if isPrime {
		fmt.Println("Prime")
	} else {
		fmt.Println("Not prime")
	}

}
