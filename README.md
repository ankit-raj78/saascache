# SaaSCache - Cache Provisioning Service

A Go-based HTTP service that provides cache recommendations and provisioning capabilities.

## Project Structure

```
saascache/
├── cmd/
│   └── server/          # main executable
├── internal/
│   ├── api/             # HTTP router & handlers
│   │   ├── router.go
│   │   └── handlers/
│   │       └── handlers.go
│   ├── config/          # configuration loader
│   │   └── config.go
│   ├── recommend/       # recommendation engine
│   │   └── engine.go
│   ├── provision/       # provisioning interfaces & stubs
│   │   ├── interface.go
│   │   └── local.go
│   └── integrate/       # integration formatter stub
│       └── formatter.go
├── pkg/                 # reusable libraries (if any)
├── db/
│   └── migrations/      # SQL migration files
├── .gitignore
└── README.md
```

## API Endpoints

### Health Check
- **GET** `/v1/health` - Returns server health status

### Recommendations
- **POST** `/v1/recommend` - Get cache recommendations (request-driven)

### Provisioning
- **POST** `/v1/provision` - Schedule cache provisioning (event-driven)
  - Request body: `{"engine":"redis","pattern":"cache-aside","nodeCount":1}`
  - Returns: `{"jobId":"job-xxxxxxxx"}`

## Running the Server

```bash
# Install dependencies
go get github.com/gin-gonic/gin

# Run the server
go run ./cmd/server

# The server will start on port 8080 by default
# Set PORT environment variable to use a different port
```

## Testing

```bash
# Health check
curl http://localhost:8080/v1/health

# Get recommendations
curl -X POST http://localhost:8080/v1/recommend \
  -H "Content-Type: application/json" \
  -d '{}'

# Schedule provisioning
curl -X POST http://localhost:8080/v1/provision \
  -H "Content-Type: application/json" \
  -d '{"engine":"redis","pattern":"cache-aside","nodeCount":1}'
```

## Next Steps

1. Replace `LocalProvisioner` with real AWS provisioning logic
2. Add PostgreSQL for job state persistence
3. Replace goroutine with SQS for event-driven provisioning
4. Implement recommendation engine logic
5. Add authentication and authorization
6. Add comprehensive testing
7. Add metrics and logging
