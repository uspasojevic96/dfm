package pkg

import "errors"

type Package struct {
	Name    string
	Version string
	Script  string
}

func (p *Package) Uninstall() error {
	return errors.New("not implemented")
}
