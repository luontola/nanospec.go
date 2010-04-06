// Copyright © 2010 Esko Luontola <www.orfjackal.net>
// This software is released under the Apache License 2.0.
// The license text is at http://www.apache.org/licenses/LICENSE-2.0

package nanospec

import (
	"strings"
	"testing"
)


func Test__The_names_of_parent_specs_are_reported_on_failure(t *testing.T) {
	tt := TT(t)

	root := newSpec(nil, "FooSpec")
	parent := newSpec(root, "When foo")
	current := newSpec(parent, "Then bar")

	r := new(DummyReporter)
	newSpecReporter(r, current, "foo.go:42").Error("error message")

	tt.AssertEquals(trim(`
FooSpec
- When foo
  - Then bar

*** error message
    at foo.go:42
`),
		trim(r.Message))
}

var trim = strings.TrimSpace


type DummyReporter struct {
	Message string
}

func (this *DummyReporter) Error(message string) {
	this.Message = message
}
