// Copyright © 2010 Esko Luontola <www.orfjackal.net>
// This software is released under the Apache License 2.0.
// The license text is at http://www.apache.org/licenses/LICENSE-2.0

package nanospec

import (
	"errors"
	"fmt"
	"testing"
)

func Test__Expect_Equals(t *testing.T) {
	tt := TT(t)

	r := new(DummyReporter)
	newExpectation(200, r).Equals(200)
	tt.AssertEquals("", r.Message)

	r = new(DummyReporter)
	newExpectation(100, r).Equals(200)
	tt.AssertEquals("Expected: equals '200'\n\tgot: '100'", r.Message)
}

func Test__Expect_NotEquals(t *testing.T) {
	tt := TT(t)

	r := new(DummyReporter)
	newExpectation(100, r).NotEquals(200)
	tt.AssertEquals("", r.Message)

	r = new(DummyReporter)
	newExpectation(200, r).NotEquals(200)
	tt.AssertEquals("Expected: NOT equals '200'\n\tgot: '200'", r.Message)
}

func Test__Expect_IsTrue(t *testing.T) {
	tt := TT(t)

	r := new(DummyReporter)
	newExpectation(true, r).IsTrue()
	tt.AssertEquals("", r.Message)

	r = new(DummyReporter)
	newExpectation(false, r).IsTrue()
	tt.AssertEquals("Expected: is true\n\tgot: 'false'", r.Message)
}

func Test__Expect_IsFalse(t *testing.T) {
	tt := TT(t)

	r := new(DummyReporter)
	newExpectation(false, r).IsFalse()
	tt.AssertEquals("", r.Message)

	r = new(DummyReporter)
	newExpectation(true, r).IsFalse()
	tt.AssertEquals("Expected: is false\n\tgot: 'true'", r.Message)
}

func Test__Expect_Satisfies(t *testing.T) {
	tt := TT(t)

	actual := "foo"

	r := new(DummyReporter)
	newExpectation(actual, r).Satisfies(len(actual) == 3)
	tt.AssertEquals("", r.Message)

	r = new(DummyReporter)
	newExpectation(actual, r).Satisfies(len(actual) == 4)
	tt.AssertEquals("Expected: satisfies the contract\n\tgot: 'foo'", r.Message)
}

func Test__Expect_Matches(t *testing.T) {
	tt := TT(t)

	actual := "foo"

	r := new(DummyReporter)
	newExpectation(actual, r).Matches(HasLength(3))
	tt.AssertEquals("", r.Message)

	r = new(DummyReporter)
	newExpectation(actual, r).Matches(HasLength(4))
	tt.AssertEquals("Expected: has length 4\n\tgot: 'foo'", r.Message)
}

func HasLength(length int) Matcher {
	return func(actual interface{}) error {
		if len(actual.(string)) != length {
			return errors.New(fmt.Sprintf("Expected: has length %v\n\tgot: '%v'", length, actual))
		}
		return nil
	}
}
