package main

func lengthOfLongestSubstring(s string) int {
	mp := make(map[byte]int)

	max := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}
	ret := 0
	left := 0
	for right, b := range s {
		mp[byte(b)]++
		for mp[byte(b)] > 1 {
			mp[s[left]]--
			left++
		}
		ret = max(ret, right-left+1)
	}
	return ret
}

func main() {
	println(lengthOfLongestSubstring("abcabcbb"))
}
