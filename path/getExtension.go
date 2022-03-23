package path

import "strings"

func GetExtension(filename string) string {
	words := strings.Split(filename, ".")
	if len(words) == 2 {
		return words[1]
	} else {
		if len(words) > 2 {
			return words[len(words)-1]
		} else {
			return ""
		}
	}
}
