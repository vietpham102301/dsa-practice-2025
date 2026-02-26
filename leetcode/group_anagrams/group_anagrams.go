package main

import (
	"fmt"
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	groupMap := map[string][]string{}

	for _, str := range strs {
		b := []byte(str)
		sort.Slice(b, func(i, j int) bool {
			return b[i] < b[j]
		})

		sortedS := string(b)
		groupMap[sortedS] = append(groupMap[sortedS], str)
	}

	res := make([][]string, 0, len(groupMap))

	for _, g := range groupMap {
		res = append(res, g)
	}

	return res
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}

	res := groupAnagrams(strs)
	fmt.Println(res)
}
