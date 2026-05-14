package medium

func MaxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	res, mn := 0, prices[0]

	for i := 1; i < len(prices); i++ {
		res = max(res, prices[i]-mn)
		mn = min(mn, prices[i])
	}

	return res
}
