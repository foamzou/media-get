package utils

import (
	"regexp"
	"sort"
	"strings"
)

func SimilarText(str1, str2 string) float64 {
	l1, l2 := len(str1), len(str2)
	if l1+l2 == 0 {
		return 0
	}
	sim := similarText(str1, str2, l1, l2)
	return float64(sim*200) / float64(l1+l2)
}

func similarText(str1, str2 string, len1, len2 int) int {
	var sum, max int
	pos1, pos2 := 0, 0

	for i := 0; i < len1; i++ {
		for j := 0; j < len2; j++ {
			for l := 0; (i+l < len1) && (j+l < len2) && (str1[i+l] == str2[j+l]); l++ {
				if l+1 > max {
					max = l + 1
					pos1 = i
					pos2 = j
				}
			}
		}
	}

	if sum = max; sum > 0 {
		if pos1 > 0 && pos2 > 0 {
			sum += similarText(str1, str2, pos1, pos2)
		}
		if (pos1+max < len1) && (pos2+max < len2) {
			s1 := []byte(str1)
			s2 := []byte(str2)
			sum += similarText(string(s1[pos1+max:]), string(s2[pos2+max:]), len1-pos1-max, len2-pos2-max)
		}
	}

	return sum
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
