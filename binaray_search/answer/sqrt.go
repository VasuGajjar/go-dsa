package answer

func SquareRoot(num int) int {
	low, high := 0, max(1, num/2)

	for low <= high {
		mid := (low + high) / 2
		if val := Pow(mid, 2); val == num {
			return mid
		} else if val < num {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return high
}

func Pow(num, exp int) int {
	if exp <= 0 {
		return 1
	}
	if exp == 1 {
		return num
	}

	ans := Pow(num, exp/2)
	ans *= ans

	if exp%2 == 1 {
		ans *= num
	}

	return ans
}
