package main

import "fmt"

func main() {

	// Make

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

	// Append

	numbers := []int{1, 2, 3, 4, 5, 6}
	var evens []int
	for _, n := range numbers {
		if n%2 == 0 {
			evens = append(evens, n)
		}
	}
	fmt.Println(evens)
	for i := range 37 {
		i++
		fmt.Print("-")
		if i == 37 {
			fmt.Println()
		}
	}

	matrix := [][]int{{1, 2}, {3, 4}, {5, 6}}
	var flat []int
	// var holder []int // I found out I do not need to do a 2 for loops to do that but not sure if that could make issues

	for _, row := range matrix {
		// row <-- matrix
		// holder = append(holder, row...)
		for _, val := range row { // val <-- row
			flat = append(flat, val)
		} // flat <-- val

	}
	// fmt.Println("Holder:", holder)
	fmt.Println("Matrix:", matrix)
	fmt.Println("Flat:", flat)
	for i := range 37 {
		i++
		fmt.Print("-")
		if i == 37 {
			fmt.Println()
		}
	}

	// Delete

}
