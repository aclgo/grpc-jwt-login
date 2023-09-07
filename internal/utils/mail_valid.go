package utils

import "regexp"

func ValidMail(email string) bool {
	regexp := regexp.MustCompile(`^[a-zA-Z0-9._%-+]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return regexp.MatchString(email)
}
