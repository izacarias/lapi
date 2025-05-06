# Location API (LAPI)

A partial implementation of ETSI GS MEC 013 Location API specification. 
This API provides location services for Mobile Edge Computing applications.

![GitHub last commit](https://img.shields.io/github/last-commit/izacarias/lapi)
![GitHub top language](https://img.shields.io/github/languages/top/izacarias/lapi)
![MongoDB](https://img.shields.io/badge/-MongoDB-4DB33D?style=flat&logo=mongodb&logoColor=white)
![Go](https://img.shields.io/badge/-Go-00ADD8?style=flat&logo=go&logoColor=white)
![Docker Compose](https://img.shields.io/badge/Docker%20Compose-2496ED?style=flat&logo=docker&logoColor=white)

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

## Development

### Update Swagger Documentation

```bash
swag init -g main.go --output docs/swagger
```

### Project Structure

```
.
├── configs/            # Utility functions to get configuration values
├── controllers         # Controllers to receive and route the requests
├── docs/               # Swagger generated documentation files
├── domain/             # Entitites and data models
├── mock/               # Functions to insert mock data into database
├── responses/          # Definition of objects used in responses as JSON
├── routes/             # Route definitions
├── services/           # Coupling between domain entities (WIP!)
├── tests/              # Test files
├── utils/              # Utility functions
...
├── docker-compose.yml  # Docker-related files
├── Dockerfile          # Docker-related files

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

Iulisloi Zacarias 
- [Personal page at TU Braunschweig](https://www.tu-braunschweig.de/en/kns/faculty-and-staff/iulisloi-zacarias)
- [@izacarias at GitHub](https://github.com/izacarias)

Project Link: [https://github.com/izacarias/lapi](https://github.com/izacarias/lapi)

## Acknowledgment

This work was supported by the German Federal Ministry of Education and Research (BMBF) project [6G-ANNA](https://6g-anna.de/), grant agreement number 16KISK100.