package main

import "fmt"

func main() {

	nums := make([]int, 5)
	for i := range nums {
		nums[i] = i * 10
	}
	fmt.Println(nums)

	for i := range 37 {
		i++
		fmt.Print("-")
		if i == 37 {
			fmt.Println()
		}
	}

	keys := []string{"a", "b", "c"}
	m := make(map[string]int, len(keys))
	for i, k := range keys {
		m[k] = i + 1
	}
	fmt.Println(m)

	for i := range 37 {
		i++
		fmt.Print("-")
		if i == 37 {
			fmt.Println()
		}
	}

}
