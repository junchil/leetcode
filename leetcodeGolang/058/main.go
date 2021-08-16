package main

func lengthOfLastWord(s string) int {
	length := len(s)

	if length == 0 {
		return 0
	}

	res := 0
	for i := length - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if res > 0 {
				return res
			}
		}
		res = res + 1
	}

	return res
}
