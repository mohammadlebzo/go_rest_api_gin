sudo docker compose up -d

cp .env.sample .env

go run src/main.go

sudo docker compose down
