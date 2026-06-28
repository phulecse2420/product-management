# Product Management API

A simple REST API built with Go (Gin + PostgreSQL) to manage a product inventory.

---

## Tech stack

- **Go** 1.26.4
- **Gin** — HTTP framework
- **PostgreSQL** — database
- **godotenv** — environment config

---

## Project structure

```
product-management/
├── main.go
├── .env
├── config/
│   └── config.go
├── internal/
│   ├── app/app.go
│   ├── routes/routes.go
│   ├── handlers/product_handler.go
│   ├── services/product_service.go
│   ├── repositories/product_repository.go
│   ├── models/product_model.go
│   └── infrastructure/database.go
└── migrations/
    └── 001_create_products_table.sql
```

---

## How to run

### 1. Prerequisites

- Go 1.26.4
- [Docker](https://docs.docker.com/get-docker/) and [Docker Compose](https://docs.docker.com/compose/install/)

### 2. Clone and install dependencies

```bash
git clone git@github.com:phulecse2420/product-management.git
cd product-management
go mod tidy
```


### 2. Configure environment

Create a `.env` file at the project root:

```env
DB_HOST=db
DB_PORT=5432
DB_USER=postgres_user
DB_PASSWORD=postgres_password
DB_NAME=product_management
PORT=8080
```

> **Note:** `DB_HOST` must be `db` (the Docker Compose service name), not `localhost`.

### 3. Start all services

```bash
docker compose up --build
```

This will:
- Start a PostgreSQL container (`product_management_db`) on port `5432`
- Build and start the API container on port `8080`
- The API waits for the database to pass its health check before starting
### 4. Run database migration

In a separate terminal, after the containers are up:

```bash
docker compose exec -T db psql -U postgres_user -d product_management \
  < migrations/001_create_products_table.sql
```

The server is available at `http://localhost:8080`.

### Useful commands

```bash
# Run in background
docker compose up --build -d
```

```bash
# View logs
docker compose logs -f api
```

```bash
# Stop all containers
docker compose down
```

```bash
# Stop and remove volumes (resets the database)
docker compose down -v
```
 
---

## API endpoints

| Method   | Path                 | Description             |
|----------|----------------------|-------------------------|
| `Get`    | `/health`            | Check health app        |
| `POST`   | `/products`          | Create a new product    |
| `GET`    | `/products`          | List all products       |
| `GET`    | `/products?keyword=` | Search products by name |
| `GET`    | `/products/:id`      | Get product by ID       |
| `PUT`    | `/products/:id`      | Update a product        |
| `DELETE` | `/products/:id`      | Delete a product        |

---

## Example curl commands


### Check health

```bash
curl http://localhost:8080/health
```


### Create a product

```bash
curl -X POST http://localhost:8080/products \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Mechanical Keyboard",
    "description": "Wireless mechanical keyboard",
    "price": 120.50,
    "quantity": 10
  }'
```

**Response `201`:**

```json
{
  "id": 1,
  "name": "Mechanical Keyboard",
  "description": "Wireless mechanical keyboard",
  "price": 120.50,
  "quantity": 10,
  "created_at": "2026-06-21T10:00:00Z",
  "updated_at": "2026-06-21T10:00:00Z"
}
```

---

### List all products

```bash
curl http://localhost:8080/products
```

### Search by keyword

```bash
curl http://localhost:8080/products?keyword=keyboard
```

**Response `200`:** array of products (empty array `[]` if none found).

---

### Get product by ID

```bash
curl http://localhost:8080/products/1
```

**Response `200`:** single product object.

**Response `404` (not found):**

```json
{ "message": "product not found" }
```

---

### Update a product

```bash
curl -X PUT http://localhost:8080/products/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Updated Keyboard",
    "description": "Updated description",
    "price": 135.00,
    "quantity": 15
  }'
```

**Response `200`:** updated product object with new `updated_at`.

---

### Delete a product

```bash
curl -X DELETE http://localhost:8080/products/1
```

**Response `200`:**

```json
{ "message": "product deleted successfully" }
```

---

## Validation rules

| Field      | Rules                            |
|------------|----------------------------------|
| `name`     | Required, minimum 3 characters   |
| `price`    | Required, must be greater than 0 |
| `quantity` | Required, must be ≥ 0            |

**Example validation error response `400`:**

```json
{ "message": "name must be at least 3 characters" }
```

---

## Common errors

| Symptom                              | Likely cause                                | Fix                                                        |
|--------------------------------------|---------------------------------------------|------------------------------------------------------------|
| `failed to connect to database`      | PostgreSQL not running or wrong credentials | Check `.env` and ensure PostgreSQL is up                   |
| `relation "products" does not exist` | Migration not run                           | Run `psql ... -f migrations/001_create_products_table.sql` |
| `address already in use`             | Port 8080 taken                             | Change `PORT` in `.env` or stop the other process          |
| `404 product not found`              | ID does not exist in DB                     | Check with `GET /products` first                           |