# Projeto Transactions

## Tópicos Chave

- Conta com dados do cartão relacionado ao cliente
- Cada transação fica salvo o evento
- Tipos de transação (Compra a vista, compra parcelada, saque ou pagamento)
- Estrutura da transação (Tipo, Valor e Data de criação)
- Compras e saques tem valor negativo e pagamentos tem o valor positivo

## Executando Local com Docker

Execute o comando abaixo para subir com docker o projeto, ira disponibilizar uma API na porta :8080.

```sh
docker-compose up
```

Quando for parar o projeto basta executar o comando `ctrl+c` em seu terminal e caso tenha rodado com a flag `-d` poderá executar o comando:

```sh
docker-compose down
```

Lembrando que os dados do banco postgres estão armazenados em um volume gerenciado pelo docker, se não for mais utilizar o serviço nesta maquina não esqueça de desmontar o volume passando a flag `--volumes` no comando `down`.

## Executando Local sem Docker

Necessário possuir o GoLang configurado em sua maquina e uma instância do postgres.

```sh
go mod download
go run main.go
```

> Nota
>
> Algumas variaveis de ambiente são necessarias para rodar o app, Ex:
> PORT: "8080", DB_HOST: "db", DB_PORT: "5432", DB_USER: "postgres", DB_PASS: "postgres", DB_BASE: "transaction"

## Rodando os Testes

Necessário possuir o GoLang configurado em sua máquina.

```sh
go test -v ./...
```
