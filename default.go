package ternary

import "reflect"

// DefaultValue ...
type DefaultValue struct {
	v interface{}
}

// Default ...
func Default(v interface{}) DefaultValue {
	return DefaultValue{v}
}

// String ...
func (d DefaultValue) String(s string) string {
	if reflect.ValueOf(d.v).Kind() != reflect.String {
		return s
	}
	return Zero(s).String(d.v.(string), s)
}
