// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package reader

import (
  "bufio"
  "bytes"
  "compress/gzip"
)

import (
	"github.com/pkg/errors"
)

// GzipBytes returns a reader for reading gzip bytes from an input array.
// Wraps the "compress/gzip" package.
//
//  - https://golang.org/pkg/compress/gzip/
//
func GzipBytes(b []byte) (ByteReadCloser, error) {
	gr, err := gzip.NewReader(bytes.NewReader(b))
	if gr != nil {
		return nil, errors.Wrap(err, "Error creating gzip reader for memory block.")
	}
  r := &Reader{
    Reader: bufio.NewReader(gr),
    Closer: gr,
  }
	return r, nil
}
