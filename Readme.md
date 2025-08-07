## initialize go module
go mod init github.com/sanukumar/go-students-api

## folder structure
create cmd folder --> go-students-api 

## run go prooject
go run --> (main file path)

## add git ignore file
1. add git ignore vs code extension
2. cmd+shif+p -> .gitignore -> go

## git initialize
git init
add commit -> (initial commit)

## DB
sqlite -> file base database --> fast

### https://github.com/ilyakaznacheev/cleanenv
to install package --> go get -u github.com/ilyakaznacheev/cleanenv

## to run the application
go run cmd/students-api/main.go -config config/local.yaml

## install request validator
https://github.com/go-playground/validator

go get github.com/go-playground/validator/v10

## install squlite db drive
https://github.com/mattn/go-sqlite3

go get github.com/mattn/go-sqlite3