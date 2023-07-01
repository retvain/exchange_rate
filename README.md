## Migrations

docker run -v /home/retvain/dev/go/exchange_rate/migrations:/migrations --network host migrate/migrate -path=/migrations/ -database "postgresql://postgres:postgres@localhost:54324/exchange_rate?sslmode=disable" up