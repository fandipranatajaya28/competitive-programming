package main

import "fmt"

func countMajoritySubarrays(nums []int, target int) int64 {
	n := len(nums)

	// Prefix sum can range from -n to +n.
	// Fenwick tree index must be positive, so we shift by offset.
	offset := n + 2

	// Fenwick size needs enough room for shifted prefix sums.
	// Possible shifted range:
	// -n + offset to n + offset
	size := 2*n + 5

	bit := NewFenwick(size)

	var answer int64 = 0
	prefixSum := 0

	// Add prefix sum 0 before processing any element.
	// This represents an empty prefix before index 0.
	bit.Add(prefixSum+offset, 1)

	for _, num := range nums {
		// Convert current number to +1 or -1.
		//
		// If num == target, it helps target become majority.
		// Otherwise, it works against target.
		if num == target {
			prefixSum++
		} else {
			prefixSum--
		}

		// We need to count previous prefix sums that are smaller
		// than the current prefix sum.
		//
		// If previousPrefix < currentPrefix:
		// currentPrefix - previousPrefix > 0
		//
		// That means the subarray has positive transformed sum,
		// so target is majority in that subarray.
		currentIndex := prefixSum + offset
		answer += bit.Query(currentIndex - 1)

		// Add current prefix sum to Fenwick tree
		// for future subarrays.
		bit.Add(currentIndex, 1)
	}

	return answer
}

// Fenwick Tree / Binary Indexed Tree.
//
// Supports:
// 1. Add value at an index.
// 2. Query prefix sum from 1 to index.
//
// Here, it stores how many times each prefix sum has appeared.
type Fenwick struct {
	tree []int64
}

func NewFenwick(size int) *Fenwick {
	return &Fenwick{
		tree: make([]int64, size+1),
	}
}

// Add delta to index idx.
func (f *Fenwick) Add(idx int, delta int64) {
	for idx < len(f.tree) {
		f.tree[idx] += delta
		idx += idx & -idx
	}
}

// Query returns sum of values from index 1 to idx.
func (f *Fenwick) Query(idx int) int64 {
	var sum int64 = 0

	for idx > 0 {
		sum += f.tree[idx]
		idx -= idx & -idx
	}

	return sum
}

func main() {
	nums := []int{1, 2, 2, 3}
	target := 2

	fmt.Println(countMajoritySubarrays(nums, target)) // Output: 5
}
