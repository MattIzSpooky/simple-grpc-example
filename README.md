# Notes Service (Go + gRPC + HTTP + Postgres)

A simple Notes service implemented in Go, using Postgres with `sqlc`, exposing both JSON HTTP and gRPC APIs.

---

## Getting Started

### Prerequisites

- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/)
- (optional) [Make](https://www.gnu.org/software/make/)

### Run with Docker Compose

```
make docker-up
```

This will:

* Start Postgres (`db`)
* Build and run the Notes app (`app`)
* Expose:

    * HTTP → [http://localhost:8080](http://localhost:8080)
    * gRPC → localhost:9090

### Tear Down

```bash
make docker-down
```

Clean everything (including Postgres volume):

```bash
make docker-clean
```

---

## HTTP API (cURL examples)

1. Create a Note

```bash
curl -X POST http://localhost:8080/notes \
  -H "Content-Type: application/json" \
  -d '{"description": "My first note"}'
```

2. Get All Notes

```bash
curl http://localhost:8080/notes
```

3. Update a Note

```bash
curl -X PUT http://localhost:8080/notes/<NOTE_ID> \
  -H "Content-Type: application/json" \
  -d '{"description": "Updated note text"}'
```

4. Delete a Note

```bash
curl -X DELETE http://localhost:8080/notes/<NOTE_ID>
```

---

## gRPC API

### Generate Protobuf Stubs

```bash
make protoc
```

### Test with grpcurl

List services:

```bash
grpcurl -plaintext localhost:9090 list
```

Create a Note:

```bash
grpcurl -plaintext -d '{"description":"My first gRPC note"}' localhost:9090 notes.NotesService/CreateNote
```

Get All Notes:

```bash
grpcurl -plaintext -d '{}' localhost:9090 notes.NotesService/GetAllNotes
```

Update a Note:

```bash
grpcurl -plaintext -d '{"id":"<NOTE_ID>", "description":"Updated text"}' localhost:9090 notes.NotesService/UpdateNote
```

Delete a Note:

```bash
grpcurl -plaintext -d '{"id":"<NOTE_ID>"}' localhost:9090 notes.NotesService/DeleteNote
```

---

## Development

Run the app locally without Docker:

```bash
make run
```

Regenerate SQLC + Protobuf:

```bash
make generate
```

---
