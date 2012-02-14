// Copyright © 2010 Esko Luontola <www.orfjackal.net>
// This software is released under the Apache License 2.0.
// The license text is at http://www.apache.org/licenses/LICENSE-2.0

package nanospec

import (
	"fmt"
	"testing"
)

func Run(gotest *testing.T, spec func(Context)) {
	c := newContext(gotest, spec)
	c.Run()
}

type Context interface {
	Specify(name string, closure func())
	Expect(actual interface{}) *Expectation
	Errorf(format string, args ...interface{})
}

type runContext struct {
	out          Reporter
	rootClosure  func(Context)
	root         *aSpec
	current      *aSpec
	backtracking bool
}

func newContext(t *testing.T, spec func(Context)) *runContext {
	root := newSpec(nil, functionName(spec))
	return &runContext{gotestReporter{t}, spec, root, root, false}
}

func (this *runContext) Run() {
	safetyLimit := 10000 // just in case this program gets stuck in an infinite loop during development
	for this.root.ShouldExecute() && safetyLimit > 0 {
		this.backtracking = false
		this.root.Execute(func() { this.rootClosure(this) })
		safetyLimit--
	}
}

func (this *runContext) Specify(name string, closure func()) {
	this.enterSpec(name)
	this.processSpec(closure)
	this.exitSpec()
}

func (this *runContext) enterSpec(name string) {
	child := this.current.EnterChild(name)
	this.current = child
}

func (this *runContext) processSpec(closure func()) {
	if !this.backtracking && this.current.ShouldExecute() {
		this.current.Execute(closure)
		this.backtracking = true
	}
}

func (this *runContext) exitSpec() {
	this.current = this.current.Parent
}

func (this *runContext) Expect(actual interface{}) *Expectation {
	reporter := newSpecReporter(this.out, this.current, callerLocation())
	return newExpectation(actual, reporter)
}

func (this *runContext) Errorf(format string, args ...interface{}) {
	reporter := newSpecReporter(this.out, this.current, callerLocation())
	reporter.Error(fmt.Sprintf(format, args...))
}

type aSpec struct {
	Parent                *aSpec
	Name                  string
	children              []*aSpec
	childrenSeenOnThisRun int
	hasBeenFullyExecuted  bool
}

func newSpec(parent *aSpec, name string) *aSpec {
	return &aSpec{parent, name, make([]*aSpec, 0), 0, false}
}

func (this *aSpec) ShouldExecute() bool {
	return !this.hasBeenFullyExecuted
}

// The closure of this spec must be passed as a parameter,
// to make sure it's fresh; no side-effects from previous runs.
func (this *aSpec) Execute(closure func()) {
	this.childrenSeenOnThisRun = 0
	closure()
	this.hasBeenFullyExecuted = this.allChildrenHaveBeenExecuted()
}

func (this *aSpec) allChildrenHaveBeenExecuted() bool {
	for _, child := range this.children {
		if !child.hasBeenFullyExecuted {
			return false
		}
	}
	return true
}

func (this *aSpec) EnterChild(name string) *aSpec {
	this.childrenSeenOnThisRun++
	childIndex := this.childrenSeenOnThisRun - 1

	isUnseen := childIndex >= len(this.children)
	if isUnseen {
		child := newSpec(this, name)
		this.children = append(this.children, child)
		return child
	}

	child := this.children[childIndex]
	return child
}
