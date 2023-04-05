package xtest

import "testing"

func TestAlways(t *testing.T) {
	// ... test here
}

func TestA(t *testing.T) {
	SkipUnless(t, "blue") // skips if not in the given test suites
	// ... test here
}
