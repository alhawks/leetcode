package main

import "fmt"

// find the two indices such that the numbers at those indices returns the target
func main() {
	fmt.Println(twoSum([]int{-3,4,3,90}, 0))
}


// numbers bigger than the target can be discarded. This didn't work because of negative numbers.
// I could keep track of the numbers I've found so far and what I would need to hit the target
func twoSum(nums []int, target int) []int {
	// needed -> index
	p := make(map[int]int)
	for i, num := range nums {
		if _, ok := p[num]; ok {
			return []int{i, p[num]}
		}
		p[target - num] = i
	}

	fmt.Println("NO SOLUTION FOUND")
	return []int{}
}