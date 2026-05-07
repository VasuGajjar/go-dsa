package easy

func MaxConsecutiveOne(nums []int) int {
	cnt, mxCnt := 0, 0

	for i := range nums {
		if nums[i] == 1 {
			cnt++
			mxCnt = max(mxCnt, cnt)
		} else {
			cnt = 0
		}
	}

	return mxCnt
}
