package main

import (
	"fmt"
)

func maxArea(req []int) int {
	maxArea := 0
	left := 0
	right := len(req)-1
	for left < right {
		width := right - left
		maxArea = max(maxArea, min(req[left], req[right])*width)
		if req[left] <= req[right] {
			left = left + 1
		} else {
			right = right - 1
		}
	}
	return maxArea
}

func max(a, b int) int{
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	height := []int{1,8,6,2,5,4,8,3,7}
	res := maxArea(height)
	fmt.Println("Max area :- ", res)
}
