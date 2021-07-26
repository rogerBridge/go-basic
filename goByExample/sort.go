package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"a", "b", "c"}
	sort.Strings(strs)
	fmt.Println(strs)

	ints := []int{1,3,2,4}
	sort.Ints(ints)
	fmt.Println(ints)

	fmt.Println(sort.IntsAreSorted(ints))
}
