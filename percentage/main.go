package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

/* Problem Name is &&& Best Average Grade &&& PLEASE DO NOT REMOVE THIS LINE. */

/*
**  Instructions:
**
**  Given a list of student test scores, find the best average grade.
**  Each student may have more than one test score in the list.
**
**  Complete the bestAverageGrade function in the editor below.
**  It has one parameter, scores, which is an array of student test scores.
**  Each element in the array is a two-element array of the form [student name, test score]
**  e.g. [ "Bobby", "87" ].
**  Test scores may be positive or negative integers.
**
**  If you end up with an average grade that is not an integer, you should
**  use a floor function to return the largest integer less than or equal to the average.
**  Return 0 for an empty input.
**
**  Example:
**
**  Input:
**  [ [ "Bobby", "87" ],
**    [ "Charles", "100" ],
**    [ "Eric", "64" ],
**    [ "Charles", "22" ] ].
**
**  Expected output: 87
**  Explanation: The average scores are 87, 61, and 64 for Bobby, Charles, and Eric,
**  respectively. 87 is the highest.
 */

func BestAverageGrade(r io.Reader) (int, error) {
	scanner := bufio.NewScanner(r)
	hashTable := make(map[string]string)
	res := make(map[string]int)
	max := 0
	for scanner.Scan() {
		t := scanner.Text()
		splitedInput := strings.Split(t, ",")
		if len(splitedInput) != 2 {
			continue
		}
		score, _ := strconv.Atoi(strings.TrimSpace(splitedInput[1]))
		_, ok := hashTable[splitedInput[0]]
		if ok {
			splitedHashValue := strings.Split(hashTable[splitedInput[0]], "_")
			count, _ := strconv.Atoi(strings.TrimSpace(splitedHashValue[1]))
			nextScore, _ := strconv.Atoi(strings.TrimSpace(splitedHashValue[0]))
			score = score + nextScore
			countNext := count + 100
			hashTable[splitedInput[0]] = fmt.Sprintf("%d_%d", score, countNext)
		} else {
			hashTable[splitedInput[0]] = fmt.Sprintf("%d_%d", score, 100)
		}
		findPercentage(hashTable, res, splitedInput[0])
	}
	for _, value := range res {
		if value > max {
			max = value
		}
	}
	return max, nil
}

func findPercentage(hash map[string]string, res map[string]int, key string) {
	splitedHash := strings.Split(hash[key], "_")
	score, _ := strconv.Atoi(strings.TrimSpace(splitedHash[0]))
	count, _ := strconv.Atoi(strings.TrimSpace(splitedHash[1]))
	prctg := score * 100/count
	res[key] = prctg
	fmt.Println(res)
}

func TestsPass() {
	// TODO: implement more test cases
	tests := []struct {
		in   io.Reader
		want int
	}{
		{
			in:   strings.NewReader(tc1),
			want: 87,
		},
	}
	ok := true
	for _, test := range tests {
		got, err := BestAverageGrade(test.in)
		if err != nil || got != test.want {
			ok = false
		}
	}
	if ok {
		fmt.Println("All tests pass")
		return
	}
	fmt.Println("There are test failures")
}

func main() {
	TestsPass()
}

// example
var tc1 = `
Boby, 87
Charles,   100
Eric, 64
Charles, 22
`
