run:
	go run main.go

proto:
	protoc --go_out=. --go_opt=paths=source_relative \
	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
	pb/*.proto

clean:
	rm pb/*.pb.go

# grpcurl execution
login:
	grpcurl --plaintext -d '{"user": "username", "pass": "password"}' localhost:5013 UserService.Login

find:
	grpcurl --plaintext -d '{"user_id": "16678292763"}' localhost:5013 UserService.Find

update:
	grpcurl --plaintext -d '{"user_id": "16678292763", "user": "umardev500", "pass": "umardev500pass"}' localhost:5013 UserService.UpdateCreds
	