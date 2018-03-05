package main

import (
  "fmt"
  "sort"
)

func ArrayPairSum(nums []int)int{
  sort.Ints(nums)
  maxSum := 0
  for i:=0; i < len(nums); i++{
    if i % 2 == 0 {
      maxSum += nums[i]
    }
  }
  return maxSum;
}

func main() {
  nums := []int{1, 3, 4, 2}
  fmt.Println(ArrayPairSum(nums))
}
