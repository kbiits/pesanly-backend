package utils

import (
	"regexp"
	"strings"
)

var RegexPrefixPhoneNumber = regexp.MustCompile("^\\+?(62)?")

func CleanPhoneNumber(p string) string {
	return "0" + strings.TrimLeft(RegexPrefixPhoneNumber.ReplaceAllString(p, ""), "0")
}
