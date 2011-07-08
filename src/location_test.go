// Copyright © 2010 Esko Luontola <www.orfjackal.net>
// This software is released under the Apache License 2.0.
// The license text is at http://www.apache.org/licenses/LICENSE-2.0

package nanospec

import (
	"strings"
	"testing"
)


func Test__Get_location_of_calling_method(t *testing.T) {
	tt := TT(t)

	location := fakeExpectationMethod() // line 16

	parts := strings.Split(location, "/")
	tt.AssertEquals("location_test.go:16", parts[len(parts)-1])
}

func fakeExpectationMethod() string {
	return callerLocation()
}

func Test__Get_function_name_from_function_literal(t *testing.T) {
	tt := TT(t)

	name := functionName(fakeExpectationMethod)

	tt.AssertEquals("nanospec.fakeExpectationMethod", name)
}
