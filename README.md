# mini-be-services-ecommerce

Mini project of ecommerce specific cart service.

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

Start the application

```bash
docker-compose up -d
```

## How to test the application

here the json collections
