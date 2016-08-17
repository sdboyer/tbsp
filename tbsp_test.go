package tablespoon

import "testing"

func sq(x int) int {
	return x * x
}

type dumbtable struct {
	in, out int
}

func (fix dumbtable) Test(t *testing.T, name string) bool {
	got := sq(fix.in)
	if got != fix.out {
		t.Errorf("Expected %v when squaring %v, got %v", fix.out, fix.out, got)
	}
	return false
}

func TestBasicTabling(t *testing.T) {

	table := map[string]Fixture{
		"x=1": dumbtable{in: 1, out: 1},
		"x=2": dumbtable{in: 2, out: 4},
		"x=3": dumbtable{in: 3, out: 9},
		"x=4": dumbtable{in: 4, out: 16},
		"x=5": dumbtable{in: 5, out: 25},
	}

	RunTableTest(t, table)
}
