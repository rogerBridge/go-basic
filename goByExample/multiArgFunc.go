package main

import "fmt"

func MultiArgFunc(nums ...int) (int, error) {
	sum := 0
	for _, i := range nums {
		sum += i
	}
	fmt.Println(sum)
	return sum, nil
}
// a := []int{1, 2, 3, 4}
// MultiArgFunc(nums...) equals MultiArgFunc(1,2,3,4)