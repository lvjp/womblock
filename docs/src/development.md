# Development

## Source code

Source code is written in Go and thus, `go` binaries should be present for compilation and running
tests. You can refer to the [official documentation](https://go.dev/doc/install) for environment
setup.

## Build

### Local build

Local build can be done with the official `go` toolchain :

```shell
wtf-go $> go build .
wtf-go $> ./wtf-go --version
wtf-go version 6bc1d5bf79864b36685990f39f04b5b57c98aa16 2026-02-15T14:37:28Z go1.26.0 darwin/arm64
```

### Docker build

A docker file located at `build/Dockerfile` permit to build. You will need to have one of those
installed:

- [Docker Engine](https://docs.docker.com/engine/)
- [Docker Desktop](https://docs.docker.com/desktop/)
- [OrbStack](https://orbstack.dev/) (macOS only)

The Dockerfile is located at `build/Dockerfile`.

Docker build command:

```shell
wtf-go $> docker build --file build/Dockerfile --tag wtf-go:local .
```

On success, the image `wtf-go:local` is created.

```shell
wtf-go $> docker image ls wtf-go:local
REPOSITORY   TAG       IMAGE ID       CREATED          SIZE
wtf-go     local     a541f56b4a99   41 seconds ago   22MB
```

## Deployment

All supported deployment stacks related files are located into the `/deployment` folder.

### Local deployement

Local deployment is based on [docker-compose](https://docs.docker.com/compose/).

Exposed ports :

| Port | Service |
|------|---------|
| 8081 | backend |

## Documenting

Documentation is under the `/docs` folder and written with [romero](https://romero.srht.site/).

The following command will help you to preview the documention by serving it via HTTP at
`localhost:8003`:

```shell
wtf-go $> romero serve
12:12AM INF Building the documentation render=html
12:12AM INF html: Building page 'README.md' -> 'book/index.html'
12:12AM INF html: Building page 'configuration.md' -> 'book/configuration.html'
12:12AM INF html: Building page 'development.md' -> 'book/development.html'
12:12AM INF html: Building page '404.md' -> 'book/404.html'
12:12AM INF Checking for dead links
12:12AM INF Documentation building is done! duration=208.610042 output=book render=html
12:12AM INF HTTP server listening on http://0.0.0.0:8003
12:12AM INF Network URL: http://192.168.1.18:8003

Scan this QR code to open on your mobile device:

[...SNIP...]
```
