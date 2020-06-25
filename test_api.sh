#!/bin/bash

set -e

echo "Construindo o pacote 'superheroapi'..."
cd src/github.com/carlsonsantana/superheroapi/
go build
echo "Pacote 'superheroapi' finalizado."

echo "Executanto testes"
cd tests
go test
