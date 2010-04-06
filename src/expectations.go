// Copyright © 2010 Esko Luontola <www.orfjackal.net>
// This software is released under the Apache License 2.0.
// The license text is at http://www.apache.org/licenses/LICENSE-2.0

package nanospec

import (
	"fmt"
)


type Reporter interface {
	Error(message string)
}

type Expectation struct {
	actual   interface{}
	reporter Reporter
}

func newExpectation(actual interface{}, reporter Reporter) Expectation {
	return Expectation{actual, reporter}
}

func (this Expectation) Equals(expected interface{}) {
	if this.actual != expected {
		this.reporter.Error(fmt.Sprintf("'%v' should equal '%v'", this.actual, expected))
	}
}

func (this Expectation) NotEquals(expected interface{}) {
	if this.actual == expected {
		this.reporter.Error(fmt.Sprintf("'%v' should NOT equal '%v'", this.actual, expected))
	}
}

func (this Expectation) IsTrue() {
	if this.actual.(bool) != true {
		this.reporter.Error(fmt.Sprintf("'%v' should be true", this.actual))
	}
}

func (this Expectation) IsFalse() {
	if this.actual.(bool) != false {
		this.reporter.Error(fmt.Sprintf("'%v' should be false", this.actual))
	}
}

func (this Expectation) Satisfies(contract bool) {
	if !contract {
		this.reporter.Error(fmt.Sprintf("'%v' should satisfy the contract", this.actual))
	}
}
