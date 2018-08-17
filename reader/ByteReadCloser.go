// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package reader

import (
	"io"
)

// ByteReader is an interface that extends io.Reader, io.ByteReader, and adds a range function.
// ByteReader provides functions for reading bytes.
type ByteReadCloser interface {
	ByteReader
	io.Closer
	ReadAt(i int) (byte, error)
	ReadAll() ([]byte, error)
	ReadRange(start int, end int) ([]byte, error)
}
