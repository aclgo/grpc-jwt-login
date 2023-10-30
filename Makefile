up:
    sudo docker compose up --build -d
migrate:
    migrate -database postgres://grpc-jwt:grpc-jwt@localhost:5432/grpc-jwt?sslmode=disable -path migrate up 1
# migrate_down:
#     migrate -database postgres://grpc-jwt:grpc-jwt@localhost:5432/grpc-jwt?sslmode=disable -path migrations down 1
gorun:
    go run cmd/grpc-jwt/main.go
# protoc:
     protoc --go_out=. --go_opt=paths=source_relative \
     --go-grpc_out=. --go-grpc_opt=paths=source_relative \
     mail.proto