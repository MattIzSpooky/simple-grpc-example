PROTO_DIR=proto
SQLC_DIR=queries
MIGRATIONS_DIR=migrations

all: generate build

# Run sqlc to regenerate Go DB code
sqlc:
	sqlc generate

# Run protoc to regenerate Go stubs
protoc:
	protoc --go_out=internal/pb --go-grpc_out=internal/pb \
		$(PROTO_DIR)/*.proto

# Regenerate everything
generate: sqlc protoc

# Build the server binary
build:
	go build -o server ./cmd/server

# Run migrations (example with psql, adapt as needed)
migrate:
	psql $$DATABASE_URL -f $(MIGRATIONS_DIR)/0001_create_notes.sql

# Run the app
run: build
	./server

clean:
	rm -f server
