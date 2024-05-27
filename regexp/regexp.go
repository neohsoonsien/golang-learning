package regexp

import (
	"regexp"
)

func Replace(str string) string {
	// matches one or more characters "\s" or "/"
	toBeReplaced, _ := regexp.Compile("[\\s/]")

	// replace the matched characters / sub-strings
	match := toBeReplaced.ReplaceAllString(str, "_")

	return match
}
