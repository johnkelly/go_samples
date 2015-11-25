package string

import "testing"

func TestReverse(t *testing.T) {
	var tests = []struct {
		s, want string
	}{
		{"Backward", "drawkcaB"},
		{"racecar", "racecar"},
		{"Ô¨∏", "∏¨Ô"},
		{"", ""},
	}
	for _, c := range tests {
		got := Reverse(c.s)
		if got != c.want {
			t.Errorf("Reverse(%q) == %q, want %q", c.s, got, c.want)
		}
	}
}

func TestVowelsOnly(t *testing.T) {
	var tests = []struct {
		s, want string
	}{
		{"Backward", "aa"},
		{"hmmmq", ""},
		{"", ""},
		{"Ô¨∏", ""},
	}
	for _, c := range tests {
		got := VowelsOnly(c.s)
		if got != c.want {
			t.Errorf("VowelsOnly(%q) == %q, want %q", c.s, got, c.want)
		}
	}
}
