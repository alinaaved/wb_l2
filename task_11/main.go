package main

import (
	"fmt"
	"sort"
	"strings"
)

func sorting(first string) string {

	firstRune := []rune(first)

	sort.Slice(firstRune, func(i, j int) bool { return firstRune[i] < firstRune[j] })
	return string(firstRune)
}

func findAnagram(arrStrings []string) map[string][]string {
	res := make(map[string][]string)
	groups := make(map[string][]string)
	for i := range arrStrings {
		lower := strings.ToLower(arrStrings[i])
		sorted := sorting(lower)
		groups[sorted] = append(groups[sorted], lower)
	}
	for _, v := range groups {
		if len(v) > 1 {
			first := v[0]
			sort.Strings(v)
			res[first] = v
		}
	}
	return res
}

func main() {
	input := []string{
		"пятка", "пятак", "тяпка", "листок", "слиток", "столик", "стол",
	}
	res := findAnagram(input)
	for k, v := range res {
		fmt.Println(k, ":", v)
	}

}
