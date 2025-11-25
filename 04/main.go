package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Median of slice 10, 9, 8, 7, 6, 5, 4, 3, 2 is? ")

	var median int
	nums := []int{10, 3, 6, 7, 8, 5, 4, 9, 2}

	sort.Ints(nums)

	i := len(nums) / 2

	if len(nums)%2 == 0 {

		median = (nums[i] + nums[i-1]) / 2

	} else {

		median = nums[i]
	}

	fmt.Println("median:", median)

}
