package util

import "regexp"

func IsSupportedPhone(phone string) bool {
	digitPattern := regexp.MustCompile(`^\d{11}$`)
	return digitPattern.MatchString(phone)
}
