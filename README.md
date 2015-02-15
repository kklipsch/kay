A command line tool to keep track of autobiography chapters organized by year.

# Usage 

## To install

To install from the kay.zip file extract the file kay and put it in a location in your path.  For autocompletion extract kay.bash and make sure it is sourced by your bashrc.  For installing locally from the repo you can:

```bash
make install
```

## Instructions

To see the command line options:

```bash
kay -h
```
# Development

This is a golang project that uses a Makefile to wrap development operations.  It builds with the assumption that it will use the current directory as the GOPATH.

## Prerequisites

- go 1.3.1
- ansible 1.7.1
- GNU Make 3.81

To get dependencies:

```bash
make dep
```

To build:

```bash
make
```

To make the distribution package:

```bash
make clean package
```

## Integration tests

In the integration folder is an ansible playbook that can be used to run integration tests:

```bash
make integration
```
