package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	type test struct {
		nums1 []int
		nums2 []int
		expected float64
	}
	tests := []test {
		{
			[]int{5,6,7},
			[]int{1,2,3,4},
			4,
		},
		{
			[]int{1,2,3,4},
			[]int{5,6,7},
			4,
		},
		{
			[]int{4,4,4},
			[]int{4, 4, 4, 4},
			4,
		},
		{
			[]int{4,5,6, 7},
			[]int{1, 1, 2, 3},
			3.5,
		},
		{
			[]int{1, 1, 2, 3},
			[]int{4,5,6, 7},
			3.5,
		},
		{
			[]int{4,5,6, 7},
			[]int{0, 1, 2, 3},
			3.5,
		},
		{
			[]int{},
			[]int{0, 1, 2},
			1,
		},
		{
			[]int{1, 3},
			[]int{2},
			2,
		},
		{
			[]int{1, 3},
			[]int{2, 7},
			2.5,
		},
		{
			[]int{2, 7},
			[]int{1, 3},
			2.5,
		},
		{
			[]int{},
			[]int{2, 3},
			2.5,
		},
		{
			[]int{},
			[]int{1, 2, 3, 4},
			2.5,
		},
		{
			[]int{1, 2},
			[]int{3, 4, 5, 6, 7},
			4,
		},
	}
	for _, t := range tests {
		ans := findMedianSortedArrays(t.nums1, t.nums2)
		fmt.Printf("GOT: %f, WANT: %f\n", ans, t.expected)
	
	}
}

// with an odd number of entries, the median is the middle value.
// with an even number of entries, the median is halfway between the two middle values (if they're not the same) I think...

// the problem: find the combined median without combining the arrays
// idea: Do some binary sort kind of thing to figure out how the two arrays could fit together.
// the numbers being inserted

// find the m + n / 2 value...j
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	answer, err := checkArray(nums1, nums2)
	if err != nil {
		answer, _ = checkArray(nums2, nums1)
	}
	return answer
}

func checkArray(nums1 []int, nums2 []int) (float64, error) {
// binary search each array
totalLen := len(nums1) + len(nums2)
medianPos := totalLen / 2

hi := len(nums1) - 1
lo := 0

for hi >= lo {
	i := (hi + lo) / 2
	if i > medianPos {
		// Too big
		hi = i - 1
		continue
	}
	// Where the element needs to fit in the next array
	neededPos := medianPos - i
	if neededPos < len(nums2) {
		if nums2[neededPos] < nums1[i] {
			// too big
			hi = i - 1
			continue
		}
	}
	if neededPos > len(nums2) {
		// too small
		lo = i + 1
		continue
	}
	if neededPos != 0 && len(nums2) != 0 && nums2[neededPos - 1] > nums1[i] {
		// too small
		lo = i + 1
		continue
	}
	// found it
	// If the total length is even, then some averaging needs to be done
	if totalLen % 2 == 0 {
		// stupid, frustrating edge case
		if len(nums2) == 0 {
			low := (len(nums1) / 2) -1
			high := (len(nums1) / 2)
			return (float64(nums1[low]) + float64(nums1[high])) / 2, nil
		}
		var s1, s2 int
		// If it's even, we have to find the next lowest element.
		// It will be one index lower in the current array
		// OR on either side of neededPos
		var f1, f2 bool
	
		if i!=0 {
			f1 = true
			s1 = nums1[i-1]
		}
		if neededPos != 0 && len(nums2) != 0 {
			f2 = true
			s2 = nums2[neededPos - 1]
		}

		if f1 && f2{
			return (math.Max(float64(s1), float64(s2)) + float64(nums1[i])) / 2, nil
		} else if f1 {
			return (float64(s1) + float64(nums1[i])) / 2, nil
		} else {
			return (float64(s2) + float64(nums1[i])) / 2, nil
		}
	}
	return float64(nums1[i]), nil
}
// based on the value I'm looking at, figure out where it would have to be in order to be the median
// If it fits, then I'm done.
// Otherwise, keep going

return 0, errors.New("didn't find it")
}