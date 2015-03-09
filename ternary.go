package ternary

// Returns returns given "yes" value if condition is true.
func Returns(cond bool, yes, no interface{}) interface{} {
	if cond {
		return yes
	}
	return no
}

// Sets sets field if condition is true, otherwise map doesn't have this field.
func Sets(m map[string]interface{}, field string, cond bool, val interface{}) map[string]interface{} {
	if !cond {
		return m
	}
	m[field] = val
	return m
}
