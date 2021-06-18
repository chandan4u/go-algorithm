package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 3}
	sum := 6
	// desiApproach(arr, sum)
	hashApproach(arr, sum)
}

func desiApproach(arr []int, sum int) {
	for i:=0; i<len(arr); i++ {
		for j:=i+1; j<len(arr); j++ {
			if arr[i] + arr[j] == sum {
				fmt.Printf("Pair with the given sum %d - %d [ %d, %d]\n",arr[i], arr[j], i, j)
			}
		}
	}
}

func hashApproach(arr []int, sum int) {
	hashTable := make(map[int]int)
	for i:=0; i<len(arr); i++ {
		temp := sum - arr[i]
		_, ok := hashTable[temp]
		if ok {
			fmt.Printf("Pair with the given sum %d - %d [%d, %d]\n",temp, arr[i], hashTable[temp], i)
		}
		hashTable[arr[i]] = i
	}
}


