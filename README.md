[![Build Status](https://travis-ci.org/spatialcurrent/go-reader.svg)](https://travis-ci.org/spatialcurrent/go-reader) [![Go Report Card](https://goreportcard.com/badge/spatialcurrent/go-reader)](https://goreportcard.com/report/spatialcurrent/go-reader)  [![GoDoc](https://godoc.org/github.com/spatialcurrent/go-reader?status.svg)](https://godoc.org/github.com/spatialcurrent/go-reader) [![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://github.com/spatialcurrent/go-reader/blob/master/LICENSE.md)

# go-reader

# Description

**go-reader** is a simple library for managing the closing of readers and underlying resources.

# Usage

**CLI**

You can use the command line tool to convert between formats.

```
Usage: go-reader -i INPUT_FORMAT -o OUTPUT_FORMAT
Options:
  -help
    	Print help.
  -i string
    	The input format: csv, hcl, hcl2, json, jsonl, properties, toml, yaml
  -o string
    	The output format: csv, hcl, hcl2, json, jsonl, properties, toml, yaml
  -version
    	Prints version to stdout.
```

**Go**

You can import **go-reader** as a library with:

```go
import (
  "github.com/spatialcurrent/go-reader/reader"
)
...
```

See [reader](https://godoc.org/github.com/spatialcurrent/go-reader/reader) in GoDoc for information on how to use Go API.

# Examples:

TBD

# Contributing

[Spatial Current, Inc.](https://spatialcurrent.io) is currently accepting pull requests for this repository.  We'd love to have your contributions!  Please see [Contributing.md](https://github.com/spatialcurrent/go-reader/blob/master/CONTRIBUTING.md) for how to get started.

# License

This work is distributed under the **MIT License**.  See **LICENSE** file.
