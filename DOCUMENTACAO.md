# Super Hero API

Este documento descreve como realizar a configuração do projeto.

## Configuração

Nesta seção é descritas as configurações necessárias para poder executar este projeto.

### Variáveis de ambiente

Para executar o projeto é necessário que as seguintes variáveis de ambiente estejam configuradas:

- `GOROOT` com o valor igual ao diretório raiz do Go;
- `GOPATH` com o valor igual ao diretório onde as dependências serão instaladas.

### Versão do GO

Para o desenvolvimento desta API foi utilizada a versão 1.14 do Go.

### Dependências

Foram utilizadas as seguintes bibliotecas para o desenvolvimento desta API:
- github.com/gorilla/mux

É possível baixar todas as dependências executando o seguinte comando no diretório raiz do projeto:

```shell
go mod download
```

## Executar

Para executar a API basta executar o seguinte comando, no diretório raiz do projeto:

```shell
go run src/github.com/carlsonsantana/superheroapi/*.go
```
