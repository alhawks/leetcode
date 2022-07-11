package main

import "fmt"

var (
	letterToNumber = map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	rules = map[string]bool {
		"IV": true,
		"IX": true,
		"XL": true,
		"XC": true,
		"CD": true,
		"CM": true,
	}
)

func main() {
	fmt.Println(romanToInteger("MCMXCIV"))
}

// Six cases
// I before V and X
// X before L and C
// C before D and M

func romanToInteger(input string) int {
	var sum int
	var previousRune rune
	// store the previous value to avoid an extra map access later. It made my leetcode time better.
	var previousVal int
	for _, char := range input {
		val := letterToNumber[char]
		var finalVal int
		if rules[string(previousRune) + string(char)] {
			finalVal = letterToNumber[char] - 2 * previousVal
		} else {
			finalVal = val
		}

		sum += finalVal
		previousRune = char
		previousVal = val
	}
	return sum
}
