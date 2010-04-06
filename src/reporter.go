// Copyright © 2010 Esko Luontola <www.orfjackal.net>
// This software is released under the Apache License 2.0.
// The license text is at http://www.apache.org/licenses/LICENSE-2.0

package nanospec

import (
	"testing"
)


type Reporter interface {
	Error(message string)
}

type theReporter struct {
	gotest   *testing.T
	location string
}

func newReporter(gotest *testing.T, location string) Reporter {
	return theReporter{gotest, location}
}

func (this theReporter) Error(message string) {
	this.gotest.Error(message, "in", this.location)
}
