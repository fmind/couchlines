package main

import "testing"

var trimTests = []struct {
	in  string
	out string
}{
	// first line
	{`{"total_rows":3897205,"rows":[` + "\n",
		``},
	// second line
	{`{"id":"0","key":"0","value":{"rev":"0"},"doc":{"_id": "a", "a": 1}},"`,
		`{"_id": "a", "a": 1}`},
	// ... line
	{`{"id":"0","key":"0","value":{"rev":"0"},"doc":{"_id": "b", "b": 2}},"`,
		`{"_id": "b", "b": 2}`},
	// last line
	{`{"id":"0","key":"0","value":{"rev":"0"},"doc":{"_id": "c", "c": 3}}"`,
		`{"_id": "c", "c": 3}`},
}

func TestTrim(t *testing.T) {
	for _, tt := range trimTests {
		if got := Trim(tt.in + "\n"); got != tt.out {
			t.Errorf("Expected %v from %v, got %s", tt.out, tt.in, got)
		}
	}
}
