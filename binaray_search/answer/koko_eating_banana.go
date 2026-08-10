package answer

import "math"

func MinEatingSpeed(piles []int, h int) int {
	if n := len(piles); n == 0 || n > h {
		return -1
	}

	mx := math.MinInt
	for i := range piles {
		mx = max(mx, piles[i])
	}

	low, high := 1, mx

	for low <= high {
		mid := (low + high) / 2
		if isEatingPossible(piles, h, mid) {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return low
}

func isEatingPossible(piles []int, h, n int) bool {
	cnt := 0

	for i := range piles {
		cnt += piles[i] / n

		if piles[i]%n != 0 {
			cnt++
		}

		if cnt > h {
			return false
		}
	}

	return true
}
