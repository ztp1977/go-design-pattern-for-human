# Go Test Summary

### Setup/Teardown

+ main_test.go

```go
package go_test_summary

import (
	"testing"
	"os"
	"fmt"
)

func setup() {
	fmt.Println("on setup")
}

func teardown() {
	fmt.Println("on tear down")
}

func TestMain(m *testing.M) {

	setup()

	ret := m.Run()

	teardown()
	os.Exit(ret)
}
```

+ function setup

```go
// Setup
func SetupComputeTest(t *testing.T, fname string) (*bufio.Reader, func()) {
	f, err := os.Open(fname)
	if err != nil {
		t.Fatal(err)
	}
	return bufio.NewReader(f), func() {
		f.Close()
	}
}
// Usage
{
	r, Teardown := SetupComputeTest(t, "testdata/compute.txt")
	defer Teardown()
	...
}
```

### Table

+ shift+cmd+t

```go
func TestComputerForTable(t *testing.T) {
	computerTests := []struct{
		in string
		out string
	}{
		{"1+1", "2"},
		{"1.0/2.0", "0.5"},
	}

	for _, test := range computerTests {
		s, err := Compute(test.in)
		if err != nil {
			t.Fatal(err)
		}
		if s != test.out {
			t.Errorf("Compute(%s) = %s, want %s", test.in, s, test.out)
		}
	}
}
```

### SubClass - t.Run执行

```go
// Sub Test, Callback for t.Run(t, func)
func testCompute(t *testing.T, in, expected string) {
	// t.Helper
	t.Helper()
	s, err := Compute(in)
	if err != nil {
		t.Fatal(err)
	}
	if s != expected {
		t.Errorf("Compute(%s) = %s, want %s", in, s, expected)
	}
}

func TestComputeByRun(t *testing.T) {
	t.Run("add sub", func(t *testing.T) {
		testCompute(t, "1+1", "2")
		testCompute(t, "-1+1", "0")
		testCompute(t, "-2+1", "-1")
	})
	t.Run("div", func(t *testing.T) {
		testCompute(t, "1.0/2.0", "0.5")
		testCompute(t, "2.0/1.0", "2")
	})
}
```

### File

```go
// Read Config File
func TestComputeByFile(t *testing.T) {
	f, err := os.Open("testdata/compute.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		line, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}
		test := strings.Split(string(line), "=")
		if len(test) != 2 {
			t.Fatalf("invalid test data: %s", string(line))
		}
		testCompute(t, test[0], test[1])
	}
}
// cat testdata/compute.txt
// 1+1=2
// 2+2=4
// 5*2=10
```

### Quick/Random

```go


func TestComputeForRandom(t *testing.T) {
	add := func(a, b int16) bool {
		s, err := Compute(fmt.Sprintf("%d+%d", a, b))
		if err != nil {
			t.Fatal(err)
		}
		expected := strconv.Itoa(int(a) + int(b))
		//pp.Println(a, b, expected)
		if s != expected {
			t.Logf("Compute(%d+%d) = %s, want %s", a, b, s, expected)
			return false
		}
		return true
	}

	if err := quick.Check(add, nil); err != nil {
		t.Fatal(err)
	}
}
```

### Struct

```go

type computeTest struct {
	testing.TB
	f *os.File
	r *bufio.Reader
}

func SetupComputeTestWithStruct(tb testing.TB, fname string) *computeTest {
	f, err := os.Open(fname)
	if err != nil {
		tb.Fatal(err)
	}

	return &computeTest{
		TB: tb,
		f:  f,
		r:  bufio.NewReader(f),
	}
}

func (t *computeTest) Teardown() {
	t.f.Close()
}

func (t *computeTest) testData() (in, out string, ok bool) {
	line, _, err := t.r.ReadLine()
	if err == io.EOF {
		return "", "", false
	}
	test := strings.Split(string(line), "=")
	if len(test) != 2 {
		t.Fatalf("invalid test data %s", string(line))
	}
	return test[0], test[1], true
}

func TestComputeForStruct(tt *testing.T) {
	t := SetupComputeTestWithStruct(tt, "testdata/compute.txt")
	defer t.Teardown()

	for {
		in, out, ok := t.testData()
		if !ok {
			break
		}
		testCompute(tt, in, out)
	}
}
```

### Example

```go
func ExampleHello() {
        fmt.Println("hello")
        // Output: hello
}
```


### Testing-Doc

