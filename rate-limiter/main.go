package main

import "fmt"

func main() {
	reqArr := []int{1,1,1,1,2,2,2,3,3,3,4,4,4,5,5,5,6,6,6,7,7,7,7,11,11,11,11}
	dropCount := rateLimiterDropCount(reqArr)
	fmt.Println("Drop Count", dropCount)
}

func rateLimiterDropCount(reqArr []int) int {
	dropCount := 0
	dropLimitIncThree := 0
	dropLimitIncTen := 0
	for i := 0; i < len(reqArr)-1; i++ {
		dropLimitIncTen++
		if reqArr[i] == reqArr[i + 1] {
			dropLimitIncThree++
			continue
		}
		if dropLimitIncThree > 2 ||  dropLimitIncTen > 20 {
			fmt.Println("Drop : ", reqArr[i], i)
			dropCount++
		}
	}
	return dropCount
}