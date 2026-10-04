package main

import "fmt"

func main() {
	primes := []int{2, 3, 5, 7, 11, 13}
	for i, prime := range primes {
		fmt.Printf("Prime %d: %d\n", i, prime)
	}
}
