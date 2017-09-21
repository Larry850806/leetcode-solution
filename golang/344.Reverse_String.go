// https://leetcode.com/problems/reverse-string/description/

package main

import "bytes"

func reverseString(s string) string {
	var reversedBuffer bytes.Buffer

	length := len(s)
	for i := range s {
		reversedBuffer.WriteByte(s[length-i-1])
	}

	return reversedBuffer.String()
}

// func reverseString(s string) string {
// 	slen := len(s)
// 	if slen <= 1 {
// 		return s
// 	}

// 	bytes := []byte(s)
// 	for i := 0; i < slen/2; i++ {
// 		bytes[i], bytes[slen-1-i] = bytes[slen-1-i], bytes[i]
// 	}

// 	return string(bytes)
// }
