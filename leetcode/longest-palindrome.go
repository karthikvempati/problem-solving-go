/*
Leetcode Problem 409 - https://leetcode.com/problems/longest-palindrome/

Given a string s which consists of lowercase or uppercase letters, return the length of the longest palindrome that can be built with those letters.

Letters are case sensitive, for example, "Aa" is not considered a palindrome here.

 

Example 1:

Input: s = "abccccdd"
Output: 7
Explanation:
One longest palindrome that can be built is "dccaccd", whose length is 7.
Example 2:

Input: s = "a"
Output: 1
Example 3:

Input: s = "bb"
Output: 2
 

Constraints:

1 <= s.length <= 2000
s consists of lowercase and/or uppercase English letters only.

*/

package leetcode

func longestPalindrome(s string) int {
    m := make(map[rune]int)
    length := 0
    maxOdd := 0
    a := 0
    for _,c := range s{
        m[c]++
    }
    for _,val := range m{ 
        a++
        if(val % 2 == 0){
            length = length+val
        } else if(val > maxOdd){ 
            if maxOdd != 0 {
                length = length + maxOdd -1
            }
            maxOdd = val
        } else {
            length = length + val - 1
        }
    }
    
    return maxOdd + length
}