package main

/*
The prime factors of 13195 are 5, 7, 13 and 29.
What is the largest prime factor of the number 600851475143 ?
*/

import (
	"net/http"
    "math"
    "encoding/json"
)

func EulerThreeHandler(rw http.ResponseWriter, req *http.Request) {
    prime := largestPrime(600851475143)
	responseData := map[string]interface{} {
		"result" : "OK",
        "prime": prime,
	}

	jsonEncoder := json.NewEncoder(rw)
	jsonEncoder.Encode(responseData)
}

func largestPrime(of uint64) uint64 {
    for i := of; i > 0; i-- {
        if of % i == 0 && isPrime(i) {
            return i
        }
    }
    return 0
}

func isPrime(p uint64) bool {
    for i := uint64(1); i < uint64(math.Sqrt(float64(p))); i++ {
        if p % i == 0 {
            return false
        }
    }
    return true
}
