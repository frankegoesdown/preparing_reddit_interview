package main

func main() {

}

func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for _, w := range wordDict {
			if i-len(w) >= 0 && dp[i-len(w)] && s[i-len(w):i] == w {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}
