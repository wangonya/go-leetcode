// https://leetcode.com/problems/implement-strstr/

package goleetcode

import "strings"

func strStr(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}

// manual solution
// const notFound = -1

// func strStr(haystack string, needle string) int {
// 	if len(needle) == 0 {
// 		return 0
// 	}

// 	if len(needle) > len(haystack) {
// 		return notFound
// 	}

// 	for i := 0; i < len(haystack)-len(needle)+1; i++ {

// 		for j := 0; j < len(needle); j++ {
// 			if haystack[i+j] != needle[j] {
// 				break
// 			}

// 			if j == len(needle)-1 {
// 				return i
// 			}
// 		}
// 	}

// 	return notFound
// }
