package main

func maxProfit(prices []int) int {
	var max = func(a int, b int) int {
		if a > b {
			return a
		}
		return b
	}

	maxprofit := 0
	l := 0
	for r := 0; r < len(prices); r++ {
		if prices[l] > prices[r] {
			l = r
		} else {
			maxprofit = max(maxprofit, prices[r]-prices[l])
		}
	}
	return maxprofit
}

func main() {

}
