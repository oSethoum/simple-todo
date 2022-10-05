SELF_DIR := $(dir $(lastword $(MAKEFILE_LIST)))

.DEFAULT_GOAL := help
NODEVER=v17.4.0
GOFLAGS=-mod=mod

#help: @ List available tasks
help:
	@clear
	@echo "Usage: make COMMAND"
	@echo "Commands :"
	@grep -E '[a-zA-Z\.\-]+:.*?@ .*$$' $(MAKEFILE_LIST)| tr -d '#' | awk 'BEGIN {FS = ":.*?@ "}; {printf "\033[32m%-16s\033[0m - %s\n", $$1, $$2}'

#deps: @ Install dependencies
deps:
	cd ./frontend && \
	curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.1/install.sh | bash && \
	. ~/.bashrc && \
	. ${NVM_DIR}/nvm.sh && \
	nvm install $(NODEVER) && \
	nvm use $(NODEVER) && \
	curl -fsSL https://get.pnpm.io/install.sh | sh - && \
	. ~/.bashrc && \
	go install github.com/99designs/gqlgen@latest

#install-frontend: @ Install frontend
install-frontend:
	cd ./frontend && \
	. ${NVM_DIR}/nvm.sh && \
	nvm use $(NODEVER) && \
	pnpm install

#build-frontend: @ Build frontend
build-frontend:
	cd ./frontend && \
	. ${NVM_DIR}/nvm.sh && \
	nvm use $(NODEVER) && \
 	pnpm run build

#run-frontend: @ Run frontend
run-frontend: install
	cd ./frontend && \
	. ${NVM_DIR}/nvm.sh && \
	nvm use $(NODEVER) && \
	pnpm run dev

#clean-backend: @ Cleanup backend
clean-backend:
	@cd ./backend && \
	sudo rm -rf ./vendor/

#generate-backend: @ Generate backend GraphQL go source code
generate-backend:
	@cd ./backend && \
	sudo rm -rf .graph/model && \
	sudo rm -rf .graph/generated && \
	export GOFLAGS=$(GOFLAGS); go run github.com/99designs/gqlgen generate

#test-backend: @ Run backend tests
test-backend: generate-backend
	@cd ./backend && \
	export GOFLAGS=$(GOFLAGS); go test -v ./...

#build-backend: @ Build backend GraphQL API
build-backend: generate-backend
	@cd ./backend && \
	export GOFLAGS=$(GOFLAGS); go build -o server server.go

#run-backend: @ Run backend GraphQL API
run-backend: generate-backend
	@cd ./backend && \
	export GOFLAGS=$(GOFLAGS); go run server.go

#get-backend: @ Download and install go backend packages
get-backend: clean-backend
	@cd ./backend && \
	export GOFLAGS=$(GOFLAGS); go get . ; go mod tidy

#update-backend: @ Update go backend dependencies to latest versions
update-backend: clean-backend
	@cd ./backend && \
	export GOFLAGS=$(GOFLAGS); go get -u; go mod tidy
