[![PkgGoDev](https://pkg.go.dev/badge/github.com/k0swe/rigctld-go)](https://pkg.go.dev/github.com/k0swe/rigctld-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/k0swe/rigctld-go)](https://goreportcard.com/report/github.com/k0swe/rigctld-go/v4)
[![Test](https://github.com/k0swe/rigctld-go/workflows/Test/badge.svg)](https://github.com/k0swe/rigctld-go/actions/workflows/test.yml)

# rigctld-go

Golang binding for the rigctld amateur radio software's TCP communication interface. This library
supports receiving and sending all rigctld message types up through rigctld v2.5.2.

## Run

This repository is designed as a library but includes a simple driver program to document basic
integration. rigctld must be running and accepting TCP connections for the driver to pick them up.

From this directory:

```shell script
go run cmd/main.go
```
