package main

import (
	"container/heap"
	"fmt"
	"math/bits"
)

func maxTotalValue(nums []int, k int) int64 {
	n := len(nums)

	// Sparse table supports O(1) range max/min query.
	st := NewSparseTable(nums)

	// Max heap stores the current best subarray for each left index.
	pq := &PriorityQueue{}
	heap.Init(pq)

	// For each left index l, the largest value is usually at r = n-1,
	// because extending the subarray cannot decrease max-min.
	for l := 0; l < n; l++ {
		val := int64(st.QueryMax(l, n-1) - st.QueryMin(l, n-1))
		heap.Push(pq, &Item{
			value: val,
			left:  l,
			right: n - 1,
		})
	}

	var answer int64 = 0

	// Pick exactly k distinct subarrays with the largest values.
	for i := 0; i < k; i++ {
		curr := heap.Pop(pq).(*Item)

		// Add the current largest available subarray value.
		answer += curr.value

		// Move to the next smaller candidate for the same left index.
		//
		// Current subarray is nums[left..right].
		// Next candidate for the same left is nums[left..right-1].
		//
		// This remains distinct because right changes.
		if curr.right > curr.left {
			nextRight := curr.right - 1
			nextValue := int64(st.QueryMax(curr.left, nextRight) - st.QueryMin(curr.left, nextRight))

			heap.Push(pq, &Item{
				value: nextValue,
				left:  curr.left,
				right: nextRight,
			})
		}
	}

	return answer
}

// ---------------- Sparse Table ----------------
//
// Sparse table lets us query max/min in nums[l..r] in O(1)
// after O(n log n) preprocessing.

type SparseTable struct {
	logTable []int
	maxTable [][]int
	minTable [][]int
}

func NewSparseTable(nums []int) *SparseTable {
	n := len(nums)
	maxLog := bits.Len(uint(n))

	logTable := make([]int, n+1)
	for i := 2; i <= n; i++ {
		logTable[i] = logTable[i/2] + 1
	}

	maxTable := make([][]int, maxLog)
	minTable := make([][]int, maxLog)

	maxTable[0] = make([]int, n)
	minTable[0] = make([]int, n)

	for i := 0; i < n; i++ {
		maxTable[0][i] = nums[i]
		minTable[0][i] = nums[i]
	}

	for level := 1; level < maxLog; level++ {
		length := 1 << level
		half := length >> 1

		maxTable[level] = make([]int, n-length+1)
		minTable[level] = make([]int, n-length+1)

		for i := 0; i+length <= n; i++ {
			maxTable[level][i] = max(
				maxTable[level-1][i],
				maxTable[level-1][i+half],
			)

			minTable[level][i] = min(
				minTable[level-1][i],
				minTable[level-1][i+half],
			)
		}
	}

	return &SparseTable{
		logTable: logTable,
		maxTable: maxTable,
		minTable: minTable,
	}
}

func (st *SparseTable) QueryMax(left int, right int) int {
	length := right - left + 1
	level := st.logTable[length]

	return max(
		st.maxTable[level][left],
		st.maxTable[level][right-(1<<level)+1],
	)
}

func (st *SparseTable) QueryMin(left int, right int) int {
	length := right - left + 1
	level := st.logTable[length]

	return min(
		st.minTable[level][left],
		st.minTable[level][right-(1<<level)+1],
	)
}

// ---------------- Priority Queue / Max Heap ----------------

type Item struct {
	value int64
	left  int
	right int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

// Greater value should come first, so this is a max heap.
func (pq PriorityQueue) Less(i int, j int) bool {
	return pq[i].value > pq[j].value
}

func (pq PriorityQueue) Swap(i int, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	*pq = append(*pq, x.(*Item))
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)

	item := old[n-1]
	*pq = old[:n-1]

	return item
}

// ---------------- Helpers ----------------

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxTotalValue([]int{4, 2, 5, 1}, 3))
}
