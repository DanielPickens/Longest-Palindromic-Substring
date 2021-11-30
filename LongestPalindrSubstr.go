//We can start by traversing from each index and growing from both sides till it is a palindrome.
//We can then get the maximum length of odd length palindrome centered at i as first pointer.
//Then we can get maximum length even length palindrome centered at i
//We will go through all possible element in string and find the longest palindrome amongst them.
//And to calculate all the possible substrings in a string of length n 
//We need to check all the substrings for it to be a Palindrome, and we know we will have to go through all the n characters to ascertain this.

//time complexity: O(n)
//Space complexity: O(n)

func longestPalindrome(s string) string {
  if len(s) == 0 {
		return ""
	}
	var (
		start, end int
		maxLen     int
	)
	for i := 0; i < len(s); i++ {
		//odd
		j, k := i-1, i+1
		for j >= 0 && k < len(s) && s[j] == s[k] {
			j--
			k++
		}
		if k-j-1 > maxLen {
			maxLen = k - j - 1
			start = j + 1
			end = k
		}
		//even
		j, k = i, i+1
		for j >= 0 && k < len(s) && s[j] == s[k] {
			j--
			k++
		}
		if k-j-1 > maxLen {
			maxLen = k - j - 1
			start = j + 1
			end = k
		}
	}
	return s[start : end]
}



/*
Success
Details 
Runtime: 4 ms, faster than 94.57% of Go online submissions for Longest Palindromic Substring.
Memory Usage: 2.7 MB, less than 100.00% of Go online submissions for Longest Palindromic Substring.
*/
