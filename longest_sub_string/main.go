package main

import (
	"fmt"
)

func main() {
	fmt.Sprintf("Longest sub string")
	requestedString := "abcdeafbdgcbb"
	lgStr := longestSubString(requestedString)
	fmt.Println("Longest sub string", lgStr)
}

func longestSubString(reqStr string) int {
	hashTable := make(map[string]int)
	ans := 0
	i := 0
	maxM := 0
	start := 0
	end := 0
	for j:=0; j<len(reqStr); j++ {
		_, ok := hashTable[string(reqStr[i])]
		if ok {
			i = max(hashTable[string(reqStr[j])], i)
		}
		ans = max(ans, j - i + 1)
		if ans > maxM {
			maxM = ans
			start = i
			end = j + 1
		}
		hashTable[string(reqStr[j])] = j + 1
		fmt.Println(">>>>>", ans, hashTable, start, end)
	}
	return ans
}
// "a   b   c   d   e   a   f   b   d   g    c     b     b"
//i 0   0   0   0   0   1   1   2   4   4    4     8     12
//j 0   1   2   3   4   5   6   7   8   9    10    11    12
//a 1   2   3   4   5   5   6   6   6   6     7    7     7
//h a=1 b=2 c=3 d=4 e=5 a=6 f=7 b=8 d=9 g=10 c=11  b=12  b=13
func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
