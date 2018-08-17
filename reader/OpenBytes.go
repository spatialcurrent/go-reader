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

// OpenBytes returns a ByteReader for a byte array with a given compression.
// alg may be "snappy", "gzip", or "none."
//
//  - https://golang.org/pkg/compress/gzip/
//  - https://godoc.org/github.com/golang/snappy
//
func OpenBytes(b []byte, alg string) (ByteReadCloser, error) {
	switch alg {
	case "snappy":
		return SnappyBytes(b)
	case "gzip":
		return GzipBytes(b)
	case "none":
		return Bytes(b)
	}
	return nil, errors.New("Unknown algorithm \"" + alg + "\"")
}
