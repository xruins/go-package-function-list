package function

import (
	"fmt"
	"regexp"
	"strings"
)

// FilterByRegexp filters the functions matches regexp from src
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

// FilterByPrefix filters the functions has given prefix from src
func FilterByPrefix(src []string, prefix string) []string {
	var ret []string
	for _, s := range src {
		if strings.HasPrefix(s, prefix) {
			ret = append(ret, s)
		}
	}
	return ret
}

// FilterBySuffix filters the functions has given suffix from src
func FilterBySuffix(src []string, suffix string) []string {
	var ret []string
	for _, s := range src {
		if strings.HasSuffix(s, suffix) {
			ret = append(ret, s)
		}
	}
	return ret
}

// FilterPublicMethod filters public methods from src
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
