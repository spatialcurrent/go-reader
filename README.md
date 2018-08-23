[![Build Status](https://travis-ci.org/spatialcurrent/go-reader.svg)](https://travis-ci.org/spatialcurrent/go-reader) [![Go Report Card](https://goreportcard.com/badge/spatialcurrent/go-reader)](https://goreportcard.com/report/spatialcurrent/go-reader)  [![GoDoc](https://godoc.org/github.com/spatialcurrent/go-reader?status.svg)](https://godoc.org/github.com/spatialcurrent/go-reader) [![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://github.com/spatialcurrent/go-reader/blob/master/LICENSE.md)

# go-reader

# Description

**go-reader** is a simple library for managing the reading of uncompressed text from underlying resources.

Using cross compilers, this library can also be called by other languages.  This library is cross compiled into a Shared Object file (`*.so`).  The Shared Object file can be called by `C`, `C++`, and `Python` on Linux machines.  See the examples folder for patterns that you can use.  This library is also compiled to pure `JavaScript` using [GopherJS](https://github.com/gopherjs/gopherjs).

# Usage

**CLI**

You can use the command line tool to convert between formats.

```
Usage: reader -uri URI [-alg [bzip2|gzip|snappy|none]] [-verbose] [-version]
Options:
  -alg string
    	Stream input compression algorithm for nodes, using: bzip2, gzip, snappy, or none. (default "none")
  -aws_access_key_id string
    	Defaults to value of environment variable AWS_ACCESS_KEY_ID
  -aws_default_region string
    	Defaults to value of environment variable AWS_DEFAULT_REGION.
  -aws_secret_access_key string
    	Defaults to value of environment variable AWS_SECRET_ACCESS_KEY.
  -aws_session_token string
    	Defaults to value of environment variable AWS_SESSION_TOKEN.
  -buffer_size int
    	The input reader buffer size (default 4096)
  -hdfs_name_node string
    	Defaults to value of environment variable HDFS_DEFAULT_NAME_NODE.
  -help
    	Print help
  -uri string
    	"stdin" or uri to input file
  -version
    	Prints version to stdout
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

**JavaScript**

```html
<html>
  <head>
    <script src="https://...reader.js"></script>
  </head>
  <body>
    <script>
      reader.open(uri, "none", function(text){ ... })
      ...
    </script>
  </body>
</html>
```

# Examples:

TBD

# Building

**CLI**

The command line go-reader program can be built with the `scripts/build_cli.sh` script.

**JavaScript**

You can compile go-reader to pure JavaScript with the `scripts/build_javascript.sh` script.

**Shared Object**

The `scripts/build_so.sh` script is used to build a Shared Object (`*.go`), which can be called by `C`, `C++`, and `Python` on Linux machines.

**Changing Destination**

The default destination for build artifacts is `go-reader/bin`, but you can change the destination with a CLI argument.  For building on a Chromebook consider saving the artifacts in `/usr/local/go/bin`, e.g., `bash scripts/build_cli.sh /usr/local/go/bin`

# Contributing

[Spatial Current, Inc.](https://spatialcurrent.io) is currently accepting pull requests for this repository.  We'd love to have your contributions!  Please see [Contributing.md](https://github.com/spatialcurrent/go-reader/blob/master/CONTRIBUTING.md) for how to get started.

# License

This work is distributed under the **MIT License**.  See **LICENSE** file.
