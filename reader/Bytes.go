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
)

// Bytes returns a reader for reading the bytes from an input array, and an error if any.
func Bytes(b []byte) (ByteReadCloser, error) {
	return NewCache(&Reader{Reader: bufio.NewReader(bytes.NewReader(b))}), nil
}
