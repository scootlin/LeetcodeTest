package main

import (
	"fmt"
	"sort"
)

func main() {
	// groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
	fmt.Println("INT to Roman : ", groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}

type sortRune []byte

func (s sortRune) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRune) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRune) Len() int {
	return len(s)
}

func groupAnagrams(strs []string) [][]string {
	var output [][]string
	dist := make(map[string]int)
	for _, str := range strs {
		chars := []byte(str)
		sort.Sort(sortRune(chars))
		sorted := string(chars)
		if idx, ok := dist[sorted]; ok {
			output[idx] = append(output[idx], str)
		} else {
			dist[sorted] = len(output)
			output = append(output, []string{str})
		}
	}
	return output
}
