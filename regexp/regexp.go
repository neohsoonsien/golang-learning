package regexp

import (
	"regexp"
)

func ReplaceSpaceWithUnderscore(str string) string {
	// matches one or more characters "\s" or "/"
	toBeReplaced, _ := regexp.Compile("[\\s/]")

	// replace the matched characters / sub-strings
	match := toBeReplaced.ReplaceAllString(str, "_")

	return match
}

func RemoveAfterDot(str string) string {
	// matches one or more characters "\s" or "/"
	toBeReplaced, _ := regexp.Compile("[^.]*$")

	// replace the matched characters / sub-strings
	match := toBeReplaced.ReplaceAllString(str, "")

	return match
}
