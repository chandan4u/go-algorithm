package main

import "fmt"

// Create one empty stack
// Create one defined hashmap for braces
// Push item inside stack
// Once we find closer then check previous stack value, If it match pop item from stack.


func isValid(req string) bool {
	if len(req)%2 != 0 {
		return false
	}
	var stack []string
	hashMap := map[string]string{
		"(" : ")",
		"{" : "}",
		"[" : "]",
	}
	for i:=0; i<len(req); i++ {
		fmt.Println(">>>>> Stack", stack)
		if closer, ok := hashMap[string(req[i])]; ok {
			fmt.Println(">>>> cl", closer)
			stack = append(stack, closer)
			continue
		}
		lg := len(stack) - 1
		if lg < 0 || string(req[i]) != stack[lg] {
			return false
		}
		stack = stack[:lg]
		fmt.Println(">>>>> Stack", stack)
	}
	return len(stack) == 0
}

func main()  {
	str := "[[[[][]]]]"
	res := isValid(str)
	fmt.Println("Result ::", res)

	var st []int
	for i:=1; i<10; i++ {
		st = append(st, i)
		if i == 9 {
			fmt.Println(st)
			st = st[3:7]
			fmt.Println(st)
		}
	}
	fmt.Println(st)

}