[![GitHub release](https://img.shields.io/github/release/sgaunet/httpfileserver.svg)](https://github.com/sgaunet/httpfileserver/releases/latest)
[![Go Report Card](https://goreportcard.com/badge/github.com/sgaunet/httpfileserver)](https://goreportcard.com/report/github.com/sgaunet/httpfileserver)
![GitHub Downloads](https://img.shields.io/github/downloads/sgaunet/httpfileserver/total)
[![Maintainability](https://api.codeclimate.com/v1/badges/55a5f8d4ec1cc30b3f82/maintainability)](https://codeclimate.com/github/sgaunet/httpfileserver/maintainability)


# httpfileserver

A simple webserver in Golang to expose a directory by http.
You can use the binary (releases) or the docker image (from scratch, it's a multi-arch image).

There is a possibility to add a basic auth by defining environment variable :

* HTTP_USER
* HTTP_PASSWORD


# Build

This project is using :

* golang 1.17+
* [task for development](https://taskfile.dev/#/)
* docker
* [docker buildx](https://github.com/docker/buildx)
* docker manifest
* [goreleaser](https://goreleaser.com/)


## Binary

```
task
```

## Docker Image

```
task image
```

# Release

## Snapshot

```
task snapshot
```

## Release

```
task release
```

