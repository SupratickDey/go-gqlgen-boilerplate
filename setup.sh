#!/bin/bash

echo "Downloading packages needed ..."
go install	github.com/99designs/gqlgen
go install	github.com/go-chi/chi
go install	github.com/google/wire
go install	github.com/jinzhu/gorm
go install	github.com/joho/godotenv
go install	github.com/stretchr/testify
go install	github.com/vektah/dataloaden
go install	github.com/vektah/gqlparser/v2


echo "Setup initial git hooks ..."
cp pre-commit .git/hooks
chmod +x .git/hooks/pre-commit