xtest - show how to create test suites in Go

This example shows how to group tests.
Use build tags or by supplying a suite name.
The latter has the benefit of always compiling your tests.


## Quick start

    $ git clone git@github.com:gregoryv/xtest.git
    $ cd xtest

In this demo there are three tests spread out over two files

1. a_test.go TestAlways and TestA
2. b_test.go TestB

TestAlways is not part of any test suite nor is the file a_test.go tagged in any way
TestA however has is skipped unless you provide a the blue test suite using --test-suit=blue
TestB is in a file tagged with `//go:build systest`

    $ go test -v
    === RUN   TestAlways
    --- PASS: TestAlways (0.00s)
    === RUN   TestA
        a_test.go:10: include with --test-suite=[blue]
    --- SKIP: TestA (0.00s)
    PASS
    ok      github.com/gregoryv/xtest       0.001s


You can see that TestA is skipped, but included in the compiled set of
tests, TestB however is not compiled at all.


    $ go test --tags=systest -v
    === RUN   TestAlways
    --- PASS: TestAlways (0.00s)
    === RUN   TestA
        a_test.go:10: include with --test-suite=[blue]
    --- SKIP: TestA (0.00s)
    === RUN   TestB
    --- FAIL: TestB (0.00s)
    FAIL
    exit status 1
    FAIL    github.com/gregoryv/xtest       0.001s


Using --tags=systest TestB is compiled And executed where as TestA is
still just compiled but not included.


    $ go test --tags=systest --test-suite=red,blue,green -v
    === RUN   TestAlways
    --- PASS: TestAlways (0.00s)
    === RUN   TestA
    --- PASS: TestA (0.00s)
    === RUN   TestB
    --- FAIL: TestB (0.00s)
    FAIL
    exit status 1
    FAIL    github.com/gregoryv/xtest       0.001s


