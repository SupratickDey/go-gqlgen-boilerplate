#!/bin/bash

echo "Downloading packages needed ..."
go get	github.com/99designs/gqlgen
go get	github.com/go-chi/chi
go get	github.com/google/wire
go get	github.com/jinzhu/gorm
go get	github.com/joho/godotenv
go get	github.com/stretchr/testify
go get	github.com/vektah/dataloaden
go get	github.com/vektah/gqlparser/v2


echo "Setup initial git hooks ..."
cp pre-commit .git/hooks
chmod +x .git/hooks/pre-commit