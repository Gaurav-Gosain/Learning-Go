package testing_app

import (
	"fmt"
	"regexp"
)

func IsEmail(s string) (string, error) {
	re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if re.MatchString(s) {
		return s + " is a valid email 🔥", nil
	}
	return "", fmt.Errorf("%s is not a valid email 🚫", s)
}