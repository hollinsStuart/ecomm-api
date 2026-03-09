# Ecomm

Following the [Ecomm API Tutorial](https://github.com/sikozonpc/ecom-go-api-project), with additional features.

We are building

- GET /health
- GET /products?name&limit=20&offset=0
- POST /orders

## TODOs

- [ ] All user apis
- [ ] GET /orders/{id}
- [ ] POST /product
- [ ] GET /products/{id}

## Placing an order

```json
{
  "customerId": 123,
  "items": [
    {"product_id":  42, "quantity":  2}
  ]
}
```