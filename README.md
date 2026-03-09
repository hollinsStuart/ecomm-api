# Ecomm

We are building

- GET /health
- GET /products?name&limit=20&offset=0
- POST /orders

## TODOs

- GET /orders/{id}
- POST /product
- GET /products/{id}

## Placing an order

```json
{
  "customerId": 123,
  "items": [
    {"product_id":  42, "quantity":  2}
  ]
}
```