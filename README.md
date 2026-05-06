# Go Microservice Template

This is a comprehensive template for building scalable Go microservices using clean architecture principles. It provides a solid foundation with gRPC, HTTP APIs, Kafka integration, Redis caching, observability, and deployment configurations.

## Features

- **Clean Architecture**: Organized into layers (delivery, usecase, repository, model) for maintainability
- **Multiple Protocols**: Support for both gRPC and HTTP APIs
- **Message Queue**: Kafka integration for event-driven architecture
- **Caching**: Redis client for high-performance caching
- **Observability**: Prometheus metrics and distributed tracing
- **Database**: Migration support for database schema management
- **Deployment**: Docker and Kubernetes configurations
- **Logging**: Structured logging with configurable levels
- **Graceful Shutdown**: Proper handling of application shutdown

## Project Structure

```
go-template-code/
├── api/                    # Generated API code (protobuf)
│   └── gen/
│   |    └── server/
│   |        └── v1/
│   └── proto/
│        └── server/
│            └── v1/
├── cmd/                    # Application entry points
│   ├── infra-init/        # Infrastructure initialization
│   ├── server/            # Main server application
│   └── worker/            # Background worker processes
├── configs/                # Configuration files
│   └── config.yaml        # Application configuration
├── deployments/            # Deployment configurations
│   ├── docker/            # Docker setup
│   │   ├── docker-compose.yml
│   │   └── Dockerfile
│   └── k8s/               # Kubernetes manifests
│       └── deployment.yaml
├── docs/                   # Documentation
│   ├── ARCHITECTURE.md    # Architecture documentation
│   ├── docs.go            # Go documentation
│   ├── swagger.json       # API documentation
│   └── swagger.yaml
├── internal/               # Private application code
│   └── server/
│       ├── config/        # Configuration management
│       ├── constant/      # Application constants
│       │   ├── error.go   # Error definitions
│       │   └── status.go  # Status codes
│       ├── delivery/      # Delivery layer (handlers)
│       │   ├── grpc/      # gRPC handlers
│       │   └── http/      # HTTP handlers
│       ├── infra/         # Infrastructure code
│       │   └── migrate/   # Database migrations
│       ├── model/         # Data models
│       ├── repository/    # Data access layer
│       └── usecase/       # Business logic layer
├── migrations/             # Database migration files
│   ├── 001_init_db.down.sql
│   └── 001_init_db.up.sql
├── pkg/                    # Shared packages
│   ├── kafka/             # Kafka client and producer
│   ├── logger/            # Logging utilities
│   ├── metric/            # Prometheus metrics
│   ├── observability/     # Tracing and monitoring
│   ├── redis/             # Redis client
│   ├── shutdown/          # Graceful shutdown utilities
│   └── utils/             # General utilities
└── scripts/                # Build and deployment scripts
    └── scripts.sh
```

### Directory Explanations

- **api/**: Contains generated code from protobuf definitions for API contracts
- **cmd/**: Main applications that can be executed
  - `infra-init`: Initializes infrastructure components (databases, queues, etc.)
  - `server`: The main API server handling requests
  - `worker`: Background processes for async tasks
- **configs/**: Configuration files for different environments
- **deployments/**: Containerization and orchestration configurations
- **docs/**: Project documentation and API specifications
- **internal/**: Private application code following Go conventions
  - `config/`: Configuration loading and validation
  - `constant/`: Application-wide constants and enums
  - `delivery/`: API handlers and controllers
  - `infra/`: Infrastructure concerns (migrations, connections)
  - `model/`: Domain models and data structures
  - `repository/`: Data persistence layer
  - `usecase/`: Business logic and use cases
- **migrations/**: Database schema migration scripts
- **pkg/**: Reusable packages that can be imported by other projects
- **scripts/**: Automation scripts for build, test, and deployment

## Prerequisites

- Go 1.22+
- Docker & Docker Compose
- Kubernetes (for production deployment)
- Kafka (if using message queues)
- Redis (if using caching)
- PostgreSQL/MySQL (depending on database choice)

## Getting Started

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd go-template-code
   ```

2. **Install dependencies**
   ```bash
   go mod tidy
   ```

3. **Configure the application**
   - Copy `configs/config.yaml` and update with your settings
   - Set environment variables as needed

4. **Run infrastructure**
   ```bash
   # Using Docker Compose
   cd deployments/docker
   docker-compose up -d
   ```

5. **Run the application**
   ```bash
   # Run server
   go run cmd/server/main.go

   # Run worker
   go run cmd/worker/main.go

   # Initialize infrastructure
   go run cmd/infra-init/main.go
   ```

## Development

### Building

```bash
make build
# or
go build ./cmd/server
```

### Testing

```bash
make test
# or
go test ./...
```

### Code Generation

```bash
# Generate protobuf code
make proto

# Generate mocks (if using)
make mocks
```

## Deployment

### Docker

```bash
# Build image
docker build -f deployments/docker/Dockerfile -t my-service .

# Run with docker-compose
cd deployments/docker
docker-compose up
```

### Kubernetes

```bash
# Apply manifests
kubectl apply -f deployments/k8s/
```

## Configuration

The application uses a YAML configuration file located at `configs/config.yaml`. Key configuration sections include:

- Database connection settings
- Kafka broker configuration
- Redis connection details
- Server ports and timeouts
- Logging levels

## API Documentation

- **Swagger UI**: Available at `/swagger/` endpoint when server is running
- **gRPC Reflection**: Enabled for gRPC clients
- **API Specs**: Located in `docs/` directory

## Monitoring & Observability

- **Metrics**: Exposed via `/metrics` endpoint (Prometheus format)
- **Tracing**: Distributed tracing with OpenTelemetry
- **Logging**: Structured JSON logging with configurable levels

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests
5. Submit a pull request

## License

This project is licensed under the MIT License - see the LICENSE file for details.
