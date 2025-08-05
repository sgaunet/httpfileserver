[![GitHub release](https://img.shields.io/github/release/sgaunet/httpfileserver.svg)](https://github.com/sgaunet/httpfileserver/releases/latest)
[![Go Report Card](https://goreportcard.com/badge/github.com/sgaunet/httpfileserver)](https://goreportcard.com/report/github.com/sgaunet/httpfileserver)
![GitHub Downloads](https://img.shields.io/github/downloads/sgaunet/httpfileserver/total)
![Coverage](https://raw.githubusercontent.com/wiki/sgaunet/httpfileserver/coverage-badge.svg)
[![coverage](https://github.com/sgaunet/httpfileserver/actions/workflows/coverage.yml/badge.svg)](https://github.com/sgaunet/httpfileserver/actions/workflows/coverage.yml)
[![linter](https://github.com/sgaunet/httpfileserver/actions/workflows/linter.yml/badge.svg)](https://github.com/sgaunet/httpfileserver/actions/workflows/linter.yml)
[![Snapshot Build](https://github.com/sgaunet/httpfileserver/actions/workflows/snapshot.yml/badge.svg)](https://github.com/sgaunet/httpfileserver/actions/workflows/snapshot.yml)
[![Release Build](https://github.com/sgaunet/httpfileserver/actions/workflows/release.yml/badge.svg)](https://github.com/sgaunet/httpfileserver/actions/workflows/release.yml)
[![License](https://img.shields.io/github/license/sgaunet/httpfileserver.svg)](LICENSE)


# httpfileserver

A simple webserver in Golang to expose a directory by http.
You can use the binary (releases) or the docker image (from scratch, it's a multi-arch image).

There is a possibility to add a basic auth by defining environment variable :

* HTTP_USER
* HTTP_PASSWORD

## Project Disclaimer

This software project is released under the MIT License and was created primarily for fun and testing purposes. While it may offer some interesting functionalities, please note:

* Intended Use
* This project is experimental in nature
* It serves as a playground for ideas and concepts
* The code may not be optimized or production-ready

## Recommendation

If you find the features provided by this project useful or intriguing, we strongly recommend exploring more mature and established solutions for your actual needs. This project is not intended to compete with or replace professional-grade software in its domain.

## Contributions

While we appreciate your interest, please understand that this project may not be actively maintained or developed further. Feel free to fork and experiment with the code as per the MIT License terms.
Thank you for your understanding and enjoy exploring!

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

