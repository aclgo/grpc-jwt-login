gpr:
	protoc --go_out=. --go_opt=paths=./proto \
    --go-grpc_out=. --go-grpc_opt=paths=./proto \
    ./proto/user.proto