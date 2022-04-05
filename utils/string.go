package utils

import (
	"regexp"
	"sort"
	"strings"

	"github.com/dexyk/stringosim"
)

func SimilarText(str1, str2 string) float64 {
	similarScore := 1 - stringosim.Cosine([]rune(str1), []rune(str2))

	return float64(int(similarScore*10000)) / 100
}

func RemoveTagFromString(str string) string {
	const pattern = `(<\/?[a-zA-A]+?[^>]*\/?>)*`
	r := regexp.MustCompile(pattern)
	groups := r.FindAllString(str, -1)
	sort.Slice(groups, func(i, j int) bool {
		return len(groups[i]) > len(groups[j])
	})
	for _, group := range groups {
		if strings.TrimSpace(group) != "" {
			str = strings.ReplaceAll(str, group, "")
		}
	}
	return str
}

func RemoveBracketsFromString(str string) string {
	const pattern = `(\(.+?\))`
	r := regexp.MustCompile(pattern)
	groups := r.FindAllString(str, -1)
	sort.Slice(groups, func(i, j int) bool {
		return len(groups[i]) > len(groups[j])
	})
	for _, group := range groups {
		if strings.TrimSpace(group) != "" {
			str = strings.ReplaceAll(str, group, "")
		}
	}
	return str
}
