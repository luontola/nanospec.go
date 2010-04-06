// Copyright © 2010 Esko Luontola <www.orfjackal.net>
// This software is released under the Apache License 2.0.
// The license text is at http://www.apache.org/licenses/LICENSE-2.0

package nanospec

import (
	"testing"
)


func Test__Get_location_of_calling_method(t *testing.T) {
	tt := TT(t)

	location := fakeExpectationMethod() // line 15

	tt.AssertEquals("location_test.go:15", location)
}

func fakeExpectationMethod() string {
	return callerLocation()
}
