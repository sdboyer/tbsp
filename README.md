# tablespoon

A simple approach to executing table-based tests using [Go 1.7's new subtests](https://golang.org/pkg/testing/#hdr-Subtests_and_Sub_benchmarks).

Basically, the idea is:

1. Define a type `T` that implements the [one-method `Fixture` interface](https://godoc.org/github.com/sdboyer/tbsp#Fixture), doing all the input/output verification in that method
1. Declare your test tables as a `map[string]Fixture`, where the key is a descriptive name for what that particular row is testing
3. Pass the `t *testing.T`, the `map[string]Fixture`, and the func to `tbsp.RunTableTest()`
4. PROFIT :tada: :tada: :tada: :tada: :sparkles: :tada: :tada: :tada: :sparkles: :sparkles: :tada:

So, it'll look something like this:

```go
// The func we're testing
func Sq(x int) int {
	return x * x
}

// The data fixture we use for the test - a "row" in the "table"
type sqfix struct {
	in, out int
}

// The method to actually do the testing
func (fix sqfix) Test(t *testing.T, name string) bool {
	got := Sq(fix.in)
	if got != fix.out {
		t.Errorf("Expected %v when squaring %v, got %v", fix.out, fix.out, got)
	}
	return false
}

func TestSq(t *testing.T) {
	table := map[string]tbsp.Fixture{
		"x=1": sqfix{in: 1, out: 1},
		"x=2": sqfix{in: 2, out: 4},
		"x=3": sqfix{in: 3, out: 9},
		"x=4": sqfix{in: 4, out: 16},
		"x=5": sqfix{in: 5, out: 25},
	}

	tbsp.RunTableTest(t, table)
}
```

_Note: this is [so trivial](https://github.com/sdboyer/tbsp/blob/master/tbsp.go) to do that it's probably a terrible idea as a **library**, but it demonstrates how TBTs should, now, probably be run to take advantage of subtests._

