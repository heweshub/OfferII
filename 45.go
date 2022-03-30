package main

import (
	"strconv"
)

// 该题主要借助排序思想，解题关键在于排序规则
func minNumber(nums []int) string {
	res := make([]string, len(nums))
	for i, v := range nums {
		res[i] = strconv.Itoa(v)
	}
	// 匿名函数比较两个字符串拼接之后的大小
	compare := func(str1, str2 string) bool {
		num1, _ := strconv.Atoi(str1 + str2)
		num2, _ := strconv.Atoi(str2 + str1)

		if num1 < num2 {
			return true
		}
		return false
	}

	var quickSort func(strS []string, left int, right int)
	quickSort = func(strS []string, left int, right int) {
		if left >= right {
			return
		}
		l, r, pivot := left, right, strS[left]
		for l < r {
			// 正序大于倒叙，右指针--
			for l < r && compare(pivot, strS[r]) {
				r--
			}
			// 正序小于倒叙，左指针++
			for l < r && !compare(pivot, strS[l]) {
				l++
			}
			strS[l], strS[r] = strS[r], strS[l]
		}
		strS[l], strS[left] = strS[left], strS[l]
		quickSort(strS, left, l-1)
		quickSort(strS, l+1, right)
	}
	quickSort(res, 0, len(res)-1)
	ans := ""
	for _, s := range res {
		ans += s
	}
	return ans
}
