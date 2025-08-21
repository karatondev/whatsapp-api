# WhatsApp API

A comprehensive WhatsApp messaging API built with Go that provides functionality for sending messages, managing contacts, and handling groups through both REST and gRPC interfaces.

## 🚀 Features

- **Message Management**: Send text, image, video, document, audio, and location messages
- **Contact Management**: Retrieve and manage WhatsApp contacts
- **Group Management**: Handle WhatsApp group operations
- **Multi-Protocol Support**: Both REST API and gRPC interfaces
- **Media Support**: Handle various media types with validation
- **Database Integration**: PostgreSQL for data persistence
- **Caching**: Redis for improved performance
- **Logging**: Structured logging with rotation
- **Configuration**: Flexible YAML-based configuration

## 📋 Prerequisites

- Go 1.23.3 or higher
- PostgreSQL database
- Redis server
- WhatsApp Core Gateway (gRPC service)

## 🛠️ Installation

1. **Clone the repository:**
   ```bash
   git clone <repository-url>
   cd whatsapp-api
   ```

2. **Install dependencies:**
   ```bash
   go mod download
   ```

3. **Configure the application:**
   - Copy and modify `config.yaml` according to your environment
   - Set up your database connection
   - Configure Redis connection
   - Set gRPC service endpoint

4. **Run the application:**
   ```bash
   go run cmd/main.go
   ```

## ⚙️ Configuration

The application uses `config.yaml` for configuration. Key sections include:

### Server Configuration
```yaml
server:
  port: 8001
  mode: "debug"
  shutdown_timeout: 5s
```

### Database Configuration
```yaml
postgres:
  host: localhost
  port: 5432
  user: your_user
  password: your_password
  dbname: whatsapp_api
```

### Redis Configuration
```yaml
redis:
  host: 127.0.0.1
  port: 6379
  database: 0
```

### gRPC Configuration
```yaml
grpc:
  dsn: "localhost:50051"
```

## 📚 API Documentation

### REST Endpoints

#### Send Message
**POST** `/api/v1/messages`

Send various types of messages including text, media, and location.

**Request Body:**
```json
{
  "device": "device_id",
  "to": "6281234567890",
  "type": "text|image|video|document|audio|location",
  "text": "Hello World",
  "image": {
    "url": "https://example.com/image.jpg",
    "caption": "Image caption",
    "mimetype": "image/jpeg"
  },
  "video": {
    "url": "https://example.com/video.mp4",
    "caption": "Video caption",
    "mimetype": "video/mp4"
  },
  "document": {
    "url": "https://example.com/document.pdf",
    "caption": "Document caption",
    "mimetype": "application/pdf",
    "title": "Document Title"
  },
  "audio": {
    "url": "https://example.com/audio.ogg",
    "ptt": false,
    "mimetype": "audio/ogg"
  },
  "location": {
    "longitude": 116.7169678,
    "latitude": -0.9727893,
    "name": "Location Name",
    "address": "Location Address"
  }
}
```

**Response:**
```json
{
  "message": {
    "id": "message_id"
  }
}
```

#### Get Contacts
**GET** `/api/v1/contacts/:id`

Retrieve contacts for a specific device.

**Response:**
```json
{
  "device_id": "device_id",
  "device_alias": "Device Alias",
  "device_name": "Device Name",
  "connect_status": "connected",
  "contacts": [
    {
      "name": "Contact Name",
      "phone": "6281234567890",
      "short": "Short Name"
    }
  ]
}
```

#### Get Groups
**GET** `/api/v1/groups/:id`

Retrieve groups for a specific device.

**Response:**
```json
{
  "device_id": "device_id",
  "device_alias": "Device Alias",
  "device_name": "Device Name",
  "groups": [
    {
      "name": "Group Name",
      "phone": "group_id",
      "short": "Group Short Name"
    }
  ]
}
```

### Message Types

The API supports the following message types:

