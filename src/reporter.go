// Copyright © 2010 Esko Luontola <www.orfjackal.net>
// This software is released under the Apache License 2.0.
// The license text is at http://www.apache.org/licenses/LICENSE-2.0

package nanospec

import (
	"fmt"
	"testing"
)


type Reporter interface {
	Error(message string)
}


type specReporter struct {
	out      Reporter
	current  *aSpec
	location string
}

func newSpecReporter(out Reporter, current *aSpec, location string) Reporter {
	return specReporter{out, current, location}
}

func (this specReporter) Error(message string) {
	context := ""
	for spec := this.current; spec != nil; spec = spec.Parent {
		context = indent(spec) + spec.Name + "\n" + context
	}
	this.out.Error(fmt.Sprintf("%v\n*** %v\n    at %v\n", context, message, this.location))
}

func indent(spec *aSpec) string {
	if spec.Parent == nil {
		return ""
	}
	s := "- "
	for spec.Parent.Parent != nil {
		spec = spec.Parent
		s = "  " + s
	}
	return s
}


type gotestReporter struct {
	t *testing.T
}

func (this gotestReporter) Error(message string) {
	this.t.Error(message)
}
