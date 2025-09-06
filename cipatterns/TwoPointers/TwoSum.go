package twopointers

import "fmt"

func TwoSumDefault(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

func TwoSum(nums []int, target int) []int {
	size := len(nums)
	indexOf := make(map[int]int)
	fmt.Printf("nums:%v", nums)
	result := make([]int, 0, size)
	for i := 0; i < size; i++ {
		complement := target - nums[i]
		_, ok := indexOf[complement]
		if ok {
			result = append(result, i)
			result = append(result, indexOf[complement])
			return result
		} else {
			indexOf[nums[i]] = i
		}
	}
	fmt.Printf("Reached end: %v", result)
	return result
}
