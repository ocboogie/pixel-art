# Server

## Migrations

```bash
$ export DATABASE=postgres://postgres:password@localhost:5432/postgres?sslmode=disable
$ migrate -path migrations -database "$DATABASE" command # (See help)
```
