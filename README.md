# Validadores Golang
- O objetivo desta lib é estudar a criação de bibliotecas em golang
- O mesmo foi criado no desafio de golang da comunidade do: https://www.torneseumprogramador.com.br/
- Toda gravação está disponilibizada aqui: https://www.torneseumprogramador.com.br/cursos/desafio_go_lang/aulas


## Lint de código
```shell
sh lint.sh
```

## Testes
```shell
sh tests.sh
```

## Estrutura de arquivos
```
validadores_golang/
├── .gitignore                  # Arquivo para ignorar arquivos no Git
├── go.mod                      # Arquivo de módulo Go, contendo dependências
├── README.md                   # Documentação inicial do projeto
├── cpf.go                      # Implementação da validação de CPF
├── cpf_test.go                 # Testes para a validação de CPF
├── cnpj.go                     # Implementação da validação de CNPJ
├── cnpj_test.go                # Testes para a validação de CNPJ
├── lint.sh                     # Formata e valida código
├── tests.sh                    # Roda os testes
```


## Como instalar
```shell
go get github.com/torneseumprogramador/validadores_golang # ultima versão
go get github.com/torneseumprogramador/validadores_golang@v1.0.0 # especificando a versão
```

## Como utilizar validador CPF
```go
package main

import (
	"fmt"
	_ "github.com/torneseumprogramador/validadores_golang"
)

func main() {
    cpf := "748.469.830-05"
    esperado := true
    resultado, err := validadores_golang.ValidarCPF(cpf)
    if resultado != esperado {
        fmt.Printf("ValidarCPF(%s) = %v; esperado %v (erro: %v)", cpf, resultado, esperado, err)
    } else {
        fmt.Println("CPF (%s) validado com sucesso", cpf)
    }


    cpf = "748.469.830-01"
    resultado, err = validadores_golang.ValidarCPF(cpf)
    fmt.Printf("ValidarCPF(%s) = %v; esperado %v (erro: %v)", cpf, resultado, esperado, err)
}
```


## Como utilizar validador CNPJ
```go
package main

import (
	"fmt"
	_ "github.com/torneseumprogramador/validadores_golang"
)

func main() {
    cnpj := "25.823.284/0001-47"
	esperado := true
	resultado, err := validadores_golang.ValidarCNPJ(cnpj)
    if resultado != esperado {
        fmt.Printf("ValidarCNPJ(%s) = %v; esperado %v (erro: %v)", cpf, resultado, esperado, err)
    } else {
        fmt.Println("CNPJ (%s) validado com sucesso", cnpj)
    }


    cnpj = "25.823.284/0001-48"
    resultado, err = validadores_golang.ValidarCNPJ(cnpj)
    fmt.Printf("ValidarCNPJ(%s) = %v; esperado %v (erro: %v)", cpf, resultado, esperado, err)
}
```