# go-gin-backend-template

## Docker Database Setup
To run a database in a docker container execute this command in the terminal

### Postgres
```bash
docker run --name <container_name> -e POSTGRES_PASSWORD=mysecretpassword -p 5432:5432 -d postgres
```
