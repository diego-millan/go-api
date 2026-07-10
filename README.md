# go-api

REST API completa com Go, Gin e PostgreSQL.

## Stack

- **Go** 1.26
- **Gin** — HTTP web framework
- **PostgreSQL** 12 — banco de dados
- **Docker Compose** — containers da aplicação e banco

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
  Dockerfile                     -- build da imagem
  docker-compose.yml             -- containers app + banco
```

## Rotas

| Método | Rota                  | Descrição               |
|--------|-----------------------|-------------------------|
| GET    | `/ping`               | Health check            |
| GET    | `/products`           | Listar produtos         |
| POST   | `/product`            | Criar um produto        |
| GET    | `/product/:productId` | Buscar produto por ID   |

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

# buscar produto por id
curl http://localhost:8001/product/1
```

## Comandos do tutorial

### Inicializar o projeto

```sh
go mod init go-api
```

### Instalar dependências

```sh
go mod tidy
```

### Subir app e banco com Docker Compose

```sh
docker compose up -d
```

> O `docker-compose.yml` sobe dois containers: `go-app` (a API) e `go_db` (PostgreSQL).
> A aplicação usa a variável `DB_HOST=go_db` para se conectar ao banco na rede Docker.

### Criar a tabela no banco

```sh
docker exec -i go_db psql -U postgres -d postgres < db/init.sql
```

### Rodar a aplicação localmente (sem Docker)

```sh
go run cmd/main.go
```

> Quando roda localmente, `DB_HOST` não está definida e o fallback é `localhost`.

### Build da imagem Docker

```sh
docker build -t go-api-tutorial .
```

### Verificar containers e imagens

```sh
# containers em execução
docker ps

# imagens disponíveis
docker images
```


---

Tutorial: [Como criar uma REST API completa do zero com GO | Golang tutorial - iniciante](https://youtu.be/3p4mpId_ZU8) por Go Lab Tutoriais.