1. **Text Messages**: Simple text messages
2. **Image Messages**: Images with optional captions
3. **Video Messages**: Videos with optional captions
4. **Document Messages**: Files with metadata
5. **Audio Messages**: Audio files with PTT support
6. **Location Messages**: Geographic coordinates with details

### Validation Rules

When sending media messages, all fields within the media object are required:

- **Image**: `url`, `caption`, `mimetype`
- **Video**: `url`, `caption`, `mimetype`
- **Document**: `url`, `filename`, `mimetype`, `title`
- **Audio**: `url`, `mimetype` (ptt is optional)
- **Location**: `latitude`, `longitude`, `name`, `address`

## 🏗️ Project Structure

```
whatsapp-api/
├── cmd/
│   └── main.go                 # Application entry point
├── internal/
│   ├── app/
│   │   └── app.go             # Application initialization
│   ├── handler/
│   │   ├── grpc.go            # gRPC handlers
│   │   ├── middleware/        # Middleware components
│   │   └── rest/              # REST API handlers
│   ├── provider/
│   │   ├── logger.go          # Logging provider
│   │   ├── postgres.go        # PostgreSQL provider
│   │   ├── redis.go           # Redis provider
│   │   └── dailylogger/       # Daily log rotation
│   ├── repository/
│   │   ├── repository.go      # Repository interface
│   │   └── public_repo.go     # Repository implementation
│   └── service/
│       ├── service.go         # Service interface
│       ├── messages.go        # Message service implementation
│       └── service_error.go   # Service error definitions
├── model/
│   ├── entity/                # Database entities
│   ├── constant/              # Application constants
│   ├── pb/                    # Protocol buffer generated files
│   └── *.go                   # Request/Response models
├── util/
│   ├── config.go             # Configuration utilities
│   └── helpers.go            # Helper functions
├── log/                      # Log files directory
├── config.yaml              # Application configuration
├── go.mod                   # Go module definition
└── README.md               # This file
```

## 🔧 Development

### Building the Application
```bash
go build -o whatsapp-api cmd/main.go
```

### Running Tests
```bash
go test ./...
```

### Code Formatting
```bash
go fmt ./...
```

### Linting
```bash
golangci-lint run
```

## 🐳 Docker Support

Create a `Dockerfile`:
```dockerfile
FROM golang:1.23.3-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o whatsapp-api cmd/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=builder /app/whatsapp-api .
COPY --from=builder /app/config.yaml .

CMD ["./whatsapp-api"]
```

### Docker Compose
```yaml
version: '3.8'
services:
  whatsapp-api:
    build: .
    ports:
      - "8001:8001"
    depends_on:
      - postgres
      - redis
    environment:
      - CONFIG_PATH=./config.yaml
    
  postgres:
    image: postgres:15
    environment:
      POSTGRES_DB: whatsapp_api
      POSTGRES_USER: user
      POSTGRES_PASSWORD: password
    ports:
      - "5432:5432"
    
  redis:
    image: redis:7-alpine
    ports:
      - "6379:6379"
```

## 📊 Monitoring and Logging

The application includes comprehensive logging with:

- **Structured Logging**: Using logrus with JSON formatting
- **Log Rotation**: Daily log rotation with configurable retention
- **Log Levels**: Debug, Info, Warn, Error levels
- **Request Logging**: HTTP request/response logging
- **Error Tracking**: Detailed error logging with stack traces

Log files are stored in the `log/` directory:
- `log/info/` - Information and debug logs
- `log/error/` - Error logs

## 🛡️ Error Handling

The application implements comprehensive error handling:

- **Validation Errors**: Input validation with detailed error messages
- **Service Errors**: Business logic error handling
- **Database Errors**: Database operation error handling
- **gRPC Errors**: gRPC communication error handling

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🆘 Support

For support and questions:

- Create an issue in the repository
- Check the documentation
- Review the logs for error details

## 🔄 Changelog

### v1.0.0
- Initial release
- Basic messaging functionality
- Contact and group management
- REST and gRPC interfaces
- Media message support with validation
- Comprehensive error handling
