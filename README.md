# tablespoon

A simple approach to converting table-based tests to [Go 1.7's new subtests](https://golang.org/pkg/testing/#hdr-Subtests_and_Sub_benchmarks).

_Note: this is a terrible idea as a **library** (because no generics), but it demonstrates how TBTs should, now, probably be run._

Basically, the idea is:

1. Declare your testing fixture tables as a `map[string]T`, where the key is a descriptive name and `T` is the `type`` for the fixture
2. Wrap your input, output, and verification logic up in a standalone func or closure with the signature `func (t *testing.T, fixture T)`
3. Pass the `*testing.T`, the map, and the func to the table test runner
4. :tada: :tada: :tada: :tada: :sparkles: :tada: :tada: :tada: :sparkles: :sparkles: :tada:

Hell, I'm sure we could write a generator pretty easily.
