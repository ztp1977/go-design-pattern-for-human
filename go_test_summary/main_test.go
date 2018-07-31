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
