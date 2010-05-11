// Copyright © 2010 Esko Luontola <www.orfjackal.net>
// This software is released under the Apache License 2.0.
// The license text is at http://www.apache.org/licenses/LICENSE-2.0

package nanospec

import (
	"fmt"
	"os"
)


type Expectation struct {
	actual   interface{}
	reporter Reporter
}

func newExpectation(actual interface{}, reporter Reporter) *Expectation {
	return &Expectation{actual, reporter}
}

func (this *Expectation) Equals(expected interface{}) {
	if this.actual != expected {
		this.reporter.Error(fmt.Sprintf("Expected: equals '%v'\n\tgot: '%v'", expected, this.actual))
	}
}

func (this *Expectation) NotEquals(expected interface{}) {
	if this.actual == expected {
		this.reporter.Error(fmt.Sprintf("Expected: NOT equals '%v'\n\tgot: '%v'", expected, this.actual))
	}
}

func (this *Expectation) IsTrue() {
	if this.actual.(bool) != true {
		this.reporter.Error(fmt.Sprintf("Expected: is true\n\tgot: '%v'", this.actual))
	}
}

func (this *Expectation) IsFalse() {
	if this.actual.(bool) != false {
		this.reporter.Error(fmt.Sprintf("Expected: is false\n\tgot: '%v'", this.actual))
	}
}

func (this *Expectation) Satisfies(contract bool) {
	if !contract {
		this.reporter.Error(fmt.Sprintf("Expected: satisfies the contract\n\tgot: '%v'", this.actual))
	}
}

func (this *Expectation) Matches(matcher Matcher) {
	err := matcher(this.actual)
	if err != nil {
		this.reporter.Error(err.String())
	}
}

type Matcher func(actual interface{}) os.Error
