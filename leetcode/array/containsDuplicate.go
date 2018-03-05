package main

import (
	"fmt"
	"sort"
)

func ContainsDuplicate(nums []int) bool {
	sort.Ints(nums)
	for idx := 0; idx < len(nums)-1; idx++ {
		if nums[idx] == nums[idx+1] {
			return true
		}
	}
	return false
}

func main() {
	nums := []int{3, 1, 4, 3, 2}
	nums2 := []int{3, 4, 6, 5}

	fmt.Println(ContainsDuplicate(nums[:]))
	fmt.Println(ContainsDuplicate(nums2[:]))
}
