package anagram

import (
	"sort"
	"strings"
)

var a []string
var b string

func sorting(s string) string {
	c := strings.ToLower(s)
	t := []rune(c)

	sort.Slice(t, func(i, j int) bool {
		return t[i] < t[j]
	})
	return string(t)
}
func Detect(subject string, candidates []string) []string {
	a = make([]string, 0)
	h := strings.ToLower(subject)

	for _, v := range candidates {

		I := strings.ToLower(v)
		if len(subject) == len(v) && subject != v {
			if sorting(subject) == sorting(v) {
				a = append(a, v)
			}
		}

	}
	return a
}

func main() {
	// fmt.Println(sorting("pavan"))
	fmt.Println(Detect("Orchestra", []string{"cashregister", "Carthorse", "radishes"}))
	// v := []rune("done")
	// fmt.Println(v)
	// fmt.Println(sorting("done"))
}
