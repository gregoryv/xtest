//go:build systest

package xtest

import "testing"

func TestB(t *testing.T) {
	// .. test here
	t.Fail() // just for demo purposes
}
