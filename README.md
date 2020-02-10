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
