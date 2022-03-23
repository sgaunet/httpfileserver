# httpfileserver

A simple webserver in Golang to expose a directory by http.
You can use the binary (releases) or the docker image (from scratch, it's a multi-arch image).

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

