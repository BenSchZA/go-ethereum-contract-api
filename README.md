# Go Contract API

Go API to interact with Solidity contract deployed on Ethereum blockchain using Geth binaries.

---

## Project setup

Install and use dep, the Go dependency manager.

### Working directory setup

Initialize project directory using `dep init`.

### Package management

To install a package:

`dep ensure -add github.com/gorilla/mux`

To install all packages from Gopkg and those referenced in code:

`dep ensure`

**Make sure to export correct GOPATH**

## Deployment

Use Dockerfile for local development and Dockerfile.prod for production.

### Docker deployment

`docker build -t 'go-api' .`

`docker run -p 3000:3000 -it 'go-api'`

### Docker Go debugging

See [Debugging containerized Go applications](https://blog.jetbrains.com/go/2018/04/30/debugging-containerized-go-applications/)

## Compile Solidity contract to Go binary

### Install SOLC, using relevant package manager

`pacaur -S solidity`

`abigen --sol=contract.sol --pkg=main --out=contract.go`


