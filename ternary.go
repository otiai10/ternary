package ternary

func Returns(cond bool, yes, no interface{}) interface{} {
	if cond {
		return yes
	}
	return no
}
