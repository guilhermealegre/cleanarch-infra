package validator

import (
	"auth/infra"
)

type validator struct {
}

func New() infra.IValidator {
	return &validator{}
}

func (*validator) Init() error {

	return nil
}
