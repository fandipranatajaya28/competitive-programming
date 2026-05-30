package main

import "fmt"

// =============================================================================
// LeetCode 3161 – Block Placement Queries
// =============================================================================
//
// DATA STRUCTURES
// ───────────────
// 1. Fenwick tree as an ordered set
//    • Stores obstacle positions as a frequency array (0 or 1 at each index).
//    • fw.Sum(i)  = number of obstacles in [0, i-1]  (prefix count)
//    • fw.Kth(k)  = position of the k-th obstacle    (binary lift, O(log M))
//    • predecessor(x) = Kth( Sum(x+1) )   — last obstacle ≤ x
//    • successor(x)   = Kth( Sum(x+1)+1 ) — first obstacle > x
//
//    *** INDEXING CONVENTION ***
//    Real position p is stored at BIT index p+1.
//    So fw.Add(p+1, 1) marks position p, and fw.Kth(k)-1 recovers position p.
//    Sentinel for position 0: fw.Add(1, 1) at startup.
//
// 2. Segment tree (point-update, range-max)
//    • st[p] = length of the gap whose RIGHT endpoint is obstacle p.
//              i.e. st[p] = p − prev(p)
//    • For type-2 query with bound x, querying [0, x] is safe: only
//      obstacle positions have non-zero values, so [prev+1 .. x] slots
//      are always 0 and don't pollute the max.
//
// TYPE-1 insert obstacle at x
//    l  = predecessor of x  (guaranteed to exist; sentinel 0 is always present)
//    r  = successor of x    (-1 if none)
//    → fw.Add(x+1, 1)
//    → st.update(x, x−l)         gap ending at x
//    → st.update(r, r−x)         gap ending at r shrinks  (if r exists)
//
// TYPE-2 query [2, x, sz]
//    prev     = predecessor(x)         last obstacle ≤ x
//    partial  = x − prev               usable width of gap straddling x
//    fullGap  = st.query(0, x)         best complete gap with right end ≤ x
//    answer   = max(partial, fullGap) ≥ sz
//
// COMPLEXITY: O(n log M) time, O(M) space  where M = max coordinate
// =============================================================================

// ─── Fenwick tree ─────────────────────────────────────────────────────────────

type Fenwick struct {
	n   int
	bit []int
}

func NewFenwick(n int) *Fenwick {
	return &Fenwick{n: n, bit: make([]int, n+1)}
}

// Add adds val to index idx (1-indexed).
func (f *Fenwick) Add(idx, val int) {
	for ; idx <= f.n; idx += idx & -idx {
		f.bit[idx] += val
	}
}

// Sum returns the prefix sum [1, idx].
func (f *Fenwick) Sum(idx int) int {
	res := 0
	for ; idx > 0; idx -= idx & -idx {
		res += f.bit[idx]
	}
	return res
}

// Kth returns the smallest index whose prefix sum reaches k.
// Uses binary lifting: start from the highest power of 2 ≤ n and walk down.
// This is O(log n) with no extra memory — tighter than calling Sum in a loop.
func (f *Fenwick) Kth(k int) int {
	// Find the highest power of 2 ≤ f.n as the starting step.
	step := 1
	for step<<1 <= f.n {
		step <<= 1
	}
	idx := 0
	for d := step; d > 0; d >>= 1 {
		next := idx + d
		// If the subtree rooted at next has fewer than k elements,
		// skip it and subtract its count from k.
		if next <= f.n && f.bit[next] < k {
			idx = next
			k -= f.bit[next]
		}
	}
	// idx is the last position where prefix sum < k, so idx+1 is the answer.
	return idx + 1
}

// ─── Segment tree (point-update, range-max) ───────────────────────────────────

type SegTree struct {
	n    int
	tree []int
}

func NewSegTree(n int) *SegTree {
	return &SegTree{n: n, tree: make([]int, 4*n+4)}
}

func (s *SegTree) update(node, l, r, pos, val int) {
	if l == r {
		s.tree[node] = val
		return
	}
	mid := (l + r) >> 1
	if pos <= mid {
		s.update(node<<1, l, mid, pos, val)
	} else {
		s.update(node<<1|1, mid+1, r, pos, val)
	}
	// Pull up: parent = max of children.
	if s.tree[node<<1] > s.tree[node<<1|1] {
		s.tree[node] = s.tree[node<<1]
	} else {
		s.tree[node] = s.tree[node<<1|1]
	}
}

