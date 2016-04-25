// Source : https://oj.leetcode.com/problems/two-sum/
// Author : DeShuai Ma
// Date   : 2016-04-26

func twoSum(nums []int, target int) []int {
	for i := 0; i <= len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}
