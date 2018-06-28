// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package reader

import (
  "bytes"
  "bufio"
)

// Bytes returns a reader for reading the bytes from an input array, and an error if any.
func Bytes(b []byte) (ByteReadCloser, error) {
  r := &Reader{
    Reader: bufio.NewReader(bytes.NewReader(b)),
  }
	return r, nil
}
