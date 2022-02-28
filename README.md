# Getting Started

This project was created by Go lang.

### `go run main.go`

Runs the app.\
Open Postman or other to view API:
- GET : http://localhost:8888/v1/books
- GET : http://localhost:8888/v1/books/[id]
- POST : http://localhost:8888/v1/books
- PUT : http://localhost:8888/v1/books/[id]
- DELETE : http://localhost:8888/v1/books/[id]

The page manual reload server when you make changes.\
You may also see any lint errors in the console.

# The Go Programming Language

Go is an open source programming language that makes it easy to build simple,
reliable, and efficient software.

![Gopher image](https://golang.org/doc/gopher/fiveyears.jpg)
*Gopher image by [Renee French][rf], licensed under [Creative Commons 3.0 Attributions license][cc3-by].*

Our canonical Git repository is located at https://go.googlesource.com/go.
There is a mirror of the repository at https://github.com/golang/go.

Unless otherwise noted, the Go source files are distributed under the
BSD-style license found in the LICENSE file.

### Download and Install

#### Binary Distributions

Official binary distributions are available at https://golang.org/dl/.

After downloading a binary release, visit https://golang.org/doc/install
for installation instructions.

#### Install From Source

If a binary distribution is not available for your combination of
operating system and architecture, visit
https://golang.org/doc/install/source
for source installation instructions.

### Contributing

Go is the work of thousands of contributors. We appreciate your help!

To contribute, please read the contribution guidelines at https://golang.org/doc/contribute.html.

Note that the Go project uses the issue tracker for bug reports and
proposals only. See https://golang.org/wiki/Questions for a list of
places to ask questions about the Go language.

[rf]: https://reneefrench.blogspot.com/
[cc3-by]: https://creativecommons.org/licenses/by/3.0/

## Gin Web Framework

Build Status codecov Go Report Card GoDoc Join the chat at https://gitter.im/gin-gonic/gin Sourcegraph Open Source Helpers Release TODOs

Gin is a web framework written in Go (Golang). It features a martini-like API with performance that is up to 40 times faster thanks to httprouter. If you need performance and good productivity, you will love Gin.

https://github.com/gin-gonic/gin

## Model binding and validation
To bind a request body into a type, use model binding. We currently support binding of JSON, XML, YAML and standard form values (foo=bar&boo=baz).

Gin uses https://github.com/go-playground/validator for validation. Check the full docs on tags usage https://godoc.org/github.com/go-playground/validator#hdr-Baked_In_Validators_and_Tags.

Note that you need to set the corresponding binding tag on all fields you want to bind. For example, when binding from JSON, set json:"fieldname".

## GORM
Database Connection
The fantastic ORM library for Golang, aims to be developer friendly.

###  Geting Started
GORM Guides https://gorm.io
