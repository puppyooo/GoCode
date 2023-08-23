package main

import "strconv"

func translateNum(num int) int {
	src := strconv.Itoa(num)
	length := len(src)
	if num < 10 {
		return 1
	}
	dp := make([]int, length)

	dp[0] = 1
	if (src[1] <= '9' && src[0] == '1') || (src[1] <= '5' && src[0] == '2') {
		dp[1] = 2
	} else {
		dp[1] = 1
	}

	for i := 2; i < length; i++ {
		if (src[i] <= '9' && src[i-1] == '1') || (src[i] <= '5' && src[i-1] == '2') {
			dp[i] = dp[i-1] + dp[i-2]
		} else {
			dp[i] = dp[i-1]
		}
	}
	return dp[length-1]
}

func main() {
	println(translateNum(25))

}
