package ternary

// ZeroableString ...
type ZeroableString string

// Or ...
func (s ZeroableString) Or(def string) string {
	if Zero(s) {
		return def
	}
	return string(s)
}

// String ...
func String(orig string) func(string) string {
	return ZeroableString(orig).Or
}
