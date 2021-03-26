# Casty protocl package for Golang
This repository contains common protocol definitions for casty services

## Install protoc-gen-go compiler
To install `protoc-gen-go` take a look at [this documentation](https://github.com/golang/protobuf#installation)!

## Directory Structure
The directory structure should match the protocol package.  
we use `/protobuf` directory for our proto files and then we compile them into `/proto` directory!

## Usage
To use this package on other Go services, simply use command below to install it.
```bash
$ go get github.com/castyapp/libcasty-protocol-go
```

## Compile protobuffers
This command will compile `.proto` files of the `/protobuf` dir into `/proto` directory!
```bash
$ make compile
```

## Contributing
Thank you for considering contributing to Casty projects!

## License
Casty is an open-source software licensed under the MIT license.
