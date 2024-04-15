package degree_of_an_array

import "fmt"

func DoAlg() {
	arr := []int{1, 2, 2, 3, 1, 4, 2}
	fmt.Println(findShortestSubArray(arr))
}

func findShortestSubArray(nums []int) int {
	/*
		1.先确定最大的度
		2.再找数组里拥有该度的连续子数组
		3.最后再找到最短的连续子数组
		===========================
		a.所以在获取度时，包含同样长度的度的连续子数组头一定是该数出现的第一个索引和最后一个索引，然后计算两个索引的差值，最小的就是该题的解
		b.希望经历1~2次循环就能取得结果
		c.需要记录：每个数字当前出现的次数和其起始索引
	*/
	totalDegree := map[int]int{} // 每个数字出现的次数，最后通过遍历该值取得最大的度
	firstIndex := map[int]int{}  // 每个数字首次出现的索引
	lastIndex := map[int]int{}   // 每个数字最后出现的索引

	for i, v := range nums {
		if _, ok := totalDegree[v]; !ok {
			totalDegree[v] = 1 // 默认从1开始
			firstIndex[v] = i  // 记录首先出现的索引
		} else {
			totalDegree[v]++
			lastIndex[v] = i // 记录最后出现的索引
		}
	}

	// 找到最大的度
	maxDegree := 0
	minLen := 0
	for n, t := range totalDegree {
		if t == 1 {
			continue // 过滤度为1的情况
		} else if t > maxDegree {
			maxDegree = t
			minLen = lastIndex[n] - firstIndex[n]
		} else if len := lastIndex[n] - firstIndex[n]; len >= 0 && len < minLen && t == maxDegree {
			minLen = lastIndex[n] - firstIndex[n]
		}
	}

	return minLen + 1 // 题目问元素个数，算式是索引相减，需要+1
}
