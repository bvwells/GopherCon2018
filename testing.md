# Testing

sub-tests.

```
func TestSub(t *testing.T) {
	tt := []struct {
		A        int
		B        int
		Expected int
	}{
		{A: 1, B: 1, Expected: 2},
		{A: 2, B: 2, Expected: 4},
		{A: 3, B: 3, Expected: 5},
		{A: 4, B: 4, Expected: 6},
	}

	globalSetup()          // run necessary setup for all the tests
	defer globalTeardown() // make sure tearing down of all the tests

	for i, x := range tt {
		t.Run(fmt.Sprintf("sub test (%d)", i), func(st *testing.T) {
			setup()          // run necessary setup for the test
			defer teardown() // make sure teardown is always called

			got := x.A + x.B
			if got != x.Expected {
				st.Errorf("expected %d, got %d", x.Expected, got)
			}
		})
	}
}
```

# Testing (black-box testing) - avoid circular dependencies.
Only access public func.
```
// foo_test.go
package foo_test
```

- get rid of mock messages?

```
func TestVerbose(t *testing.T) {
	if testing.Verbose() {
		t.Log("put extra logging here that normally we don't care about")
	} else {
		// silence my normal loggers
		log.SetOutput(ioutil.Discard)
	}
}
```

# Test in parallel 
```
func Test_Parallel(t *testing.T) {
	t.Parallel()
	t.Log("Running parallel test...")
}
```

# Race detector
 -race!

# Example tests
 ```
 // This is some text from your example.
// You can add additional information about your example here.
func ExampleDuplicate() {
	d := src.Duplicate("foo")
	fmt.Println(d)
}
```