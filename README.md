# Boilerplate for REST Microservice in Go
[![Build Status](https://travis-ci.org/goboilerplates/micro-rest.svg?branch=master)](https://travis-ci.org/goboilerplates/micro-rest)
[![codecov](https://codecov.io/gh/goboilerplates/micro-rest/branch/master/graph/badge.svg)](https://codecov.io/gh/goboilerplates/micro-rest)
[![Go Report Card](https://goreportcard.com/badge/github.com/goboilerplates/micro-rest)](https://goreportcard.com/report/github.com/goboilerplates/micro-rest)
[![GoDoc](https://godoc.org/github.com/goboilerplates/micro-rest?status.svg)](https://godoc.org/github.com/goboilerplates/micro-rest)
[![](https://images.microbadger.com/badges/image/goboilerplates/micro-rest.svg)](https://microbadger.com/images/goboilerplates/micro-rest)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/goboilerplates/micro-rest/blob/master/LICENSE)

## Features
- REST API and Swagger UI
- CI with Travis
- Docker Build

## Installation

Get the micro-rest repository

```
go get github.com/goboilerplates/micro-rest

cd echo $GOPATH/src/github.com/goboilerplates/micro-rest
```

And install dependencies

```
go get -u github.com/golang/dep/cmd/dep

dep ensure
```

## Running the tests

Run all tests

```
go test ./...
```

Or run all tests with coverage

```
bash script/coverage.sh
```

## Build and Run

Run main.go
``` bash
go run main.go
# serve at localhost:9000
```

Build and run native binary

``` bash
bash script/Build.sh

./micro-rest.out
```
Build native binary for multiple platforms (darwin, windows and linux)

```
bash script/BuildMulti.sh
```

## Environment variables

```bash
    # enable production mode, default is false
    env GBP_PROMODE=true
```
## Docker support 

Build docker image

```
bash script/Dockerbuild.sh
```

Run docker container

```
docker run -d --name micro-rest -p 9000:9000 goboilerplates/micro-rest
```
## Contributing

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests to us.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/goboilerplates/micro-rest/tags). 

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details

