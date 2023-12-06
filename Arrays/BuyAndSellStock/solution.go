package BuyAndSellStock

// link - https://leetcode.com/problems/best-time-to-buy-and-sell-stock/description/
// TC - O(n)
// SC - O(1)

func MaxProfit(prices []int) int {
	profit, n := 0, len(prices)
	if n == 0 {
		return profit
	}
	buyPrice := prices[0]
	for i := 1; i < n; i++ {
		// if we find a better buy price (as min as possible), we pick that
		if prices[i] < buyPrice {
			buyPrice = prices[i]
		} else {
			// or we find some profit at currentIndex given buyPrice, we will 
			// update our profit if our newer profit is greater
			profit = max(profit, prices[i]-buyPrice)
		}
	}
	return profit
}

func max(a, b int) int {
	if a > b {return a}
	return b
}