package regex

import "regexp"

// Regex ...
type Regex struct {
	PhoneRegexp *regexp.Regexp
}

// NewRegex is
func NewRegex() *Regex {

	r, _ := regexp.Compile("^\\+?[1-9]\\d{1,14}$")

	return &Regex{
		PhoneRegexp: r,
	}
}

// PhoneE164 ...
func (t *Regex) PhoneE164(no string) bool {
	return t.PhoneRegexp.MatchString(no)
}
