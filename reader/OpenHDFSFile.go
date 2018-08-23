// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package reader

import (
	"bufio"
	"compress/bzip2"
	"compress/gzip"
)

import (
	"github.com/colinmarc/hdfs"
	"github.com/golang/snappy"
	"github.com/pkg/errors"
)

// OpenHDFSFile returns a ByteReadCloser for an object in AWS S3.
// alg may be "bzip2", "gzip", "snappy", or "".
//
//  - https://golang.org/pkg/compress/bzip2/
//  - https://golang.org/pkg/compress/gzip/
//  - https://godoc.org/github.com/golang/snappy
//
func OpenHDFSFile(path string, alg string, cache bool, hdfs_client *hdfs.Client) (ByteReadCloser, error) {

	fileReader, err := hdfs_client.Open(path)
	if err != nil {
		return &Reader{}, errors.Wrap(err, "Error opening file on HDFS at path "+path)
	}

	if alg == "gzip" {

		gr, err := gzip.NewReader(fileReader)
		if err != nil {
			return &Reader{}, errors.Wrap(err, "Error creating gizp reader for AWS s3 object at hdfs://"+path+".")
		}

		if cache {
			return &Cache{
				Reader:  &Reader{Reader: bufio.NewReader(gr), Closer: gr},
				Content: &[]byte{},
			}, nil
		}

		return &Reader{Reader: bufio.NewReader(gr), Closer: gr}, nil

	}

	if alg == "bzip2" {

		br := bzip2.NewReader(fileReader)

		if cache {
			return &Cache{
				Reader:  &Reader{Reader: bufio.NewReader(br), Closer: fileReader},
				Content: &[]byte{},
			}, nil
		}

		return &Reader{Reader: bufio.NewReader(br), Closer: fileReader}, nil

	}

	if alg == "snappy" {

		sr := snappy.NewReader(bufio.NewReader(fileReader))

		if cache {
			return &Cache{
				Reader:  &Reader{Reader: bufio.NewReader(sr), Closer: fileReader},
				Content: &[]byte{},
			}, nil
		}

		return &Reader{Reader: bufio.NewReader(sr), Closer: fileReader}, nil
	}

	if cache {
		return &Cache{
			Reader:  &Reader{Reader: bufio.NewReader(fileReader), Closer: fileReader},
			Content: &[]byte{},
		}, nil
	}

	return &Reader{Reader: bufio.NewReader(fileReader), Closer: fileReader}, nil

}
