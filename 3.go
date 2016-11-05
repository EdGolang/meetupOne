package main

/*
The prime factors of 13195 are 5, 7, 13 and 29.
What is the largest prime factor of the number 600851475143 ?
*/

import (
	"encoding/json"
	"fmt"
	_ "math"
	"net/http"
)

func EulerThreeHandler(rw http.ResponseWriter, req *http.Request) {
	prime := largestPrime(600851475143)
	responseData := map[string]interface{}{
		"result": "OK",
		"prime":  prime,
	}

	jsonEncoder := json.NewEncoder(rw)
	jsonEncoder.Encode(responseData)
}

func largestPrime(of uint64) uint64 {
	guesses := []uint64{}
	for i := uint64(1); i < of; i++ {
		if of%i == 0 {
			if isPrime(i) {
				fmt.Printf("guess: %v\n", i)
				guesses = append(guesses, i)
			}
		}

	}
	return guesses[len(guesses)-1]
}

func isPrime(p uint64) bool {
	if p == uint64(2) {
		return true
	}
	if p == uint64(3) {
		return true
	}
	if p%uint64(2) == 0 {
		return false
	}
	if p%uint64(3) == 0 {
		return false
	}

	i := uint64(5)
	w := uint64(2)

	for i*i <= p {
		if p%i == 0 {
			return false
		}

		i += w
		w = 6 - w
	}

	return true

}
