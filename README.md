# Documentação especifica dessa instância desafio

## Dependências

Esse projeto usa módulos então você precisa da linguagem Go >= 1.13.

## Execute os testes (não precisa de postgresql)

```
go test ./...
```

## Compile o projeto

```
go build -o superheroapi
```

## Execute 

```
./superheroapi --help
```

## Testando com cURL

Todos os supers
```
curl localhost:8000/supers/ -H "Content-Type: application/json"
```

Todos os supers bonzinhos
```
curl localhost:8000/supers/good -H "Content-Type: application/json"
```

Todos os supers malvados
```
curl localhost:8000/supers/bad -H "Content-Type: application/json"
```

Adicione novo super
```
curl localhost:8000/supers/ -H "Content-Type: application/json" -d '{"name":"superman",
"groups": [{"name": "Justice League"}]}'
```

Consulte um super, uuid abaixo é um exemplo!
```
curl localhost:8000/supers/68ca3ad5-9dc6-4ee0-b11a-ad0d1513c1d5 -H "Content-Type: application/json"
```

Procure um super pelo nome
```
curl -H 'Content-Type: application/json' localhost:8000/supers/search/superman
```

Delete um super, uuid abaixo é um exemplo!
```
curl localhost:8000/supers/68ca3ad5-9dc6-4ee0-b11a-ad0d1513c1d5 -H "Content-Type: application/json" -X DELETE
```

## PostgreSQL
O projeto conta com migrações automáticas e para isso precisamos de um banco de 
dados e um usuário com permissões para modificar o banco em questão. Alguns exemplos abaixo.

```
createdb --owner=superadminuser super
```

Sempre existe a opção de rodar um docker de postgres.

Você também precisa renomear o arquivo ```configuration.toml.example``` para ```configuration.toml```.

Precisamos editar o arquivo com as configurações do seu postgres e banco criado.

```
[database]
dsn = "postgres://<usuário>:<senha>@<endereço>/<nome do banco>?<opções>"
```

_dsn_ significa data source name, mais informações podem ser encontradas em
https://pkg.go.dev/github.com/lib/pq?tab=doc

## Outras configurações

Edite o ```configuration.toml``` e escolha o endereço e parta de execução.

```
[server]
bind = "127.0.0.1:8000"
```

# Levpay

A Levpay é do ramo de meios de pagamentos e tem como maior meta atender com excelência os seus clientes e parceiros, por isso, estamos procurando profissionais que gostem de impactar a vida de outras pessoas especialmente através da tecnologia e que queiram crescer junto com a empresa.

Nossa equipe de engenharia é 100% remoto e usamos as seguintes tecnologias:
- Golang, Angular, PostgreSQL, MongoDB, AMQP
- Containers, Continuous Delivery, TDD

# Como faço para trabalhar como desenvolvedor na Levpay?

1. Envie um email para rh@levpay.com contendo:
    - Assunto: golang developer
    - Corpo: Github / LinkedIn / Curriculo / breve texto sobre sua experiência profissional
2. Faça um fork (ou clone) esse repositório
3. Complete o desafio abaixo no seu tempo
4. Nos envie o desafio da forma que lhe for conveniente - email ou pull request.

# Desafio
## Golang Developer
Para o desafio gostaríamos que você crie uma API inspirada em super heróis e vilões para servir um jogo utilizando a SuperHeroAPI (https://superheroapi.com/).

## Requisitos

A API deve ser escrita em **Golang** e utilizar **PostgreSQL** como armazenamento de dados.

### Gerais
Através da API deve ser possível:
- Cadastrar um Super/Vilão
- Listar todos os Super's cadastrados
- Listar apenas os Super Heróis
- Listar apenas os Super Vilões
- Buscar por nome
- Buscar por 'uuid'
- Remover o Super

### Específicos
- API deve seguir a arquitetura [REST](https://restfulapi.net/)
- API deve seguir os principios do [12 factor app](https://12factor.net/pt_br/)
- Cada super deve ser cadastrado somente a partir do seu `name`.
- A pesquisa por um super deve conter os seguintes campos:
    - uuid
    - name
    - full name
    - intelligence
    - power
    - occupation
    - image
- A pesquisa por um super também precisa conter:
    - lista de grupos em que tal super está associado
    - número de parentes

## Como será a avaliação

A ideia aqui é entender como você toma suas decisões frente a certas adversidades e como você desenvolve através de multiplas funcionalidades.

Pontos que vamos avaliar:
- Commits
    - como você evoluiu seu pensamento durante o projeto, pontualidade e clareza.
- Testes
    - Quanto mais testes melhor! Vide https://code.tutsplus.com/pt/tutorials/lets-go-testing-golang-programs--cms-26499 .
- Complexidade
    - Código bom é código legivel e simples (https://medium.com/trainingcenter/golang-d94e16d4b383).
- Dependências
    - O ecosistema (https://github.com/avelino/awesome-go) da linguagem possui diversas ferramentas para serem usadas, use-as bem!
- Documentação
    - Qual versão de Go você usou?
    - Quais bibliotecas e ferramentas usou?
    - Como se utiliza a sua aplicação?
    - Como executamos os testes?
- Considerações
    - as regras de negócio não foram definidas intencionalmente
    - cabe a você decidir como vai manter os cadastros no banco da aplicação
    - cabe a você decidir como vai tratar cadastros repetidos
