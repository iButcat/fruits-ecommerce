package utils

import "regexp"

func RegexEmailChecker(email string) bool {
	var validator = `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`
	var regexEmail = regexp.MustCompile(validator)
	return regexEmail.MatchString(email)
}

func ValidateIfNotEmpty(fields ...int) bool {
	for _, field := range fields {
		if field == 0 {
			return false
		}
	}
	return true
}
