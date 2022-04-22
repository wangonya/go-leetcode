// https://leetcode.com/problems/longest-common-prefix/

package goleetcode

import "sort"

// divide and conquer

func findCommonPrefix(s1 string, s2 string) string {
	result := ""
	n1, n2 := len(s1), len(s2)
	i, j := 0, 0

	for {
		if i > n1-1 || j > n2-1 {
			break
		}

		if s1[i] != s2[i] {
			break
		}

		result += string(s1[i])
		i, j = i+1, j+1
	}

	return result
}

func longestCommonPrefixUtil(strs []string, low int, high int) string {
	result := ""
	if low == high {
		result = strs[low]
	}

	if high > low {
		mid := low + (high-low)/2
		s1 := longestCommonPrefixUtil(strs, low, mid)
		s2 := longestCommonPrefixUtil(strs, mid+1, high)
		result = findCommonPrefix(s1, s2)
	}
	return result
}

func longestCommonPrefix2(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}
	
	return longestCommonPrefixUtil(strs, 0, len(strs)-1)
}

// simple loop and compare

func longestCommonPrefix(strs []string) string {
    	sort.Strings(strs)

    first, rest := strs[:1][0], strs[1:]
	prefix := first

	for _, word := range rest {
		for i, letter := range word {
			if i >= len(prefix) {
				break
			}
			if string(letter) == string(prefix[i]) {
				continue
			} else {
				prefix = prefix[:i]
				break
			}
		}
	}
	return prefix
}
