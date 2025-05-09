<div align="center">

<img src="readme/nx-s3-server-transparent.png" width="400px">

# NX S3 Server

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/mjarmoc/nx-s3-server?style=flat)
![GitHub forks](https://img.shields.io/github/forks/mjarmoc/nx-s3-server?style=flat)

NXS3Server is a lightweight self-hosted cache server for Nx written in GO.<br/>
It is using AWS S3 Buckets as a backend.

</div>

## Features

- [x] Fully satisfies specification: [https://nx.dev/recipes/running-tasks/self-hosted-caching#build-your-own-caching-server]
- [x] Endpoint to get all Cache Artifats

## Installation

```sh
make deps
```

## Local Development

```sh
docker-compose up
```
