# tablespoon

A simple approach to executing table-based tests using [Go 1.7's new subtests](https://golang.org/pkg/testing/#hdr-Subtests_and_Sub_benchmarks).

Basically, the idea is:

1. Define a type `T` that implements the [one-method `Fixture` interface](https://godoc.org/github.com/sdboyer/tablespoon#Fixture), doing all the input/output verification in that method
1. Declare your test tables as a `map[string]Fixture`, where the key is a descriptive name for what that particular row is testing
3. Pass the `t *testing.T`, the `map[string]Fixture`, and the func to `tbsp.RunTableTest()`
4. PROFIT :tada: :tada: :tada: :tada: :sparkles: :tada: :tada: :tada: :sparkles: :sparkles: :tada:

_Note: this is so trivial to do that it's probably a terrible idea as a **library**, but it demonstrates how TBTs should, now, probably be run to take advantage of subtests._

