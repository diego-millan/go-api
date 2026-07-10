# go-api

REST API completa com Go, Gin e PostgreSQL.

## Stack

- **Go** 1.26
- **Gin** — HTTP web framework
- **PostgreSQL** 12 — banco de dados
- **Docker Compose** — container do banco

## Estrutura do projeto

```
go-api/
  cmd/main.go                    -- entrypoint
  controller/
    product_controller.go        -- handlers HTTP
  db/
    connection.go                -- conexão com o banco
    init.sql                     -- script de criação da tabela
  model/
    product.go                   -- struct Product
  repository/
    product_repository.go        -- queries SQL
  usecase/
    products_usecase.go          -- lógica de negócio
  docker-compose.yml             -- container PostgreSQL
```

## Rotas

| Método | Rota        | Descrição          |
|--------|------------|--------------------|
| GET    | `/ping`    | Health check       |
| GET    | `/products` | Listar produtos    |
| POST   | `/product`  | Criar um produto   |

### Exemplos

```sh
# health check
curl http://localhost:8001/ping

# listar produtos
curl http://localhost:8001/products

# criar um produto
curl -X POST http://localhost:8001/product \
  -H "Content-Type: application/json" \
  -d '{"name": "Notebook", "price": 3500.00}'
```

## Comandos do tutorial

### Inicializar o projeto

```sh
go mod init go-api
```

### Subir o banco com Docker

```sh
docker compose up -d
```

### Criar a tabela no banco

```sh
docker exec -i go_db psql -U postgres -d postgres < db/init.sql
```

### Instalar dependências

```sh
go mod tidy
```

### Rodar a aplicação

```sh
go run cmd/main.go
```


---

Tutorial: [Como criar uma REST API completa do zero com GO | Golang tutorial - iniciante](https://youtu.be/3p4mpId_ZU8) por Go Lab Tutoriais.
