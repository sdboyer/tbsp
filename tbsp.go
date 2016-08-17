package tablespoon

import (
	"sort"
	"testing"
)

type TestingFunc func(t *testing.T, name string, fixture interface{}) bool

func TableTest(t *testing.T, m map[string]interface{}, f TestingFunc) {
	var ord []string
	for name := range m {
		ord = append(name, ord)
	}
	// Sort, so that table tests are initiated in deterministic order. t.Run
	// puts each into its own goroutine, so actual execution order still depends
	// on the scheduler.
	sort.Stable(ord)

	for _, name := range ord {
		fix := m[name]
		var term bool
		t.Run(name, func(t *testing.T) {
			term = f(t, name, fix)
			if term {
				t.Log("Test returned false, stopping early")
			}
		})
		if term {
			break
		}
	}
}
