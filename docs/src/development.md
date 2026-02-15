# Development

## Source code

Source code is written in Go and thus, `go` binaries should be present for compilation and running
tests. You can refer to the [official documentation](https://go.dev/doc/install) for environment
setup.

## Build

### Local build

Local build can be done with the official `go` toolchain :

```shell
womblock $> go build .
womblock $> ./womblock --version
womblock version 6bc1d5bf79864b36685990f39f04b5b57c98aa16 2026-02-15T14:37:28Z go1.26.0 darwin/arm64
```

## Documenting

Documentation is under the `/docs` folder and written with [mdBook](https://rust-lang.github.io/mdBook/).

The following command will help you to preview the documention by serving it via HTTP at
`localhost:3000` by default :

```shell
womblock $> mbdook serve ./docs
 INFO Book building has started
 INFO Running the html backend
 INFO HTML book written to `.../womblock/./docs/book`
 INFO Serving on: http://localhost:3000
```
