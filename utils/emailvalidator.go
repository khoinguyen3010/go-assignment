package utils

import (
	"net"
	"regexp"
	"strings"

	"github.com/badoux/checkmail"
)

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func ValidateEmailRegex(email string) bool {
	if len(email) < 3 && len(email) > 254 {
		return false
	}
	if !emailRegex.MatchString(email) {
		return false
	}
	parts := strings.Split(email, "@")
	host, err := net.LookupMX(parts[1])
	if err != nil || len(host) == 0 {
		return false
	}
	return true
}

func CheckEmail(email string) bool {
	if err := checkmail.ValidateFormat(email); err != nil {
		return false
	}
	if err := checkmail.ValidateHost(email); err != nil {
		return false
	}
	return true
}
