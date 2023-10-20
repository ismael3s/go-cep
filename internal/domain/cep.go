package domain

import (
	"errors"
	"regexp"
)

var INVALID_CEP_FORMAT = errors.New("CEP deve ser válido")

type CEP struct {
	value string
}

func (c CEP) GetValue() string {
	return c.value
}

func (c *CEP) validate() error {
	ONLY_DIGITS_REGEX := regexp.MustCompile(`\D`)
	formatedCEP := ONLY_DIGITS_REGEX.ReplaceAllString(c.value, "")
	if len(formatedCEP) != 8 {
		return INVALID_CEP_FORMAT
	}
	c.value = formatedCEP
	return nil
}

func NewCEP(value string) (*CEP, error) {
	cep := &CEP{value: value}
	if err := cep.validate(); err != nil {
		return nil, err
	}
	return cep, nil
}
