package regex

import (
	"fmt"
	"testing"
)

func TestRegex(t *testing.T) {

	r := NewRegex()
	fmt.Println("234567890", r.PhoneE164("234567890"))
	fmt.Println("+6282112345678", r.PhoneE164("+6282112345678"))
}
