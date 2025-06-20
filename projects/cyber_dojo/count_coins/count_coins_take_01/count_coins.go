package main

// CountCoins represents a coin counter that calculates ways to make change
type CountCoins struct {
}

var coins = map[string]int{
	"quarters": 25,
	"dimes":    10,
	"nickels":  5,
	"pennies":  1,
}

// Changes calculates the number of ways to make change for a given amount
func (cc *CountCoins) Changes(amount int) int {
	if amount < 1 {
		return 0
	}

	ways := make([]int, amount+1)
	ways[0] = 1

	for _, coin := range coins {
		for j := coin; j <= amount; j++ {
			ways[j] += ways[j-coin]
		}
	}

	return ways[amount]
}
