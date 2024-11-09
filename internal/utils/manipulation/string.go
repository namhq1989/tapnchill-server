package manipulation

import (
	"regexp"
	"strings"
	"unicode"

	"golang.org/x/text/unicode/norm"
)

func removeDiacritics(input string) string {
	t := norm.NFD.String(input)
	var b strings.Builder
	b.Grow(len(t))
	for _, r := range t {
		if unicode.Is(unicode.Mn, r) {
			continue
		}
		b.WriteRune(r)
	}
	return b.String()
}

func NormalizeText(input string) string {
	result := removeDiacritics(input)

	result = strings.ReplaceAll(result, "đ", "d")
	result = strings.ReplaceAll(result, "Đ", "d")

	result = strings.ToLower(result)

	re := regexp.MustCompile(`[^a-z0-9\s]+`)
	result = re.ReplaceAllString(result, " ")

	words := strings.Fields(result)

	uniqueWords := make([]string, 0, len(words))
	wordSet := make(map[string]bool)
	for _, word := range words {
		if !wordSet[word] {
			wordSet[word] = true
			uniqueWords = append(uniqueWords, word)
		}
	}

	return strings.Join(uniqueWords, " ")
}
