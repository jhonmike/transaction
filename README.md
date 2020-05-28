# Projeto Transactions

## Topicos chave

- Conta com dados do cartão relacionado ao cliente
- Cada transação fica salvo o evento
- Tipos de transação (Compra a vista, compra parcelada, saque ou pagamento)
- Estrutura da transação (Tipo, Valor e Data de criação)
- Compras e saques tem valor negativo e pagametnos valor positivo

## Executando

```sh
go mod download
go run main.go
```

## Rodando os testes

```sh
go test -v ./...
```
