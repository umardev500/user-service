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
	grpcurl --plaintext -d '{"user_id": "16678292763s"}' localhost:5013 UserService.Find

update:
	grpcurl --plaintext -d '{"user_id": "16678292763", "user": "umardev500", "pass": "umardev500pass"}' localhost:5013 UserService.UpdateCreds

updateDetail:
	grpcurl --plaintext -d '{"user_id": "16678292763", "detail": {"name": "Mackenzie Shawley Flash", "email": "mack@gmail.com", "phone": "+62 8387915430", "location": {"address": "Kp. Cimedang","village": "Menes", "district": "Menes", "city": "Pandeglang", "province": "Banten", "postal_code": "42265"}}}' localhost:5013 UserService.UpdateDetail
	