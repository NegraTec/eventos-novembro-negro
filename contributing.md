## Usando

- docker
- Go

> gock: Simula http requests

> testify: para os asserts nos testes

## Comandos

`docker build -t eventos-negros .`

`docker run --rm -p 8080:8080 -v "$PWD":/go/src/app -w /go/src/app -e FACEBOOK_ACCESS_TOKEN -it eventos-negros go run novembro.go`
