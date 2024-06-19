package regexp

import (
	"regexp"
)

func ReplaceSpaceWithUnderscore(str string) string {
	// matches one or more characters of "\s", "/" and "."
	toBeReplaced, _ := regexp.Compile("[\\s/\\.]")

	// replace the matched characters / sub-strings
	match := toBeReplaced.ReplaceAllString(str, "_")

	return match
}

func RemoveAfterDot(str string) string {
	// matches characters after "."
	toBeRemoved, _ := regexp.Compile("[^.]*$")

	// replace the matched characters / sub-strings
	match := toBeRemoved.ReplaceAllString(str, "")

	return match
}

func RemoveIncludeDot(str string) string {
	// matches characters including ".*"
	toBeRemoved, _ := regexp.Compile("\\.\\w+")

	// replace the matched characters / sub-strings
	match := toBeRemoved.ReplaceAllString(str, "")

	return match
}
