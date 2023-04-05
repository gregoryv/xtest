package xtest

import (
	"flag"
	"strings"
	"testing"
)

var ts = flag.String("test-suite", "", "")

func SkipUnless(t testing.TB, suites ...string) {
	t.Helper()
	if *ts == "all" {
		return // run all
	}
	selected := make(map[string]bool)
	for _, k := range strings.Split(*ts, ",") {
		selected[k] = true
	}
	for _, s := range suites {
		if selected[s] {
			return // means run this test
		}
	}
	t.Skipf("include with --test-suite=%v", suites)

}
