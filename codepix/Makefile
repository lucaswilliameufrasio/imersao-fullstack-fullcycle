up-sudo:
	sudo docker-compose up -d
.PHONY: up-sudo

down-sudo:
	sudo docker-compose down
.PHONY: down-sudo

attach-sudo:
	sudo docker exec -it codepix_app_1 bash
.PHONY: attach-sudo

up:
	docker-compose up -d
.PHONY: up

down:
	docker-compose down
.PHONY: down

attach:
	docker exec -it codepix_app_1 bash
.PHONY: attach

## Download all dependencies
get-deps:
	$ go get -d ./...
.PHONY: get-deps

evansr:
	$ evans -r repl
.PHONY: evansr

cc-prtfiles:
	protoc --go_out=app/grpc/pb --go_opt=paths=source_relative --go-grpc_out=app/grpc/pb --go-grpc_opt=paths=source_relative --proto_path=app/grpc/protofiles app/grpc/protofiles/*.proto
.PHONY: cc-prtfiles