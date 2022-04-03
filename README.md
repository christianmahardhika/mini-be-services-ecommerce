# mini-be-services-ecommerce

Mini project of ecommerce specific cart service.

## **Description**

This application only contain cart, order service and light products service

the products service only show searched product(s)

What you can do in this app?

- Search Product(s)

- Add item to cart

- Get cart info

- Delete item from cart

- Reset the cart

- Checkout the cart

## **How to use?**

Makesure install Docker in your machine
<https://docs.docker.com/engine/install/>

clone this repo

```bash
git clone https://github.com/christianmahardhika/mini-be-services-ecommerce.git
```

create <.env> file and fill this variable

```text
DB_USER=postgres
DB_PASS=root
DB_PORT=5432
DB_HOST=localhost
DB_NAME=ecommerce
DB_SSL_MODE=disable
PORT=8080

ALLOWED_ORIGINS=["*"]
ALLOWED_HEADERS=["*"]
```

this application use makefile to run build, test, deploy, etc

Start the application

```bash
make docker-start
```




## How to test the application

here the json collections
<https://documenter.getpostman.com/view/4625022/UVysxFmE>

run testing command

```bash
make test
```
