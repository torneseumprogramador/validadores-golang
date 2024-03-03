package validadores_golang

import "testing"

func TestValidarCNPJ(t *testing.T) {
	casosDeTeste := []struct {
		cnpj     string
		esperado bool
	}{
		{"11222333000181", true},
		{"11222333000182", false},
		{"00.000.000/0000-00", false},
		{"11.111.111/1111-11", false},
		{"12345678000195", false},
	}

	for _, caso := range casosDeTeste {
		resultado, err := ValidarCNPJ(caso.cnpj)
		if resultado != caso.esperado {
			t.Errorf("ValidarCNPJ(%s) = %v; esperado %v, (Erro: %v)", caso.cnpj, resultado, caso.esperado, err)
		}
	}
}

func TestValidarCNPJFormatado(t *testing.T) {
	cnpj := "25.823.284/0001-47"
	esperado := true
	resultado, err := ValidarCNPJ(cnpj)
	if resultado != esperado {
		t.Errorf("ValidarCNPJ(%s) = %v; esperado %v (erro: %v)", cnpj, resultado, esperado, err)
	}
}

func TestValidarCNPJDesformatado(t *testing.T) {
	cnpj := "25823284000147"
	esperado := true
	resultado, err := ValidarCNPJ(cnpj)
	if resultado != esperado {
		t.Errorf("ValidarCNPJ(%s) = %v; esperado %v (erro: %v)", cnpj, resultado, esperado, err)
	}
}

func TestValidarCNPJDesformatado2(t *testing.T) {
	cnpj := "25823284/0001-47"
	esperado := true
	resultado, err := ValidarCNPJ(cnpj)
	if resultado != esperado {
		t.Errorf("ValidarCNPJ(%s) = %v; esperado %v (erro: %v)", cnpj, resultado, esperado, err)
	}
}

func TestValidarCNPJEmpty(t *testing.T) {
	cnpj := ""
	esperado := false
	resultado, err := ValidarCNPJ(cnpj)
	if resultado != esperado {
		t.Errorf("ValidarCNPJ(%s) = %v; esperado %v (erro: %v)", cnpj, resultado, esperado, err)
	}
}
