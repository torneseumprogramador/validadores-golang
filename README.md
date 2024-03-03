# Validadores
- O objetivo desta lib é estudar a criação de bibliotecas em golang
- O mesmo foi criado no desafio de golang da comunidade do: https://www.torneseumprogramador.com.br/
- Toda gravação está disponilibizada aqui: https://www.torneseumprogramador.com.br/cursos/desafio_go_lang/aulas


## Lint de código
```shell
sh lint.sh
```

## Rodar o código
```shell
sh run.sh
```

## Testes
```shell
sh tests.sh
```

## Estrutura de arquivos
```
validadores/
├── .gitignore                  # Arquivo para ignorar arquivos no Git
├── go.mod                      # Arquivo de módulo Go, contendo dependências
├── go.sum                      # Soma de verificação para as dependências
├── README.md                   # Documentação inicial do projeto
├── cpf.go                      # Implementação da validação de CPF
├── cpf_test.go                 # Testes para a validação de CPF
├── cnpj.go                     # Implementação da validação de CNPJ
├── cnpj_test.go                # Testes para a validação de CNPJ
└── examples/                   # Exemplos de como usar sua biblioteca
    ├── example_cpf.go          # Exemplo de como validar CPF
    └── example_cnpj.go         # Exemplo de como validar CNPJ
```


## Como utilizar

```go
package main

import (
	"fmt"
	_ "github.com/torneseumprogramador/validadores-golang"
)

func main() {
    cpf := "748.469.830-05"
    esperado := true
    resultado, err := validadores.ValidarCPF(cpf)
    if resultado != esperado {
        fmt.Printf("ValidarCPF(%s) = %v; esperado %v (erro: %v)", cpf, resultado, esperado, err)
    } else {
        fmt.Println("CPF (%s) validado com sucesso", cpf)
    }


    cpf = "748.469.830-01"
    resultado, err = validadores.ValidarCPF(cpf)
    fmt.Printf("ValidarCPF(%s) = %v; esperado %v (erro: %v)", cpf, resultado, esperado, err)
}
```