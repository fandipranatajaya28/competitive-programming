package main

import "fmt"

// func findThePrefixCommonArray(A []int, B []int) []int {
// 	var (
// 		result               []int
// 		numOccuranceCountMap = make(map[int]int)
// 		commonNumCount       int
// 	)

// 	for i := 0; i < len(A); i++ {
// 		if A[i] == B[i] {
// 			commonNumCount++
// 		} else {
// 			numOccuranceCountMap[A[i]]++
// 			numOccuranceCountMap[B[i]]++
// 			if numOccuranceCountMap[A[i]] == 2 {
// 				commonNumCount++
// 			}
// 			if numOccuranceCountMap[B[i]] == 2 {
// 				commonNumCount++
// 			}
// 		}
// 		result = append(result, commonNumCount)
// 	}

// 	return result
// }

func findThePrefixCommonArray(A []int, B []int) []int {
	n := len(A)
	bitmask, common := 0, 0
	ans := make([]int, n)

	for i := 0; i < n; i++ {
		if bitmask&(1<<A[i]) != 0 {
			common++
		} else {
			bitmask |= (1 << A[i])
		}

		if bitmask&(1<<B[i]) != 0 {
			common++
		} else {
			bitmask |= (1 << B[i])
		}

		ans[i] = common
	}
	return ans
}

func main() {
	fmt.Println(findThePrefixCommonArray([]int{1, 3, 2, 4}, []int{3, 1, 2, 4}))
}
