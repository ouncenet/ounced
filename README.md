

# Ounced


## Requirements

Go 1.18 or later.

## Installation

#### Build from Source

- Install Go according to the installation instructions here:
  http://golang.org/doc/install

- Ensure Go was installed properly and is a supported version:

```bash
$ go version
```

- Run the following commands to obtain and install ounced including all dependencies:

```bash
$ git clone https://github.com/ouncenet/ounced
$ cd ounced
$ go install . ./cmd/...
```

- Ounced (and utilities) should now be installed in `$(go env GOPATH)/bin`. If you did
  not already add the bin directory to your system path during Go installation,
  you are encouraged to do so now.

## Getting Started

Ounced has several configuration options available to tweak how it runs, but all
of the basic operations work with zero configuration.

```bash
$ ounced
```



## License

Ounced is licensed under the copyfree [ISC License](https://choosealicense.com/licenses/isc/).
