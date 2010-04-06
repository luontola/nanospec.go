// Copyright © 2010 Esko Luontola <www.orfjackal.net>
// This software is released under the Apache License 2.0.
// The license text is at http://www.apache.org/licenses/LICENSE-2.0

package nanospec

import (
	"fmt"
	"testing"
)


func Test__Sibling_specs_are_executed(t *testing.T) {
	tt := TT(t)
	spy := ""

	Run(t, func(c Context) {
		c.Specify("a", func() {
			spy += "a,"
		})
		c.Specify("b", func() {
			spy += "b,"
		})
	})

	tt.AssertEquals("a,b,", spy)
}

func Test__Nested_specs_are_executed(t *testing.T) {
	tt := TT(t)
	spy := ""

	Run(t, func(c Context) {
		c.Specify("a", func() {
			spy += "a,"

			c.Specify("aa", func() {
				spy += "aa,"
			})
		})
	})

	tt.AssertEquals("a,aa,", spy)
}

func Test__Nested_sibling_specs_are_executed_in_isolation(t *testing.T) {
	tt := TT(t)
	spy := ""

	Run(t, func(c Context) {
		c.Specify("a", func() {
			spy += "a,"

			c.Specify("aa", func() {
				spy += "aa,"
			})
			c.Specify("ab", func() {
				spy += "ab,"
			})
		})
		c.Specify("b", func() {
			spy += "b,"

			c.Specify("ba", func() {
				spy += "ba,"
			})
			c.Specify("bb", func() {
				spy += "bb,"
			})
			c.Specify("bc", func() {
				spy += "bc,"
			})
		})
	})

	tt.AssertEquals("a,aa,a,ab,b,ba,b,bb,b,bc,", spy)
}

func Test__Variables_declared_inside_specs_are_isolated_from_side_effects(t *testing.T) {
	tt := TT(t)
	spy := ""

	Run(t, func(c Context) {
		common := 0

		c.Specify("a", func() {
			common++
			spy += fmt.Sprintf("%v,", common)
		})
		c.Specify("b", func() {
			common++
			spy += fmt.Sprintf("%v,", common)
		})
	})

	tt.AssertEquals("1,1,", spy)
}
