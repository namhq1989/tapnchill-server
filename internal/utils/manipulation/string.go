package manipulation

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"

	"golang.org/x/text/unicode/norm"
)

func ToTsQueryTerms(text string) string {
	text = Slugify(text)

	var (
		parts      = strings.Split(text, "-")
		totalParts = len(parts)
	)

	if totalParts == 1 {
		return fmt.Sprintf("%s:*", text)
	}

	terms := ""
	for i, part := range parts {
		if i == totalParts-1 {
			terms += fmt.Sprintf("%s:*", part)
		} else {
			terms += fmt.Sprintf("%s:* & ", part)
		}
	}

	return terms
}

func Slugify(text string) string {
	noDiacritics := removeDiacritics(text)

	noDiacritics = strings.ReplaceAll(noDiacritics, "Đ", "D")
	noDiacritics = strings.ReplaceAll(noDiacritics, "đ", "d")

	lower := strings.ToLower(noDiacritics)
	hyphens := strings.ReplaceAll(lower, " ", "-")
	reg, err := regexp.Compile("[^a-zA-Z0-9-]+")
	if err != nil {
		return ""
	}
	safe := reg.ReplaceAllString(hyphens, "")
	return strings.Trim(safe, "-")
}

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

func ValidateAndCleanHumanName(s string) string {
	cleanedName := strings.Join(strings.Fields(s), " ")

	for _, ch := range cleanedName {
		if !unicode.IsLetter(ch) && !unicode.IsSpace(ch) {
			return ""
		}
	}

	return cleanedName
}

func RemoveSpacesAndValidateNumber(s string) string {
	digitsOnly := strings.ReplaceAll(s, " ", "")

	for _, ch := range digitsOnly {
		if !unicode.IsDigit(ch) {
			return ""
		}
	}

	return digitsOnly
}
