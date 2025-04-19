# learn-cicd-starter (Notely)

This repo contains the starter code for the "Notely" application for the "Learn CICD" course on [Boot.dev](https://boot.dev).

![Badge](https://github.com/sirtaylor88/learn-cicd-starter/actions/workflows/ci.yml/badge.svg)

## Local Development

### Prerequisites

Make sure you're on Go version 1.22+.

Create a `.env` file in the root of the project with the following contents:

```bash
PORT="8080"
```

### Build and run

Run the server:

```bash
go build -o notely && ./notely
```

*This starts the server in non-database mode.* It will serve a simple webpage at `http://localhost:8080`.

You do *not* need to set up a database or any interactivity on the webpage yet. Instructions for that will come later in the course!

#### using Docker

Build on Dockerhub, replace `DOCKERHUB_NAMESPACE` with your dockerhub name.

```bash
docker build -t DOCKERHUB_NAMESPACE/notely:latest .
```

Run container

```bash
docker run -e PORT=8080 -p 8080:8080 DOCKERHUB_NAMESPACE/notely:latest
```