func (s *SegTree) query(node, l, r, ql, qr int) int {
	if ql > r || qr < l {
		return 0
	}
	if ql <= l && r <= qr {
		return s.tree[node]
	}
	mid := (l + r) >> 1
	a := s.query(node<<1, l, mid, ql, qr)
	b := s.query(node<<1|1, mid+1, r, ql, qr)
	if a > b {
		return a
	}
	return b
}

func (s *SegTree) Update(pos, val int) { s.update(1, 0, s.n, pos, val) }
func (s *SegTree) Query(ql, qr int) int {
	if ql > qr {
		return 0
	}
	return s.query(1, 0, s.n, ql, qr)
}

// ─── Solution ─────────────────────────────────────────────────────────────────

func getResults(queries [][]int) []bool {
	// Compute the actual max coordinate so we allocate only what we need.
	// Tight bound → smaller arrays → better cache behaviour.
	mx := 0
	for _, q := range queries {
		if q[1] > mx {
			mx = q[1]
		}
	}

	// BIT is 1-indexed; we need indices up to mx+1 (for position mx).
	// Add +2 for safety when calling Sum(mx+2) to count all elements.
	fw := NewFenwick(mx + 2)
	st := NewSegTree(mx + 1)

	// Sentinel: treat position 0 as a permanent obstacle.
	// Stored at BIT index 0+1 = 1.
	fw.Add(1, 1)

	ans := make([]bool, 0, len(queries))

	for _, q := range queries {
		t := q[0]
		x := q[1]

		if t == 1 {
			// ── TYPE-1: insert obstacle at x ─────────────────────────────

			// countUpToX = number of obstacles at positions [0, x] (BIT indices [1, x+1]).
			// Used for both predecessor and successor — computed once, used twice.
			// IMPORTANT: must be Sum(x+1), not Sum(x+2). Sum(x+2) would include
			// position x+1, which may be occupied, falsely hiding the right neighbour.
			countUpToX := fw.Sum(x + 1)

			// predecessor: last obstacle <= x.
			leftPos := fw.Kth(countUpToX) - 1 // BIT index -> real position

			// successor: first obstacle > x.
			totalOccupied := fw.Sum(mx + 2) // all obstacles in [0, mx]
			rightPos := -1
			if countUpToX < totalOccupied {
				rightPos = fw.Kth(countUpToX+1) - 1
			}

			// Register x in the BIT (BIT index = x+1).
			fw.Add(x+1, 1)

			// Update segment tree: gap ending at x = x - leftPos.
			st.Update(x, x-leftPos)

			// The gap that used to end at rightPos (= rightPos - leftPos)
			// is now split; update it to rightPos - x.
			if rightPos != -1 {
				st.Update(rightPos, rightPos-x)
			}

		} else {
			// ── TYPE-2: can block of size sz fit in [0, x]? ──────────────
			sz := q[2]

			// Find the last obstacle at or before x.
			leftCount := fw.Sum(x + 1)
			leftPos := fw.Kth(leftCount) - 1

			// Partial gap: the open space from leftPos to x (gap straddles x).
			partial := x - leftPos

			// Best complete gap with right endpoint anywhere in [0, x].
			// Positions between leftPos+1 and x have st value 0 (never written),
			// so querying [0, x] rather than [0, leftPos] is safe and avoids
			// a second Kth call.
			bestPrefix := st.Query(0, x)

			ans = append(ans, partial >= sz || bestPrefix >= sz)
		}
	}

	return ans
}

// ─── Smoke tests ──────────────────────────────────────────────────────────────

func main() {
	// Example 1: expected [false true true]
	fmt.Println(getResults([][]int{{1, 2}, {2, 3, 3}, {2, 3, 1}, {2, 2, 2}}))

	// Example 2: expected [true true false]
	fmt.Println(getResults([][]int{{1, 7}, {2, 7, 6}, {1, 2}, {2, 7, 5}, {2, 7, 6}}))

	// Example 3: expected [true false true]
	fmt.Println(getResults([][]int{{2, 5, 4}, {1, 2}, {2, 3, 4}, {2, 4, 2}}))
}
