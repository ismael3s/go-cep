//go:build e2e
// +build e2e

package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ismael3s/go-cep/internal/infra/rest"
)

func Test_ShouldReturnAnAddressBased_OnAnCEP(t *testing.T) {
	ts := httptest.NewServer(rest.SetupRestServer())
	defer ts.Close()
	res, err := http.Get(ts.URL + "/cep/40283-310")
	if err != nil {
		t.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		t.Fatal(err)
	}

	expectedBody := `{"cep":"40283310","logradouro":"Rua Rio Branco","bairro":"Brotas","cidade":"Salvador"}`
	if string(body) != strings.ReplaceAll(expectedBody, `\n`, "") {
		t.Errorf("Expected %s, got %s", expectedBody, string(body))
	}

}
