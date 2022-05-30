package main

//A = [3, 4, 2]
//output = [[1, 2], [1, 3], [2, 2]]
//The above-average subarrays are [3, 4], [3, 4, 2], and [4].

import "fmt"

func aboveAverage(input []int) [][]int {
	avg := 4
	sum := 0
	k := 1
	var result [][]int
	for i:=1; i<=len(input); i++ {
		sum = sum + input[i-1]
		if avg <= sum {
			avgSubArr := []int{k, i}
			result = append(result, avgSubArr)
		}
		if i == len(input) {
			k++
			i=k - 1
			sum = 0
		}
	}
	return result
}

func main() {
	input := []int{3, 4, 2}
	res := aboveAverage(input)
	fmt.Println("Result :", res)
}