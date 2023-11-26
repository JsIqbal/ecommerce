serve:
	@echo "Choose how to run the project:"
	@echo "1. Docker"
	@echo "2. Local"
	@read -p "Enter your choice (1/2): " choice; \
	if [ $$choice -eq 1 ]; then \
		docker-compose up --build; \
	elif [ $$choice -eq 2 ]; then \
		go run main.go serve-rest; \
	else \
		echo "Invalid choice"; \
	fi


seed:
	go run main.go seed

test:
	CGO_ENABLED=1 go test -gcflags=-l -cover -race ./...

perfect_go_get:
	go get package_module@commit_hash

deploy_redis:
	docker run --name redis -p 6379:6379 redis redis-server --requirepass "12345"