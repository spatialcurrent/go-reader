// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package reader

import (
  "github.com/pkg/errors"
)

// OpenFile returns a ByteReader for a file with a given compression.
// alg may be "snappy", "gzip", or "none."
//
//  - https://golang.org/pkg/compress/gzip/
//  - https://godoc.org/github.com/golang/snappy
//
func OpenFile(uri string, alg string) (ByteReadCloser, error) {
  switch alg {
  case "snappy":
    return SnappyFile(uri)
  case "gzip":
    return GzipFile(uri)
  case "none":
    return File(uri)
  }
  return nil, errors.New("Unknown algorithm \""+alg+"\"")
}
