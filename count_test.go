package count

import "testing"

func TestCount(t *testing.T) {
	s := "qwertyeeud"
	e := 3
	if c := Count(s, 'e'); c != e {
		t.Fatalf("bad  count for %s: got %d expected %d", s, c, e)
	}
}
