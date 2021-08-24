package main

import (
	"fmt"
	"strings"
)

func strStr(haystack string, needle string) int {
	leng := len(needle)

	if needle == "" {
		return 0
	}

	if !strings.Contains(haystack, needle) {
		return -1
	} else {
		for index := range haystack {
			if haystack[index:index+leng] == needle {
				return index
			}
			continue
		}
	}
	return 0
}

func main() {
	res := strStr("helolo", "ll")
	fmt.Println(res)
}
