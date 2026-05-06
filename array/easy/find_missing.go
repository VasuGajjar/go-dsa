package easy

func FindMissingWithSum(nums []int) int {
	total, sum := 0, 0

	for i, v := range nums {
		total += (i + 1)
		sum += v
	}

	return total - sum
}

func FindMissingWithXor(nums []int) int {
	val := 0

	for i, v := range nums {
		val ^= (i + 1)
		val ^= v
	}

	return val
}
