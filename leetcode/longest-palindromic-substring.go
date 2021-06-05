/*

Leetcode 5 - https://leetcode.com/problems/longest-palindromic-substring/

Given a string s, return the longest palindromic substring in s.

Example 1:

Input: s = "babad"
Output: "bab"
Note: "aba" is also a valid answer.
Example 2:

Input: s = "cbbd"
Output: "bb"
Example 3:

Input: s = "a"
Output: "a"
Example 4:

Input: s = "ac"
Output: "a"


Constraints:

1 <= s.length <= 1000
s consist of only digits and English letters (lower-case and/or upper-case),
*/

package leetcode

func longestPalindromicSubString(s string) string {
	p := ""
	for i := 0; i < len(s); i++ {
		for j := i + len(p); j < len(s); j++ {
			if isPalindrome(s[i:j+1]) && j-i+1 > len(p) {
				p = s[i : j+1]
			}
		}
	}

	return p
}

func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i <= j; {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
