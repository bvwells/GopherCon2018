# Testing 

This document contains a few tips and tricks for tests in Go.

# Running sub-tests.

This is example shows how to run sub-tests in table driven tests.

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

# Black-box testing 

Black-box testing can be used to test only public functions exposed from a package. It can also be used to avoid circular dependencies in tests.

```
// foo_test.go
package foo_test
```

# Verbosity of tests

It is possible to add extra test output depending on whether the test is run with the verbose option or not. It is also possible to silence logging.

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

It is possible to run tests in parallel by adding ```t.Parallel()```. For example:

```
func Test_Parallel(t *testing.T) {
	t.Parallel()
	t.Log("Running parallel test...")
}
```

# Race detector
 
 The race detector is a useful tool for finding race conditions. NEVER QUESTION THE RACE DETECTOR!

 Tests can be ran with the race detector using the option ```-race```.

 ```
 go test -v -race ./... 
 ```

Race detector can also be used when running applications. 

# Example tests

Tests can be documented in the following manner:

 ```
// This is some text from your example.
// You can add additional information about your example here.
func ExampleDuplicate() {
	d := src.Duplicate("foo")
	fmt.Println(d)
}
```