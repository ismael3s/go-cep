package domain_test

import (
	"testing"

	"github.com/ismael3s/go-cep/internal/domain"
)

func Test_GivenAnCEP_ShouldBeAbleToRemoveFormating(t *testing.T) {
	tt := []struct {
		Input    string
		Expected string
	}{
		{
			Input:    "40283-310",
			Expected: "40283310",
		},
		{
			Input:    "40281-310",
			Expected: "40281310",
		},
	}

	for _, tc := range tt {
		cep, err := domain.NewCEP(tc.Input)
		if err != nil {
			t.Errorf("Expected that error should be nil")
		}
		if cep.GetValue() != tc.Expected {
			t.Errorf("Expected %s, got %s", tc.Expected, tc.Input)
		}
	}
}

func Test_GivenAnCEP_WithMissingDigits_ShouldReturnAn_Invalid_Formating_ERROR(t *testing.T) {
	tt := []struct {
		Input       string
		Expected    string
		ExpectedErr string
	}{
		{
			Input:       "40283-31",
			Expected:    "",
			ExpectedErr: "CEP deve ser válido",
		},
	}

	for _, tc := range tt {
		_, err := domain.NewCEP(tc.Input)
		if err == nil {
			t.Errorf("Want error 'CEP deve ser válido', but get nothing")
		}
		if err.Error() != tc.ExpectedErr {
			t.Errorf("Expected %s, got %s", tc.ExpectedErr, err.Error())
		}
	}
}
