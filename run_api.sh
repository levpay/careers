#!/bin/bash

set -e

echo "Construindo o pacote 'superheroapi'..."
cd src/github.com/carlsonsantana/superheroapi/
go build
echo "Pacote 'superheroapi' finalizado."

echo "Migrando o banco de dados..."
cd migrations
go run *.go migrate
cd ..
echo "Banco de dados migrado."

echo "Instalando pacote 'main'..."
cd main
go install
echo "Pacote 'main' instalado."

echo "Executanto aplicação"
go run github.com/carlsonsantana/superheroapi/main
