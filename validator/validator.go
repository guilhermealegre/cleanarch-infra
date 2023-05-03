package validator

import (
	infra "github.com/guilhermealegre/cleanarch-infra"
)

type validator struct {
}

func New() infra.IValidator {
	return &validator{}
}

func (*validator) Init() error {

	return nil
}
