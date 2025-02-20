Based on: https://dev.to/hackmamba/build-a-rest-api-with-golang-and-mongodb-gin-gonic-version-269m



## MongoDB
### Docker
```
docker run --name mongodb -p 27017:27017 -d mongodb/mongodb-community-server:latest
```
or
```
docker start CONTAINDER_ID
```
### Podman
**Pull the MongoDB image from registy**
```
podman pull docker.io/mongodb/mongodb-community-server:latest
```
**Run the server as a container**
```
podman run --name mongodb -p 27017:27017 -d docker.io/mongodb/mongodb-community-server:latest
```
**To persist data**
```
mkdir mongodata

podman run --name mongodb -p 27017:27017 -v ./mongodata:/data/db:rw,U,Z --rm docker.io/mongodb/mongodb-community-server:latest
```

# Update Swagger Specifications
```
swag init -g main.go --output docs/swagger
```

or just

```

```