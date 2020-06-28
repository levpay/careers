# Super Hero API

Este documento descreve como realizar a configuração do projeto.

## Configuração

Nesta seção é descritas as configurações necessárias para poder executar este projeto.

### Variáveis de ambiente

Para executar o projeto é necessário que as seguintes variáveis de ambiente estejam configuradas:

- `HTTP_PORT` com o valor igual a port HTTP do WebService;
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
- github.com/arthurkushman/buildsqlx
- github.com/lib/pq
- github.com/satori/go.uuid
- github.com/go-pg/pg/v9
- github.com/robinjoseph08/go-pg-migrations/v2

É possível baixar todas as dependências executando o seguinte comando no diretório raiz do projeto:

```shell
cd src/github.com/carlsonsantana/superheroapi/
go mod download
cd migrations
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

## API

Nesta seção será descrito o funcionamento da API.

### Resposta da API

A API retorna a resposta em um object JSON com a seguinte estrutura:

- `status` (**string**): informa se a operação foi realizada sem nenhuma falha, quando a operação for realizada com sucesso o deste campo será igual à `success` e quando ocorrer alguma falha seja da requisição ou da aplicação o seu valor será `failed`;
- `supers` (**object[]**): retorna os Supers que foram cadastrados, pesquisados ou excluídos:
  - `uuid` (**string**):  identificador único do Super;
  - `superheroapi-id` (**int**): ID do Super para SuperHero API;
  - `name` (**string**): nome do Super;
  - `full-name` (**string**): nome completo do Super;
  - `intelligence` inteligência do Super;
  - `power` (**int**): poder do Super;
  - `occupation` (**int**): (**string**): ocupação do Super;
  - `image` (**string**): imagem do Super;
  - `groups` (**string[]**): grupos aos quais o Super pertence ou já pertenceu;
  - `category` (**string**): categoria do Super: héroi (`hero`), vilão (`villain`) ou outro (`neutral`);
  - `number-relatives` (**int**): número de parentes do Super;
- `message` (**string**): mensagem de erro da operação, este campo só é preenchido quando o campo `status` for igual a `failed`.

Além da propriedade `status`, falhas nas operações também são informadas através dos [Códigos de Status HTTP](https://en.wikipedia.org/wiki/List_of_HTTP_status_codes), sendo retornado o código **200** para todas as operações realizadas com sucesso e códigos **4XX** e **5XX** para todas as operações onde ocorreram falhas.

### Operações na API

Nesta seção serão descritas as operações possíveis na API.

#### Cadastro de Supers

Para cadastrar um conjunto de Supers nesta API é necessário realizar uma requisição através do método `POST`, cujo o cabeçalho `Content-Type` tenha um valor igual a `application/x-www-form-urlencoded` para o caminho `/super`, sendo necessário também informar um parâmetro `name` com parte do nome dos Supers.

Mesmo que todos os Supers já estejam cadastrados na API, ela não irá cadastrá-los novamente, porém a API não retornará uma mensagem de erro, assim como retornará os Supers com que seus respectivos identificadores únicos.

Caso não seja informando o parâmetro `name` a API não realizará nenhuma operação e tornará uma mensagem de erro.

```shell
# Cadastrará todos os Supers que possuem o nome Batman
curl -d "name=batman" http://localhost:8080/super
```

#### Listagem de Supers

Para listar os Supers cadastrados nesta API é necessário realizar uma requisição `GET` para o caminho `/super`.

É possível filtrar os Supers passando como parâmetros os atributos do Super, conforme a lista informada na seção **Resposta da API**. Caso não seja informado nenhum parâmetro de filtro todos os Supers serão listados.

##### Parâmetros de filtros

Esta seção descreve os parâmetros de pesquisa da listagem de Supers.

Caso seja informado um parâmetro que não corresponde a um atributo de um Super ou seja informado um valor inválido a API retornará uma mensagem de erro.

A seguir os parâmetros de filtros:

- **string exata** os parâmetros deste tipo são aqueles em que retornará um resultado apenas se todo o texto for correspondente e são case-sensitive, são eles:
  - `uuid`
  - `image`
  - `category`
- **string parcial** os parâmetros deste tipo são aqueleas em que é possível pesquisar apenas por uma parte dela, permitem ignorar certas partes do nome utilizando o símbolo `%` e são caso-insentive, são eles:
  - `name`
  - `full-name`
  - `occupation`
  - `groups`
- **número exato** os parâmetros deste tipo só poderão filtram os Supers que tiverem exatamente este valor como seu atributo, são eles:
  - `superheroapi-id`
- **número compativo** os parâmetros deste tipo permitem filtrar pelo número exato (quando é informado apenas o número), números maiores que o número informado (iniciado com o símbolo `>` e após isso o número a qual será comparado), números maiores ou iguais ao número informado (iniciado com os símbolos `>=` e após isso o número a qual será comparado), números menores que o número informado (iniciado com o símbolo `<` e após isso o número a qual será comparado) e números menores ou iguais ao número informado (iniciado com os símbolos `<=` e após isso o número a qual será comparado), são eles:
  - `intelligence`
  - `power`
  - `number-relatives`

```shell
# Listar todos os Supers
curl http://localhost:8080/super

# Listar todos os Supers com o nome 'Green Arrow'
curl http://localhost:8080/super?name=green%20arrow

# Listar todos os Supers em que seu nome real possuam o primeiro nome 'Bruce'
curl http://localhost:8080/super?full-name=bruce%20%25

# Listar todos os Supers que possuem no nome os caracteres 'man'
curl http://localhost:8080/super?name=%25man%25

# Listar todos os Supers que possuem um poder igual a 100
curl http://localhost:8080/super?power=100

# Listar todos os Supers que possuem menos de cinco parentes
curl http://localhost:8080/super?number-relatives=%3c5
```

#### Exclusão de Super

Para deletar um Super cadastrado nesta API é necessário realizar uma requisição `DELETE` para o caminho `/super/{uuid}`, onde **{uuid}** deve ser substituído pelo identificador único do Super.

Caso não exista nenhum Super com o identificador único informado a API informará que nenhum Super foi encontrado.

```shell
# Excluindo o Super com o identificador único "291a9629-a020-4921-808d-cb821a66969e"
curl -X "DELETE" http://localhost:8080/super/291a9629-a020-4921-808d-cb821a66969e
```
