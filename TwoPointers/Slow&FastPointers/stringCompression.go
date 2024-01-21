package SlowAndFastPointers

import "strconv"

// link - https://leetcode.com/problems/string-compression/description/

// Given an array of characters chars, compress it using the following algorithm:

// Begin with an empty string s. For each group of consecutive repeating characters in chars:

// If the group's length is 1, append the character to s.
// Otherwise, append the character followed by the group's length.
// The compressed string s should not be returned separately, but instead, be stored in the input character array chars. Note that group lengths that are 10 or longer will be split into multiple characters in chars.


func compress(chars []byte) int {
    start, end, currentPointer, n := 0, 0, 0, len(chars)
	for end < n {
		if chars[start] != chars[end] {

			count := end - start
			currentPointer++
			if count > 1 {
				for _, val := range []byte(strconv.Itoa(count)) {
					chars[currentPointer] = val
					currentPointer++
				}
			}
			chars[currentPointer] = chars[end]
			start = end
		}
		end++
	}
    // residue 
    count := end - start
	if count > 1 {
		currentPointer++
		for _, val := range []byte(strconv.Itoa(count)) {
			chars[currentPointer] = val
			currentPointer++
		}
		currentPointer--
	}
    return currentPointer+1
}