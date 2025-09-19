# How to generate and run

1. Install sqlc and protoc + protoc-gen-go + protoc-gen-go-grpc if you need the proto stubs.
2. Run migrations (e.g. with psql or a migration tool): ``psql $DATABASE_URL -f migrations/0001_create_notes.sql``
3. Generate sqlc code: ``sqlc generate``
4. Generate protobuf Go code (if you want gRPC): ``protoc --go_out=. --go-grpc_out=. proto/notes.proto``
5. Build & Run: ``go build ./cmd/server`` && ``DATABASE_URL=postgres://postgres:postgre``