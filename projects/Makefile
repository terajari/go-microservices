BROKER_BINARY=brokerApp
AUTH_BINARY=authApp

up:
	@echo "Starting broker"
	docker compose up -d
	@echo "Docker images started"

down:
	@echo "Stopping docker compose..."
	docker-compose down
	@echo "Done!"

build_broker:
	@echo "Building broker"
	cd ../broker-service && env GOOS=linux CGO_ENABLED=0 go build -o ${BROKER_BINARY} ./cmd/api
	@echo "Done!"

build_auth:
	@echo "Building auth"
	cd ../auth-service && env GOOS=linux CGO_ENABLED=0 go build -o ${AUTH_BINARY} ./cmd/api
	@echo "Done!"

up_build: build_broker build_auth
	@echo "Stopping docker images (if running...)"
	docker-compose down
	@echo "Building (when required) and starting docker images..."
	docker-compose up --build -d
	@echo "Docker images built and started!"

start_frontend:
	@echo "Starting frontend"
	cd ../frontend && npm run dev
	@echo "Frontend started"