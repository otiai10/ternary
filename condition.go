package ternary

// Condition represents boolean flag itself.
type Condition bool

// If construct Condition struct, to prepare to return anything.
func If(flag bool) Condition {
	return Condition(flag)
}

// Int concludes to return `int` type value, according to given condition.
func (cond Condition) Int(yes, no int) int {
	if cond {
		return yes
	}
	return no
}

// String concludes to return `string` type value, according to given condition.
func (cond Condition) String(yes, no string) string {
	if cond {
		return yes
	}
	return no
}

// Interface concludes to return `string` type value, according to given condition.
func (cond Condition) Interface(yes, no interface{}) interface{} {
	if cond {
		return yes
	}
	return no
}

// Put provides put-func to place key-value on map following.
//
// It allows putting key-value on given target if flag is true,
// or do nothing otherwise.
// It means the given key will not exist on the target map.
func (cond Condition) Put(key string, value interface{}) func(map[string]interface{}) {
	return func(target map[string]interface{}) {
		if cond {
			target[key] = value
		}
		// else, do nothing
	}
}
