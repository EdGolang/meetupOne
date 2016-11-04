package main

/*
If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
Find the sum of all the multiples of 3 or 5 below 1000.
*/

import (
	"encoding/json"	
	"net/http"
)

func EulerOneHandler(rw http.ResponseWriter, req *http.Request) {
    sum := findSum(1000)
	responseData := map[string]interface{} {
		"result" : "OK",
        "sum": sum,
	}

	jsonEncoder := json.NewEncoder(rw)
	jsonEncoder.Encode(responseData)
}

func findSum(under int) int {
    sum := 0
    for i := 1; i < under; i++ {
        if i % 3 == 0 || i % 5 == 0 {
            sum += i
        }
    }
    return sum
}
