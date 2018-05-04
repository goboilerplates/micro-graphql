# Sample Core
[![Build Status](https://travis-ci.org/goboilerplates/core.svg?branch=master)](https://travis-ci.org/goboilerplates/core)
[![codecov](https://codecov.io/gh/goboilerplates/core/branch/master/graph/badge.svg)](https://codecov.io/gh/goboilerplates/core)
[![Go Report Card](https://goreportcard.com/badge/github.com/goboilerplates/core)](https://goreportcard.com/report/github.com/goboilerplates/core)
[![GoDoc](https://godoc.org/github.com/goboilerplates/core?status.svg)](https://godoc.org/github.com/goboilerplates/core)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/goboilerplates/core/blob/master/LICENSE)

## Installation

Get the core repository

```
go get github.com/goboilerplates/core

cd echo $GOPATH/src/github.com/goboilerplates/core
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
go run test/main.go
```

Build and run native binary

``` bash
bash script/Build.sh

./core.out
```
## Contributing

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests to us.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/goboilerplates/core/tags). 

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details

