package ternary

type Provider struct {
	cond bool
}

func If(cond bool) *Provider {
	return &Provider{cond}
}

func (p *Provider) Int(yes, no int) int {
	if p.cond {
		return yes
	}
	return no
}

func (p *Provider) String(yes, no string) string {
	if p.cond {
		return yes
	}
	return no
}