```go
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package testing provides support for automated testing of Go packages.
// It is intended to be used in concert with the ``go test'' command, which automates
// execution of any function of the form
//     func TestXxx(*testing.T)
// where Xxx does not start with a lowercase letter. The function name
// serves to identify the test routine.
//
// Within these functions, use the Error, Fail or related methods to signal failure.
//
// To write a new test suite, create a file whose name ends _test.go that
// contains the TestXxx functions as described here. Put the file in the same
// package as the one being tested. The file will be excluded from regular
// package builds but will be included when the ``go test'' command is run.
// For more detail, run ``go help test'' and ``go help testflag''.
//
// Tests and benchmarks may be skipped if not applicable with a call to
// the Skip method of *T and *B:
     func TestTimeConsuming(t *testing.T) {
         if testing.Short() {
             t.Skip("skipping test in short mode.")
         }
        //  ...
     }
//
// Benchmarks
//
// Functions of the form
//     func BenchmarkXxx(*testing.B)
// are considered benchmarks, and are executed by the "go test" command when
// its -bench flag is provided. Benchmarks are run sequentially.
//
// For a description of the testing flags, see
// https://golang.org/cmd/go/#hdr-Description_of_testing_flags.
//
// A sample benchmark function looks like this:
     func BenchmarkHello(b *testing.B) {
         for i := 0; i < b.N; i++ {
             fmt.Sprintf("hello")
         }
     }
//
// The benchmark function must run the target code b.N times.
// During benchmark execution, b.N is adjusted until the benchmark function lasts
// long enough to be timed reliably. The output
//     BenchmarkHello    10000000    282 ns/op
// means that the loop ran 10000000 times at a speed of 282 ns per loop.
//
// If a benchmark needs some expensive setup before running, the timer
// may be reset:
//
     func BenchmarkBigLen(b *testing.B) {
         big := NewBig()
         b.ResetTimer()
         for i := 0; i < b.N; i++ {
             big.Len()
         }
     }
//
// If a benchmark needs to test performance in a parallel setting, it may use
// the RunParallel helper function; such benchmarks are intended to be used with
// the go test -cpu flag:
//
     func BenchmarkTemplateParallel(b *testing.B) {
         templ := template.Must(template.New("test").Parse("Hello, {{.}}!"))
         b.RunParallel(func(pb *testing.PB) {
             var buf bytes.Buffer
             for pb.Next() {
                 buf.Reset()
                 templ.Execute(&buf, "World")
             }
         })
     }
//
// Examples
//
// The package also runs and verifies example code. Example functions may
// include a concluding line comment that begins with "Output:" and is compared with
// the standard output of the function when the tests are run. (The comparison
// ignores leading and trailing space.) These are examples of an example:
//
     func ExampleHello() {
         fmt.Println("hello")
         // Output: hello
     }

     func ExampleSalutations() {
         fmt.Println("hello, and")
         fmt.Println("goodbye")
         // Output:
         // hello, and
         // goodbye
     }
//
// The comment prefix "Unordered output:" is like "Output:", but matches any
// line order:
//
     func ExamplePerm() {
         for _, value := range Perm(4) {
             fmt.Println(value)
         }
         // Unordered output: 4
         // 2
         // 1
         // 3
         // 0
     }
//
// Example functions without output comments are compiled but not executed.
//
// The naming convention to declare examples for the package, a function F, a type T and
// method M on type T are:
//
//     func Example() { ... }
//     func ExampleF() { ... }
//     func ExampleT() { ... }
//     func ExampleT_M() { ... }
//
// Multiple example functions for a package/type/function/method may be provided by
// appending a distinct suffix to the name. The suffix must start with a
// lower-case letter.
//
//     func Example_suffix() { ... }
//     func ExampleF_suffix() { ... }
//     func ExampleT_suffix() { ... }
//     func ExampleT_M_suffix() { ... }
//
// The entire test file is presented as the example when it contains a single
// example function, at least one other function, type, variable, or constant
// declaration, and no test or benchmark functions.
//
// Subtests and Sub-benchmarks
//
// The Run methods of T and B allow defining subtests and sub-benchmarks,
// without having to define separate functions for each. This enables uses
// like table-driven benchmarks and creating hierarchical tests.
// It also provides a way to share common setup and tear-down code:
//
     func TestFoo(t *testing.T) {
         // <setup code>
         t.Run("A=1", func(t *testing.T) { ... })
         t.Run("A=2", func(t *testing.T) { ... })
         t.Run("B=1", func(t *testing.T) { ... })
         // <tear-down code>
     }
//
// Each subtest and sub-benchmark has a unique name: the combination of the name
// of the top-level test and the sequence of names passed to Run, separated by
// slashes, with an optional trailing sequence number for disambiguation.
//
// The argument to the -run and -bench command-line flags is an unanchored regular
// expression that matches the test's name. For tests with multiple slash-separated
// elements, such as subtests, the argument is itself slash-separated, with
// expressions matching each name element in turn. Because it is unanchored, an
// empty expression matches any string.
// For example, using "matching" to mean "whose name contains":
//
//     go test -run ''      # Run all tests.
//     go test -run Foo     # Run top-level tests matching "Foo", such as "TestFooBar".
//     go test -run Foo/A=  # For top-level tests matching "Foo", run subtests matching "A=".
//     go test -run /A=1    # For all top-level tests, run subtests matching "A=1".
//
// Subtests can also be used to control parallelism. A parent test will only
// complete once all of its subtests complete. In this example, all tests are
// run in parallel with each other, and only with each other, regardless of
// other top-level tests that may be defined:
//
     func TestGroupedParallel(t *testing.T) {
         for _, tc := range tests {
             tc := tc // capture range variable
             t.Run(tc.Name, func(t *testing.T) {
                 t.Parallel()
                 // ...
             })
         }
     }
//
// Run does not return until parallel subtests have completed, providing a way
// to clean up after a group of parallel tests:
//
     func TestTeardownParallel(t *testing.T) {
         // This Run will not return until the parallel tests finish.
         t.Run("group", func(t *testing.T) {
             t.Run("Test1", parallelTest1)
             t.Run("Test2", parallelTest2)
             t.Run("Test3", parallelTest3)
         })
         // <tear-down code>
     }
//
// Main
//
// It is sometimes necessary for a test program to do extra setup or teardown
// before or after testing. It is also sometimes necessary for a test to control
// which code runs on the main thread. To support these and other cases,
// if a test file contains a function:
//
	func TestMain(m *testing.M)
//
// then the generated test will call TestMain(m) instead of running the tests
// directly. TestMain runs in the main goroutine and can do whatever setup
// and teardown is necessary around a call to m.Run. It should then call
// os.Exit with the result of m.Run. When TestMain is called, flag.Parse has
// not been run. If TestMain depends on command-line flags, including those
// of the testing package, it should call flag.Parse explicitly.
//
// A simple implementation of TestMain is:
//
	func TestMain(m *testing.M) {
		// call flag.Parse() here if TestMain uses flags
		os.Exit(m.Run())
	}

```


