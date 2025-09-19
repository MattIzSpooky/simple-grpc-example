PROTO_DIR=proto
SQLC_DIR=queries
MIGRATIONS_DIR=migrations

all: generate build

# Run sqlc to regenerate Go DB code
sqlc:
	sqlc generate

# Run protoc to regenerate Go stubs
protoc:
	protoc --go_out=. --go-grpc_out=. \
		$(PROTO_DIR)/*.proto

# Regenerate everything
generate: sqlc protoc

# Build the server binary
build:
	go build -o server ./cmd/server

# Run migrations (example with psql, adapt as needed)
migrate:
	psql $$DATABASE_URL -f $(MIGRATIONS_DIR)/0001_create_notes.sql

# Run the app locally
run: build
	./server

# ------------------------------
# Docker / Compose targets
# ------------------------------

# Build docker image
docker-build:
	docker build -t notes-app .

# Run app + db with docker-compose
docker-up:
	docker-compose up --build

# Run detached (background)
docker-up-d:
	docker-compose up -d --build

# Stop services
docker-down:
	docker-compose down

# Stop and remove volumes
docker-clean:
	docker-compose down -v

# ------------------------------
clean:
	rm -f server
