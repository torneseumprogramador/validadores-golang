package validadores_golang

import (
	"errors"
	"regexp"
	"strconv"
)

// Função para validar CNPJ
func ValidarCNPJ(cnpj string) (bool, error) {
	// Remover caracteres não numéricos
	re := regexp.MustCompile(`\D`)
	cnpj = re.ReplaceAllString(cnpj, "")

	// Verificar se o CNPJ tem 14 dígitos
	if len(cnpj) != 14 {
		return false, errors.New("CNPJ deve conter 14 dígitos")
	}

	// Lista de CNPJs inválidos
	cnpjsInvalidos := []string{
		"00000000000000", "11111111111111", "22222222222222", "33333333333333",
		"44444444444444", "55555555555555", "66666666666666", "77777777777777",
		"88888888888888", "99999999999999", "12345678000195",
	}
	for _, v := range cnpjsInvalidos {
		if cnpj == v {
			return false, errors.New("CNPJ não pode ter numeros repetidos")
		}
	}

	// Cálculo dos dígitos verificadores
	tamanho := len(cnpj) - 2
	numeros := cnpj[:tamanho]
	digitos := cnpj[tamanho:]

	d1 := calculaDigito(numeros, tamanho)
	d2 := calculaDigito(numeros+d1, tamanho+1)

	return digitos == d1+d2, nil
}

// Função auxiliar para calcular cada dígito verificador
func calculaDigito(numeros string, tamanho int) string {
	var soma int
	pos := tamanho - 7
	for i := tamanho; i >= 1; i-- {
		num, _ := strconv.Atoi(string(numeros[tamanho-i]))
		soma += num * pos
		pos--
		if pos < 2 {
			pos = 9
		}
	}

	resultado := soma % 11
	if resultado < 2 {
		return "0"
	} else {
		return strconv.Itoa(11 - resultado)
	}
}
