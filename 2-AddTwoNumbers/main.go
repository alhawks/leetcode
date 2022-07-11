package main

func main() {
	l1 := &ListNode{Val: 9, Next: &ListNode{Val: 9}}
	l2 := &ListNode{Val: 9}
	addTwoNumbers(l1, l2)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// what do I do about numbers of different length?
	// How do I add in overflow (a single digit over 10)

	final := &ListNode{}
	c := final
	overflow := 0
	// start with numbers of same length
	for l1 != nil || l2 != nil || overflow != 0 {
		var val1, val2 int
		if l1 != nil {
			val1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			val2 = l2.Val
			l2 = l2.Next
		}
		sum := val1 + val2 + overflow
		if sum < 10 {
			c.Val = sum
			overflow = 0
		} else {
			c.Val = sum % 10
			overflow = sum / 10
		}
		if l1 != nil || l2 != nil || overflow != 0{
			c.Next = &ListNode{}
			c = c.Next
		}
	}

	return final
}



// My original attempt was all about converting everything between strings integers and ListNodes
// It didn't work, because eventually the ListNodes got bigger than the max allowed integer.

// func getNum(l *ListNode) int {
// 	b := strings.Builder{}
// 	current := l
// 	b.WriteString(fmt.Sprint(current.Val))
// 	current = current.Next
// 	for current != nil {
// 		b.WriteString(fmt.Sprint(current.Val))
// 		current = current.Next
// 	}

// 	fs := b.String()
// 	nb := strings.Builder{}
// 	for i := len(fs) - 1; i >= 0; i-- {
// 		nb.WriteString(string(fs[i]))
// 	}

// 	final, err := strconv.Atoi(nb.String())
// 	if err != nil {
// 		fmt.Printf("strconv failed %v\n", err)
// 	}
// 	return final
// }

// func numToList(n int) *ListNode {
// 	ln := &ListNode{}
// 	current := ln
// 	s := fmt.Sprint(n)
// 	for i := len(s) - 1; i >= 0; i-- {
// 		val, err := strconv.Atoi(string(s[i]))
// 		if err != nil {
// 			fmt.Printf("strconv failed %v\n", err)
// 		}
// 		current.Val = val
// 		if i > 0 {
// 			current.Next = &ListNode{}
// 			current = current.Next
// 		}
// 	}
// 	current.Next = nil
// 	return ln
// }
