generate-proto:
	mkdir -p ./pkg/api
	protoc \
		--go_out=./pkg \
		--go_opt=paths=source_relative \
		--go-grpc_out=./pkg \
		--go-grpc_opt=paths=source_relative \
		./api/info.proto

generate-ent:
	mkdir -p ./pkg/ent
	cp -R ./schema ./pkg/ent/schema
	go run -mod=mod entgo.io/ent/cmd/ent generate ./pkg/ent/schema
# rm -r ./pkg/ent/schema

generate-ts:
	mkdir -p ./npm/code
	protoc -I=./api --js_out=import_style=commonjs:./npm/code --grpc-web_out=import_style=typescript,mode=grpcwebtext:./npm/code ./api/info.proto