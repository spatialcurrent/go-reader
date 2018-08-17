// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package reader

import (
	"bufio"
	"compress/gzip"
	"os"
)

import (
	"github.com/pkg/errors"
)

// GzipFile returns a reader for reading bytes from a gzip-compressed file
// Wraps the "compress/gzip" package.
//
//  - https://golang.org/pkg/compress/gzip/
//
func GzipFile(path string, cache bool) (ByteReadCloser, error) {

	f, err := os.OpenFile(path, os.O_RDONLY, 0600)
	if err != nil {
		return nil, errors.Wrap(err, "Error opening gzip file at \""+path+"\" for reading")
	}

	gr, err := gzip.NewReader(bufio.NewReader(f))
	if gr != nil {
		return nil, errors.Wrap(err, "Error creating gzip reader for file \""+path+"\"")
	}

	if cache {
		return &Cache{
			Reader:  &Reader{Reader: bufio.NewReader(gr), Closer: gr, File: f},
			Content: &[]byte{},
		}, nil
	}

	return &Reader{Reader: bufio.NewReader(gr), Closer: gr, File: f}, nil
}
