package boyer_moore_search

import (
	"unicode/utf8"
)

// build skip table of needle for Boyer-Moore search.
func Original_BuildSkipTable(needle string) map[rune]int {
	l := utf8.RuneCountInString(needle)
	runes := []rune(needle)

	table := make(map[rune]int)

	for i := 0; i < l-1; i++ {
		j := runes[i]
		table[j] = l - i - 1
	}

	return table
}

// search a needle in haystack and return count of needle.
// table is build by BuildSkipTable.
func Original_SearchBySkipTable(haystack, needle string, table map[rune]int) int {

	i, c := 0, 0
	hrunes := []rune(haystack)
	nrunes := []rune(needle)
	hl := utf8.RuneCountInString(haystack)
	nl := utf8.RuneCountInString(needle)

	if hl == 0 || nl == 0 || hl < nl {
		return 0
	}

	if hl == nl && haystack == needle {
		return 1
	}

loop:
	for i+nl <= hl {
		for j := nl - 1; j >= 0; j-- {
			if hrunes[i+j] != nrunes[j] {
				if _, ok := table[hrunes[i+j]]; !ok {
					if j == nl-1 {
						i += nl
					} else {
						i += nl - j - 1
					}
				} else {
					n := table[hrunes[i+j]] - (nl - j - 1)
					if n <= 0 {
						i++
					} else {
						i += n
					}
				}
				goto loop
			}
		}

		if _, ok := table[hrunes[i+nl-1]]; ok {
			i += table[hrunes[i+nl-1]]
		} else {
			i += nl
		}

		c++
	}

	return c
}

// search a needle in haystack and return count of needle.
func Original_Search(haystack, needle string) int {
	table := Original_BuildSkipTable(needle)
	return Original_SearchBySkipTable(haystack, needle, table)
}
