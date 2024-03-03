package validadores

import (
	"testing"
)

// TestValidarCPF testa a função ValidarCPF com diferentes entradas.
func TestValidarCPF(t *testing.T) {
	casosDeTeste := []struct {
		cpf      string
		esperado bool
	}{
		{"00000000000", false},
		{"11111111111", false},
		{"12345678909", false},
		{"987.654.321-00", true},
		{"98765432100", true},
		{"987.654.321-01", false},
		{"123.456.789-09", false},
	}

	for _, caso := range casosDeTeste {
		resultado, err := ValidarCPF(caso.cpf)
		if resultado != caso.esperado {
			t.Errorf("ValidarCPF(%s) = %v; esperado %v (erro: %v)", caso.cpf, resultado, caso.esperado, err)
		}
	}
}

// TestValidarCPFFormatado testa a função ValidarCPF com CPFs formatados.
func TestValidarCPFFormatado(t *testing.T) {
	cpf := "748.469.830-05"
	esperado := true
	resultado, err := ValidarCPF(cpf)
	if resultado != esperado {
		t.Errorf("ValidarCPF(%s) = %v; esperado %v (erro: %v)", cpf, resultado, esperado, err)
	}
}

func TestValidarCPFDesformatado(t *testing.T) {
	cpf := "74846983005"
	esperado := true
	resultado, err := ValidarCPF(cpf)
	if resultado != esperado {
		t.Errorf("ValidarCPF(%s) = %v; esperado %v (erro: %v)", cpf, resultado, esperado, err)
	}
}

func TestValidarCPFEmpty(t *testing.T) {
	cpf := ""
	esperado := false
	resultado, err := ValidarCPF(cpf)
	if resultado != esperado {
		t.Errorf("ValidarCPF(%s) = %v; esperado %v (erro: %v)", cpf, resultado, esperado, err)
	}
}
