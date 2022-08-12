package morestrings

import "testing"

func TestReverseRunes(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{"Hello, 世界", "界世 ,olleH"},
		{"", ""},
	}

	for _, c := range cases {
		result := ReverseRunes(c.in)
		if result != c.want {
			t.Errorf("ReverseRunes(%q) == %q, want %q", c.in, result, c.want)
		}
	}
}
