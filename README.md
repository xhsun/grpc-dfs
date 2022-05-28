# Distributed File System Using gRPC

A simple distributed file system (DFS) that supports upload, download, list, and remove to interact with files on a remote server. This DFS contains a basic SSL/TLS authentication to encrypt all data exchange between a remote server and a CLI.

## Prerequisites

Make sure you have the version of Go from [go.mod](./go.mod) installed. Then clone the repository.

_If you would like to try to run the application inside Docker, make sure you installed Docker following their [official guide](https://docs.docker.com/get-docker/)._

### Generate Certificate For Local Testing

Make sure you have `openssl` installed in your local machine, then you can follow the steps below to generate your own certificate that will be used for SSL/TLS authentication between the server and the CLI.

*Note: Certificates generated following the following steps are for local testing purpose only!*

#### Generate CA private key and self-signed certificate

```
openssl req -x509 -newkey rsa:4096 -nodes -keyout ca-key.pem -out ca-cert.pem -subj "/C=CA/ST=STATE/L=CITY/O=DEV/OU=EXAMPLE/CN=localhost/emailAddress=example@localhost"
```

#### Generate Web Server’s Private Key and CSR (Certificate Signing Request)
```
openssl req -newkey rsa:4096 -nodes -keyout server-key.pem -out server-req.pem -subj "/C=CA/ST=STATE/L=CITY/O=DEV/OU=EXAMPLE/CN=localhost/emailAddress=example@localhost"
```

#### Sign the Web Server Certificate Request (CSR)
```
openssl x509 -req -in server-req.pem -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out server-cert.pem -extfile server-ext.cnf
```

Content of `server-ext.cnf`:
```
subjectAltName=DNS:localhost,IP:127.0.0.1
```

#### (Optional) See what is inside the certificate
```
openssl x509 -in server-cert.pem -noout -text
```

### Development
If you would like to add additional features to this project, make sure to install `protobuf-compiler` by following [the official guide](https://grpc.io/docs/protoc-installation/).

#### Regenerate gRPC client and server interfaces 
```
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./dfs/dfs.proto
```

## Run The Project

Once you have installed the necessary dependencies by following the [prerequisites](#prerequisites), you will be able to run this project locally.

### Run The Server

There is two ways you can run the server, you can either run it directly on your local machine or you can run it inside a docker container.

_Note: If you would like to override default configuration of the server, take a look at [the server custom configuration section](#server-custom-configuration)._

#### Run It Directly

To run the server directly, go to the project root and run:

```
make build-server
./dfsServer
```

#### Run It Inside Docker Container

To run the server inside a docker container, go to the project root and run:

```
docker compose up
```

_Note: Default source folder for server side certificate is `./x509/`. If you are storing your server side certificate in a different location, you will need to update the source volume location in `docker-compose.yml`_

If you prefer to build your own image, you can run the following command:

```
make docker-server
docker run -v /<path>/<to>/<cert>:/server/x509/:ro dfs-server
```

### Run The CLI

In order to run the CLI, you need to install it first by go to the project root and run:

```
make install-cli
```

### Supported CLI Commands

DFS CLI supports the following commands. If you would like to learn more about each command, please run `dfs-cli help`.

#### Upload

![Upload Example](https://i.imgur.com/LhD8ZKr.png)

```
dfs-cli upload [FILE_PATH]
```

Uploads content of `FILE_PATH` to DFS server.

*Note: The server will store the file content at `FILE_PATH` in its own file system.*

#### Copy

![Copy Example](https://i.imgur.com/anGRed0.png)

```
dfs-cli copy [SERVER_FILE_PATH] [LOCAL_FILE_PATH]
```

Copy file content of `SERVER_FILE_PATH` from DFS server to `LOCAL_FILE_PATH` in the local file system.

#### Remove

![Remove Example](https://i.imgur.com/cUajTX0.png)

```
dfs-cli copy [SERVER_FILE_PATH]
```

Remove `SERVER_FILE_PATH` from DFS server.

#### List

![List Example](https://i.imgur.com/qGcFGN0.png)

```
dfs-cli list
```

List all files in DFS server's file storage.


#### Help

![Help Example](https://i.imgur.com/FRk6FfZ.png)

```
dfs-cli help
```

Display help text.

## Custom Configurations

Both server and CLI offer some custom configuration.

### Server Custom Configuration

The server will read configuration from a configuration file or environment variable.

_Note: Configuration from environment variable will override configuration from file._

#### Via environment variable

Server will read configuration from the following environment variables:

- `DFS_SERVER_CONFIG_PATH`: Path to configuration JSON file
- `DFS_SERVER_PORT`: Server port
- `DFS_SERVER_CERTIFICATE`: Path to x509 certificate file
- `DFS_SERVER_KEY`: Path to x509 private key file
- `DFS_SERVER_STORAGE`: Server will store files at this location

#### Via configuration file

The server will attempt to read content of configuration JSON file from `DFS_SERVER_CONFIG_PATH` environment variable. If that environment variable it not set, it will attempt to read from `config/config.json`.

The following is all the supported configurations via file:

``` json
{
    // Server port
    "port": 50051, 
    // x509 certificate file path
    "certFilePath": "x509/server-cert.pem", 
    // x509 private key file path
    "keyFilePath": "x509/server-key.pem",
    // Server will store files at this location
    "fileStoragePath": "serverFiles" 
}
```

### CLI Custom Configuration

The CLI swill read configuration from the following environment variables:

- `DFS_CLI_SERVER_ADDR`: DFS server IP address (format: `<IP>:<PORT>`)
- `DFS_CLI_SERVER_HOST_OVERRIDE`: DFS server host name override
- `DFS_CLI_CERTIFICATE_PATH`: Path to x509 certificate file