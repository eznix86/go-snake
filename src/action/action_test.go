package action

import (
	"testing"
)

var actiontests = []struct {
	in  string
	out Direction
}{
	{"w", UP},
	{"a", LEFT},
	{"s", DOWN},
	{"d", RIGHT},
	{" ", NONE},
	{"dd", NONE},
	{"21312", NONE},
	{"k o", NONE},
	{"##", NONE},
	{"^[[A", NONE},
	{"%-+1.2abc", NONE},
	{"%-1.2abc", NONE},
}
func DirectionPrompter(t *testing.T) {
	for _, tt := range actiontests {
		t.Run(tt.in, func(t *testing.T) {
			s := GetDirection(tt.in)
			if s != tt.out {
				t.Errorf("got %q, want %q", s, tt.out)
			}
		})
	}
}