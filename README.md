# go-rest-template

Go (Golang) API REST Template/Boilerplate with Gin Framework

[![Go Report Card](https://goreportcard.com/badge/github.com/bingoohuang/go-rest-template)](https://goreportcard.com/report/github.com/bingoohuang/go-rest-template)
[![Open Source Love](https://badges.frapsoft.com/os/mit/mit.svg?v=102)](https://github.com/ellerbrock/open-source-badge/)
[![Build Status](https://travis-ci.com/bingoohuang/go-rest-template.svg?branch=master)](https://travis-ci.com/bingoohuang/go-rest-template)

## 0. Run

- `make run`
- access **http://localhost:3000/docs/index.html**

## 1. Run with Docker

- Build `make build-docker`
- Run `docker run -p 3000:3000 api-rest`
- Test `go test -v ./test/...`

## 2. Generate Docs

- Get swag `go get -u github.com/swaggo/swag/cmd/swag`
- Generate docs `swag init --dir cmd/api --parseDependency --output docs`
