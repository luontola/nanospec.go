// Copyright © 2010 Esko Luontola <www.orfjackal.net>
// This software is released under the Apache License 2.0.
// The license text is at http://www.apache.org/licenses/LICENSE-2.0

package nanospec

import (
	"container/vector"
	"testing"
)


func Run(gotest *testing.T, spec func(Context)) {
	c := newContext(gotest, spec)
	c.Run()
}

type Context interface {
	Specify(name string, closure func())
	Expect(actual interface{}) Expectation
}


type runContext struct {
	gotest       *testing.T
	rootClosure  func(Context)
	root         *aSpec
	current      *aSpec
	backtracking bool
}

func newContext(gotest *testing.T, spec func(Context)) *runContext {
	root := newSpec(nil, "<root>")
	return &runContext{gotest, spec, root, root, false}
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

func (this *runContext) Expect(actual interface{}) Expectation {
	return newExpectation(actual, newReporter(this.gotest, callerLocation()))
}


type aSpec struct {
	Parent                *aSpec
	name                  string
	children              *vector.Vector
	childrenSeenOnThisRun int
	hasBeenFullyExecuted  bool
}

func newSpec(parent *aSpec, name string) *aSpec {
	return &aSpec{parent, name, new(vector.Vector), 0, false}
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
	for _, v := range *this.children {
		child := v.(*aSpec)
		if !child.hasBeenFullyExecuted {
			return false
		}
	}
	return true
}

func (this *aSpec) EnterChild(name string) *aSpec {
	this.childrenSeenOnThisRun++
	childIndex := this.childrenSeenOnThisRun - 1

	isUnseen := childIndex >= this.children.Len()
	if isUnseen {
		child := newSpec(this, name)
		this.children.Push(child)
		return child
	}

	child := this.children.At(childIndex).(*aSpec)
	return child
}
