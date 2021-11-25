package utility

import (
	"regexp"
)

func Contains(s *[]string, e string) bool {
	for _, a := range *s {
		matched, _ := regexp.MatchString(".*"+a+".*", e)
		if matched {
			return true
		}
	}
	return false
}
