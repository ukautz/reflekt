package reflekt

import (
	"regexp"
	"strings"
)

var snakeCaseRegex = regexp.MustCompile(`(^[\p{Ll}\p{N}]+)?(\p{Lu}[\p{Ll}\p{N}]+)`)

func snakeCase(str string) string {
	chunks := snakeCaseRegex.FindAllStringSubmatch(str, -1)
	if chunks == nil {
		return strings.ToLower(str)
	}
	sep := "_"
	found := []string{}
	for _, chunk := range chunks {
		for i := 1; i < len(chunk); i++ {
			if len(chunk[i]) > 0 {
				found = append(found, strings.ToLower(chunk[i]))
			}
		}
	}
	return strings.Join(found, sep)
}
