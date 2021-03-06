// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package reader

import (
	"bufio"
	"os"
)

import (
	"github.com/pkg/errors"
)

// File returns a ByteReader for reading bytes without any transformation from a file, and an error if any.
func File(path string, cache bool, buffer_size int) (ByteReadCloser, error) {

	f, err := os.OpenFile(path, os.O_RDONLY, 0600)
	if err != nil {
		return nil, errors.Wrap(err, "Error opening file at \""+path+"\" for reading")
	}

	br := bufio.NewReaderSize(f, buffer_size)

	if cache {
		return NewCache(&Reader{Reader: br, File: f}), nil
	}

	return &Reader{Reader: br, File: f}, nil
}
