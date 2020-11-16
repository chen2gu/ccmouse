package main

import "testing"

func TestSubstr(t *testing.T) {
	tests := []struct {
		s   string
		ans int
	}{
		{"abcabcaaacbb", 3},
		{"abcabffcaaacbb", 4},
		{"abasdcabcaaacbb", 5},
		{"asbcabcaaacbb", 4},
		{"", 0},
		{"bb", 1},
		{"abcd", 4},
	}

	for _, tt := range tests {
		actual := length0fNonRepeatingSubStr(tt.s)
		if actual != tt.ans {
			t.Errorf("calcTriangle %d ; got %s; expected %d.", actual, tt.s, tt.ans)
		}
	}

}
