package ternary

import (
	"testing"

	. "github.com/otiai10/mint"
)

func TestReturns(t *testing.T) {
	Expect(t, Returns(true, nil, "foo")).ToBe(nil)
	Expect(t, Returns(false, nil, "bar")).ToBe("bar")
}

func TestIf(t *testing.T) {
	Expect(t, If(true)).TypeOf("ternary.Condition")
}

func TestCondition_String(t *testing.T) {
	cond := If(true)
	Expect(t, cond.String("yes", "no")).ToBe("yes")
}

func TestCondition_Int(t *testing.T) {
	cond := If(false)
	Expect(t, cond.Int(200, 404)).ToBe(404)
}
