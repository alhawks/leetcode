package main

import "fmt"

func main(){
	fmt.Println(lengthOfLongestSubstring("bbtablud"))

	fmt.Println(lengthOfLongestSubstringSlower("bbtablud"))
}

func lengthOfLongestSubstringSlower(s string) int {
	// character -> index
	seen := make(map[rune]int)
	var currentLen, longestSoFar int
    for i, char := range s {
		if firstSeenAt, ok := seen[char]; ok {
			// tally length
			if currentLen > longestSoFar {
				longestSoFar = currentLen
			}
			// reset
			seen = make(map[rune]int)
			currentLen = 0
			// re-create the map from the new beginning index
			// would subtracting be faster than re-adding?
			for j := firstSeenAt + 1; j < i; j++ {
				// TODO: indexing into a string seems to return bytes, not runes. I wonder if there's a way to
				// index by rune. It shouldn't matter for this problem though.
				seen[rune(s[j])] = j
				currentLen++
			}
		}
		seen[char] = i
		currentLen++
	}
	if currentLen > longestSoFar {
		longestSoFar = currentLen
	}
	return longestSoFar
}

// This version was much much faster on the leetcode problem set. I wonder why.
// theory: The worst case scenario for this version adds and removes every letter once.
// 		   The worst case on the other one ends up re-adding letters a bunch of times.
func lengthOfLongestSubstring(s string) int {
	// character -> index
	seen := make(map[rune]int)
	var currentLen, longestSoFar, slow int
    for i, char := range s {
		if firstSeenAt, ok := seen[char]; ok {
			// tally length
			if currentLen > longestSoFar {
				longestSoFar = currentLen
			}
			// remove everything up to and including the first time we saw the repeat letter
			for j := slow; j <= firstSeenAt; j++ {
				delete(seen, rune(s[j]))
				currentLen--
			}
			slow = firstSeenAt + 1
		}
		seen[char] = i
		currentLen++
	}
	if currentLen > longestSoFar {
		longestSoFar = currentLen
	}
	return longestSoFar
}