package regexp

import (
	"log"
	"regexp"
	"strings"
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

func GetFileName(fileUrl string, keyword string) string {
	// split the string into list
	stringList := regexp.MustCompile("[/?]").Split(fileUrl, -1)
	log.Printf("The extracted string list is %v", stringList)

	// match for sub-string which contains the keyword
	fileName := ""
	for _, str := range stringList {
		if strings.Contains(str, keyword) {
			fileName = str
		}
	}

	log.Printf("The matching filename is %v", fileName)

	return fileName
}
