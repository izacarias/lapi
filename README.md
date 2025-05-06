# Location API (LAPI)

A partial implementation of ETSI GS MEC 013 Location API specification. 
This API provides location services for Mobile Edge Computing applications.

## Features

- Zone management and queries
- Access point information
- User location tracking
- Distance calculations
- Swagger documentation
- RESTful API endpoints

## Prerequisites

- Go 1.23.3 or higher
- MongoDB (v7.0)
- Docker (optional)
- Docker Compose (optional)

## Quick Start with Docker Compose

The easiest way to run the application is using Docker Compose:

```bash
docker compose up -d
```

his will start both the API service and MongoDB. The API will be available at:
- API: http://localhost:8080
- Ping Endpoint: http://localhost:8080/ping
- Swagger UI: http://localhost:8080/docs

## Running Locally

### 1. Environment Setup

Create a `.env` file in the project root:

```env
MONGOURI=mongodb://localhost:27017/lapi
APIROOT=http://localhost:8080
APIVERSION=v3
PURGEDATABASE=1
INSERTMOCKDATA=1
```
### 2. Start MongoDB

Using Docker:
```bash
docker run --name mongodb -p 27017:27017 -d mongodb/mongodb-community-server:latest
```

Or using Podman:
```bash
podman run --name mongodb --privileged -p 27017:27017 -d docker.io/mongodb/mongodb-community-server:latest
```

or if is was already initialized

```bash
docker start mongodb
```

or 

```bash
podman start mongodb
```

### 3. Run the Application

```bash
go run main.go
```

## API Documentation

The API documentation is available through Swagger UI at `/docs` endpoint. You can also find the OpenAPI specification at `/swagger/doc.json`.

> [!Note]
> Depending on your system, it might require to run podman with root privilleges (e.g., Fedora)
> This is due to a requirement of MongoDB and the rights to access files in folders. Without the 
> `sudo` command, SElinux prevents the MondoDB container to write in the filesystem causing an error.

## Testing

### Integration Tests

The integration tests assume the MongoDB instance is running an is accessible. 
Please keep in mind that the database will be recreated and filled with mock data 
(make sure that the options `PURGEDATABASE` and  `INSERTMOCKDATA` are enabled in the [.env](.env) file).

Change to the project directory and run the tests as following:

```bash
go test -v
```
### Manual Testing

You can test the API endpoints using:

1. The provided REST Client file in `./tests/Requests.rest`
2. Swagger UI at `/docs`
3. Any HTTP client (like Postman or cURL)

## API Endpoints

- `GET /ping` - Health check
- `GET /location/v3/queries/zones` - List all zones
- `GET /location/v3/queries/zones/{zoneId}` - Get zone by ID
- `GET /location/v3/queries/zones/{zoneId}/accessPoints` - List access points in a zone
- `GET /location/v3/queries/users` - Query users with various filters
- `GET /location/v3/queries/distance` - Calculate distances between points

## Development

### Update Swagger Documentation

```bash
swag init -g main.go --output docs/swagger
```

### Project Structure

```
.
├── configs/         # Configuration and database setup
├── docs/           # Swagger documentation
├── handlers/       # HTTP request handlers
├── models/         # Data models
├── repositories/   # Data access layer
├── routes/         # Route definitions
├── tests/         # Test files
└── docker/        # Docker-related files
```

## License

This project is licensed under the Apache 2.0 License - see the [LICENSE](LICENSE) file for details.

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## Contact

Iulisloi Zacarias - [@izacarias](https://github.com/izacarias)

Project Link: [https://github.com/izacarias/lapi](https://github.com/izacarias/lapi)