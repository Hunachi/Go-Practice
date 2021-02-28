// Implemented by Hunachi

package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	ans := make(map[string]int)
	texts := strings.Fields(s)
	for _, v := range texts {
		ans[v] += 1
	}
	return ans
}

func main() {
	wc.Test(WordCount)
}