package function

import (
	"fmt"
	"regexp"
	"strings"
)

// FilterBySuffix filters string matches regexp from given string slice.
func FilterByRegexp(src []string, regex string) ([]string, error) {
	r, err := regexp.Compile(regex)
	if err != nil {
		return nil, fmt.Errorf("failed to compile regexp. err: %s", err)
	}

	var ret []string
	for _, s := range src {
		if r.Match([]byte(s)) {
			ret = append(ret, s)
		}
	}
	return ret, nil
}

// FilterBySuffix filters string has suffix from given string slice.
func FilterBySuffix(src []string, suffix string) []string {
	var ret []string
	for _, s := range src {
		if strings.HasSuffix(s, suffix) {
			ret = append(ret, s)
		}
	}
	return ret
}

// FilterPublicMethod filters string starts with capital letter from given string slice.
func FilterPublicMethod(src []string) []string {
	var ret []string
	for _, s := range src {
		firstByte := byte(s[0])
		// 0x41: A, 0x5a: Z
		if 0x41 <= firstByte && firstByte <= 0x5a {
			ret = append(ret, s)
		}
	}
	return ret
}
