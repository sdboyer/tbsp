package tbsp

import "testing"

func sq(x int) int {
	return x * x
}

type sqfix struct {
	in, out int
}

func (fix sqfix) Test(t *testing.T, name string) bool {
	got := sq(fix.in)
	if got != fix.out {
		t.Errorf("Expected %v when squaring %v, got %v", fix.out, fix.out, got)
	}
	return false
}

func TestBasicTabling(t *testing.T) {
	table := map[string]Fixture{
		"x=1": sqfix{in: 1, out: 1},
		"x=2": sqfix{in: 2, out: 4},
		"x=3": sqfix{in: 3, out: 9},
		"x=4": sqfix{in: 4, out: 16},
		"x=5": sqfix{in: 5, out: 25},
	}

	RunTableTest(t, table)
}
