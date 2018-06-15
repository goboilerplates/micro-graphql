# Boilerplate for GraphQL Microservice in Go
[![Build Status](https://travis-ci.org/goboilerplates/micro-graphql.svg?branch=master)](https://travis-ci.org/goboilerplates/micro-graphql)
[![codecov](https://codecov.io/gh/goboilerplates/micro-graphql/branch/master/graph/badge.svg)](https://codecov.io/gh/goboilerplates/micro-graphql)
[![Go Report Card](https://goreportcard.com/badge/github.com/goboilerplates/micro-graphql)](https://goreportcard.com/report/github.com/goboilerplates/micro-graphql)
[![GoDoc](https://godoc.org/github.com/goboilerplates/micro-graphql?status.svg)](https://godoc.org/github.com/goboilerplates/micro-graphql)
[![](https://images.microbadger.com/badges/image/goboilerplates/micro-graphql.svg)](https://microbadger.com/images/goboilerplates/micro-graphql)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/goboilerplates/micro-graphql/blob/master/LICENSE)

## Features
- GraphQL API
- Middlewares (cors, gzip and static)
- CI with Travis
- Docker Build

## Installation

Get the micro-graphql repository

```
go get github.com/goboilerplates/micro-graphql

cd echo $GOPATH/src/github.com/goboilerplates/micro-graphql
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
bash script/Run.sh
# serve at localhost:9000
```

Build and run native binary

``` bash
bash script/Build.sh

./micro-graphql.out
```
Build native binary for multiple platforms (darwin, windows and linux)

```
bash script/BuildMulti.sh
```

## Environment variables

```bash
    # enable production mode, default is true
    env GBP_PROMODE=false
```
## Docker support 

Build docker image

```
bash script/Dockerbuild.sh
```

Run docker container

```
docker run -d --name micro-graphql -p 9000:9000 goboilerplates/micro-graphql
```
## Contributing

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests to us.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/goboilerplates/micro-graphql/tags). 

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details

