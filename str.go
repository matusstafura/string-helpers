package str

import (
	"regexp"
	"strings"
)

func After(s, c string) string {
	r := strings.Split(s, c)
	return string(r[1])
}

func Append(s ...string) string {
	return strings.Join(s, "")
}

func Before(s, c string) string {
	r := strings.Split(s, c)
	return string(r[0])
}

func EndsWith(s string, c string) bool {
	r := len(s) - len(c)
	if string(s[r:]) == c {
		return true
	} 
	
	return false
}

func Mask(s string, n int, r string) string {
	l := len(s)
	x := s[n:l]
	re := regexp.MustCompile(`.`)

	before := Limit(s, n)
	after := re.ReplaceAllString(x, r)

	return (before + after)
}

func Length(s string) int {
	return len(s)
}

func Limit(s string, n int) string {
	return s[0:n]
}

func Lower(s string) string {
	return strings.ToLower(s)
}

func Repeat(s string, n int) string {
	x := ""
	i := 1
	for i < 5 {
		x += s
		i++
	}
	return x
}

func Reverse(s string) string {
	f := make([]string, len(s))

	for i, c := range s {
		f[len(s)-1-i] = string(c)
	}
	return strings.Join(f, "")
}

func StartsWith(s string, c string) bool {
	if string(s[0:len(c)]) == c {
		return true
	} 
	
	return false
}

func Title(s string) string {
	return strings.Title(Lower(s))
}

func Upper(s string) string {
	return strings.ToUpper(s)
}
