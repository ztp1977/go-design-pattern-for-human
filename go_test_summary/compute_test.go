package go_test_summary

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"testing"
	"testing/quick"
)

func TestComputer(t *testing.T) {
	s, err := Compute("1+1")
	if err != nil {
		t.Fatal(err)
	}
	if s != "2" {
		t.Errorf("Compute(1+1) = %s, want 2", s)
	}
}

func TestComputerForTable(t *testing.T) {
	computerTests := []struct {
		in  string
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

func TestComputeWithSetup(t *testing.T) {
	r, Teardown := SetupComputeTest(t, "testdata/compute.txt")
	defer Teardown()

	for {
		line, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}
		test := strings.Split(string(line), "=")
		if len(test) != 2 {
			t.Fatalf("invalid test data : %s", string(line))
		}
		testCompute(t, test[0], test[1])
	}
}

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

func BenchmarkCompute(b *testing.B) {
	for i := 0; i < b.N; i++ {
		quick.Check(func(a, b int16) {
			Compute(fmt.Sprintf("%d+%d", a, b))
		}, nil)
	}
}
