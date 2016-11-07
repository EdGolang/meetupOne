package main

/*
A palindromic number reads the same both ways.
The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.
Find the largest palindrome made from the product of two 3-digit numbers.
*/

import (
	"encoding/json"
	"math"
	"net/http"
	"strconv"
)

func EulerFourHandler(rw http.ResponseWriter, req *http.Request) {
	responseData := map[string]interface{}{
		"result": "OK",
		"answer": getLargestPalindromeByDigits(3),
	}

	jsonEncoder := json.NewEncoder(rw)
	jsonEncoder.Encode(responseData)
}

func getLargestPalindromeByDigits(digits int) int {
	productA := getLargestNumberForDigits(digits)
	productB := productA
	limit := getLargestNumberForDigits(digits - 1)
	palindromes := []int{}

	for i := productA; i > limit; i-- {
		for n := productB; n > limit; n-- {
			palindromes = palindromeCollector(n, i, palindromes)
		}
		palindromes = palindromeCollector(i, productB, palindromes)
	}
	return findGreatestNumberInSlice(palindromes)
}

func getLargestNumberForDigits(digits int) int {
	products := math.Pow(10, (float64(digits))) - 1
	return int(products)
}

func palindromeCollector(a int, b int, list []int) []int {
	if isPalindrome(a * b) {
		list = append(list, a*b)
	}
	return list
}

func isPalindrome(num int) bool {
	numSlice := convertIntToSliceOfString(num)
	for i := 0; i < len(numSlice); i++ {
		if numSlice[i] != numSlice[len(numSlice)-(i+1)] {
			return false
		}
	}
	return true
}

func convertIntToSliceOfString(num int) []string {
	numString := strconv.Itoa(num)
	var numList []string

	for i := 0; i < len(numString); i++ {
		numList = append(numList, string(numString[i]))
	}
	return numList
}

func findGreatestNumberInSlice(slice []int) int {
	highest := 0
	for i := 0; i < len(slice); i++ {
		if slice[i] > highest {
			highest = slice[i]
		}
	}
	return highest
}
