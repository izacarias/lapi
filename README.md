Based on: https://dev.to/hackmamba/build-a-rest-api-with-golang-and-mongodb-gin-gonic-version-269m

# Requirements

## MongoDB

The easy way to run MongoDB for this project is through Docker or Podman

### Docker

    ```
    docker run --name mongodb -p 27017:27017 -d mongodb/mongodb-community-server:latest
    ```

or if is was already initialized

    ```
    docker start mongodb
    ```

### Podman (another option to run containers)

**Pull the MongoDB image from registy**

    ```
    podman pull docker.io/mongodb/mongodb-community-server:latest
    ```

**Run the server as a container**

    ```
    podman run --name mongodb --privileged -p 27017:27017 -d docker.io/mongodb/mongodb-community-server:latest
    ```

**If you want to persist data**

    ```
    mkdir mongodata

    podman run --name mongodb -p 27017:27017 -v ./mongodata:/data/db:rw,U,Z --rm docker.io/mongodb/mongodb-community-server:latest
    ```

**Attach to a running container**
    
    ```
    podman attach mongodb
    ```
> [!Note]
> Depending on your system, it might require to run podman with root privilleges (e.g., Fedora)
> This is due to a requirement of MongoDB and the rights to access files in folders. Without the 
> `sudo` command, SElinux prevents the MondoDB container to write in the filesystem causing an error.

# Testing

The integration tests assume the MongoDB instance is running an is accessible. Please keep in mind that the
database will be recreated and filled with mock data (make sure that the options `PURGEDATABASE` and 
`INSERTMOCKDATA` are enabled in the (.env)[.env] file).

Change to the project directory and run the tests as following:
```
cd <PROJECT_DIRECTORY>
go test -v
```

# Development Instructions

## Update Swagger Specifications

    ```
    swag init -g main.go --output docs/swagger
    ```

# Running tests

For testing purposes, I am using the REST Client extension for VSCode, but any REST Client (Like Postman)
should help. For the REST Client extension, there is a file with the requests under the `./test` folder 
called [Requests.rest](./tests/Requests.rest).