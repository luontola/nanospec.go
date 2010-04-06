// Copyright © 2010 Esko Luontola <www.orfjackal.net>
// This software is released under the Apache License 2.0.
// The license text is at http://www.apache.org/licenses/LICENSE-2.0

package nanospec

import (
	"fmt"
	"os"
	"testing"
)


func Test__Expect_Equals(t *testing.T) {
	tt := TT(t)

	r := new(DummyReporter)
	newExpectation(200, r).Equals(200)
	tt.AssertEquals("", r.Message)

	r = new(DummyReporter)
	newExpectation(100, r).Equals(200)
	tt.AssertEquals("'100' should equal '200'", r.Message)
}

func Test__Expect_NotEquals(t *testing.T) {
	tt := TT(t)

	r := new(DummyReporter)
	newExpectation(100, r).NotEquals(200)
	tt.AssertEquals("", r.Message)

	r = new(DummyReporter)
	newExpectation(200, r).NotEquals(200)
	tt.AssertEquals("'200' should NOT equal '200'", r.Message)
}

func Test__Expect_IsTrue(t *testing.T) {
	tt := TT(t)

	r := new(DummyReporter)
	newExpectation(true, r).IsTrue()
	tt.AssertEquals("", r.Message)

	r = new(DummyReporter)
	newExpectation(false, r).IsTrue()
	tt.AssertEquals("'false' should be true", r.Message)
}

func Test__Expect_IsFalse(t *testing.T) {
	tt := TT(t)

	r := new(DummyReporter)
	newExpectation(false, r).IsFalse()
	tt.AssertEquals("", r.Message)

	r = new(DummyReporter)
	newExpectation(true, r).IsFalse()
	tt.AssertEquals("'true' should be false", r.Message)
}

func Test__Expect_Satisfies(t *testing.T) {
	tt := TT(t)

	actual := "foo"

	r := new(DummyReporter)
	newExpectation(actual, r).Satisfies(len(actual) == 3)
	tt.AssertEquals("", r.Message)

	r = new(DummyReporter)
	newExpectation(actual, r).Satisfies(len(actual) == 4)
	tt.AssertEquals("'foo' should satisfy the contract", r.Message)
}

func Test__Expect_Matches(t *testing.T) {
	tt := TT(t)

	actual := "foo"

	r := new(DummyReporter)
	newExpectation(actual, r).Matches(HasLength(3))
	tt.AssertEquals("", r.Message)

	r = new(DummyReporter)
	newExpectation(actual, r).Matches(HasLength(4))
	tt.AssertEquals("'foo' should have length 4", r.Message)
}

func HasLength(length int) Matcher {
	return func(actual interface{}) os.Error {
		if len(actual.(string)) != length {
			return os.ErrorString(fmt.Sprintf("'%v' should have length %v", actual, length))
		}
		return nil
	}
}
