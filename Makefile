protogen:
	protoc  --proto_path=api --go_out=pkg  \
	--go-grpc_out=pkg  \
	--grpc-gateway_out=pkg \
	api/*.proto


userdb:
	docker run --name userdb  -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15-alpine

createuserdb:
	docker exec -it userdb createdb --username=root --owner=root user

dropuserdb:
	docker exec -it userdb dropdb user




test:
	go test ./...