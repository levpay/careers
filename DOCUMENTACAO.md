# Super Hero API

Este documento descreve como realizar a configuração do projeto.

## Configuração

Nesta seção é descritas as configurações necessárias para poder executar este projeto.

### Variáveis de ambiente

Para executar o projeto é necessário que as seguintes variáveis de ambiente estejam configuradas:

- `GOROOT` com o valor igual ao diretório raiz do Go;
- `GOPATH` com o valor igual ao diretório onde as dependências serão instaladas.
- `SUPERHEROAPI_TOKEN` com o valor do token da API do webservice [SuperHero API](https://superheroapi.com/);
- `DATABASE_HOST` com o valor do IP ou DNS do banco de dados;
- `DATABASE_PORT` com o valor da porta do banco de dados;
- `DATABASE_NAME` com o valor do nome do banco de dados;
- `DATABASE_USER` com o valor do login do usuário do banco de dados;
- `DATABASE_PASSWORD` com o valor da senha do usuário do banco de dados;
- `DATABASE_SSLMODE` com o valor informado se o SSL esta habilitado ou não (`disable` = SSL desabilitado; e `enable` = SSL habilitado).

### Versão do GO

Para o desenvolvimento desta API foi utilizada a versão 1.14 do Go.

### Dependências

Foram utilizadas as seguintes bibliotecas para o desenvolvimento desta API:
- github.com/gorilla/mux

É possível baixar todas as dependências executando o seguinte comando no diretório raiz do projeto:

```shell
cd src/github.com/carlsonsantana/superheroapi/
go mod download
```

## Executar

Para executar a API basta executar os seguintes comandos, no diretório raiz do projeto:

```shell
# Construindo o pacote "superheroapi"
cd src/github.com/carlsonsantana/superheroapi/
go build

# Migrando o banco de dados
cd migrations
go run *.go migrate
cd ..

# Instalando pacote "main"
cd main
go install

# Executanto aplicação
go run github.com/carlsonsantana/superheroapi/main
```

## Testar

Para testar a API basta executar os seguintes comandos, no diretório raiz do projeto:

```shell
# Construindo o pacote "superheroapi"
cd src/github.com/carlsonsantana/superheroapi/
go build

# Executanto testes
cd tests
go test
```
