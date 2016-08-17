package tablespoon

import (
	"sort"
	"testing"
)

// Fixture represents a single row in a table-based test. It is called with the
// descriptive name (the key associated with the fixture provided in RunTableTest)
// and the testing apparatus.
//
// Returning true from Test will cause the table test run to abort early.
type Fixture interface {
	Test(t *testing.T, name string) (stopearly bool)
}

func RunTableTest(t *testing.T, m map[string]Fixture) {
	var ord []string
	for name := range m {
		ord = append(ord, name)
	}

	// Sort, so that table tests are initiated in deterministic order. Note that
	// t.Run puts each into its own goroutine, so actual execution order still
	// depends on the scheduler.
	sort.Strings(ord)

	for _, name := range ord {
		fix := m[name]
		var term bool
		t.Run(name, func(t *testing.T) {
			term = fix.Test(t, name)
			if term && testing.Verbose() {
				t.Log("Fixture.Test() returned true, stopping early")
			}
		})
		if term {
			break
		}
	}
}
