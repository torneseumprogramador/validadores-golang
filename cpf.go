package validadores

import (
	"errors"
	"strconv"
	"strings"
)

// ValidarCPF valida um CPF, retornando true se for válido e false caso contrário.
// Um CPF válido deve ter 11 dígitos e passar no teste dos dois dígitos verificadores.
func ValidarCPF(cpf string) (bool, error) {
	// Remove caracteres não numéricos
	cpf = strings.Join(strings.FieldsFunc(cpf, func(r rune) bool {
		return !strings.ContainsRune("0123456789", r)
	}), "")

	if cpf == "12345678909" {
		return false, errors.New("CPF 12345678909 não é permitido")
	}

	if len(cpf) != 11 {
		return false, errors.New("CPF deve conter 11 dígitos")
	}

	if todosDigitosIguais(cpf) {
		return false, nil
	}

	// Calcula e verifica os dígitos verificadores
	if !verificaDigito(cpf, 10) || !verificaDigito(cpf, 11) {
		return false, nil
	}

	return true, nil
}

// todosDigitosIguais verifica se todos os dígitos do CPF são iguais.
func todosDigitosIguais(cpf string) bool {
	for i := 1; i < len(cpf); i++ {
		if cpf[i] != cpf[0] {
			return false
		}
	}
	return true
}

// verificaDigito verifica um dígito verificador específico do CPF.
func verificaDigito(cpf string, posicao int) bool {
	soma, peso := 0, posicao
	for i := 0; i < posicao-1; i++ {
		num, _ := strconv.Atoi(string(cpf[i]))
		soma += num * peso
		peso--
	}

	resto := soma % 11
	if resto < 2 {
		resto = 0
	} else {
		resto = 11 - resto
	}

	digito, _ := strconv.Atoi(string(cpf[posicao-1]))
	return resto == digito
}
