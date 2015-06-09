package ternary

type Condition bool

func If(cond bool) Condition {
	return Condition(cond)
}

func (cond Condition) Int(yes, no int) int {
	if cond {
		return yes
	}
	return no
}

func (cond Condition) String(yes, no string) string {
	if cond {
		return yes
	}
	return no
}
