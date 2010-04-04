// Copyright © 2010 Esko Luontola <www.orfjackal.net>
// This software is released under the Apache License 2.0.
// The license text is at http://www.apache.org/licenses/LICENSE-2.0

package nanospec

import (
	"testing"
)


func Test__At_first_no_specs_have_been_seen(t *testing.T) {
	tt := TT(t)

	c := newContext(t, func(c Context) {})
	c.Run()

	tt.AssertEquals(0, c.root.children.Len())
}

func Test__When_direct_child_specs_are_seen__Then_they_are_remembered(t *testing.T) {
	tt := TT(t)

	c := newContext(t, func(c Context) {
		c.Specify("a", func() {})
		c.Specify("b", func() {})
	})
	c.Run()

	tt.AssertEquals(2, c.root.children.Len())
}

func Test__When_nested_child_specs_are_seen__Then_they_are_remembered(t *testing.T) {
	tt := TT(t)

	c := newContext(t, func(c Context) {
		c.Specify("a", func() {
			c.Specify("aa", func() {})
		})
	})
	c.Run()

	tt.AssertEquals(1, c.root.children.Len())
	tt.AssertEquals(1, c.root.children.At(0).(*aSpec).children.Len())
}

func Test__When_sibling_specs_are_seen__Then_they_are_remembered_only_once(t *testing.T) {
	tt := TT(t)

	c := newContext(t, func(c Context) {
		c.Specify("a", func() {
			c.Specify("aa", func() {})
			c.Specify("ab", func() {})
		})
		c.Specify("b", func() {
			c.Specify("ba", func() {})
			c.Specify("bb", func() {})
		})
	})
	c.Run()

	tt.AssertEquals(2, c.root.children.Len())
	tt.AssertEquals(2, c.root.children.At(0).(*aSpec).children.Len())
	tt.AssertEquals(2, c.root.children.At(1).(*aSpec).children.Len())
}
