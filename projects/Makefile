BROKER_BINARY=brokerApp

up:
	@echo "Starting broker"
	docker compose up -d
	@echo "Broker started"

down:
	@echo "Stopping docker compose..."
	docker-compose down
	@echo "Done!"

build_broker:
	@echo "Building broker"
	cd ../broker-service && env GOOS=linux CGO_ENABLED=0 go build -o ${BROKER_BINARY} ./cmd/api
	@echo "Done!"

up_build: build_broker
	@echo "Stopping docker images (if running...)"
	docker-compose down
	@echo "Building (when required) and starting docker images..."
	docker-compose up --build -d
	@echo "Docker images built and started!"

start_frontend:
	@echo "Starting frontend"
	cd ../frontend && npm run dev
	@echo "Frontend started"