#!/bin/bash

set -e

echo "Construindo o pacote 'superheroapi'..."
cd src/github.com/carlsonsantana/superheroapi/
go build
echo "Pacote 'superheroapi' finalizado."

echo "Instalando pacote 'main'..."
cd main
go install
echo "Pacote 'main' instalado."

echo "Executanto aplicação"
go run github.com/carlsonsantana/superheroapi/main
