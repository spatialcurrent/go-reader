// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package reader

import (
	"archive/zip"
	"bufio"
	"bytes"
	"compress/bzip2"
	"compress/gzip"
	"io/ioutil"
	"net/http"
)

import (
	"github.com/golang/snappy"
	"github.com/pkg/errors"
)

// OpenHTTPFile returns a ByteReadCloser for an object for a web address
// alg may be "bzip2", "gzip", "snappy", or "".
//
//  - https://golang.org/pkg/compress/bzip2/
//  - https://golang.org/pkg/compress/gzip/
//  - https://godoc.org/github.com/golang/snappy
//
func OpenHTTPFile(uri string, alg string, cache bool) (ByteReadCloser, *Metadata, error) {
	resp, err := http.Get(uri)
	if err != nil {
		return &Reader{}, nil, errors.Wrap(err, "Error opening file from uri "+uri)
	}

	metadata := NewMetadataFromHeader(resp.Header)

	if alg == "gzip" {

		gr, err := gzip.NewReader(resp.Body)
		if err != nil {
			return &Reader{}, nil, errors.Wrap(err, "Error creating gizp reader for file at uri "+uri+".")
		}

		if cache {
			return &Cache{
				Reader:  &Reader{Reader: bufio.NewReader(gr), Closer: gr},
				Content: &[]byte{},
			}, NewMetadataFromHeader(resp.Header), nil
		}

		return &Reader{Reader: bufio.NewReader(gr), Closer: gr}, metadata, nil

	}

	if alg == "bzip2" {

		br := bzip2.NewReader(resp.Body)

		if cache {
			return &Cache{
				Reader:  &Reader{Reader: bufio.NewReader(br), Closer: resp.Body},
				Content: &[]byte{},
			}, NewMetadataFromHeader(resp.Header), nil
		}

		return &Reader{Reader: bufio.NewReader(br), Closer: resp.Body}, metadata, nil

	}

	if alg == "snappy" {

		sr := snappy.NewReader(bufio.NewReader(resp.Body))

		if cache {
			return &Cache{
				Reader:  &Reader{Reader: bufio.NewReader(sr), Closer: resp.Body},
				Content: &[]byte{},
			}, NewMetadataFromHeader(resp.Header), nil
		}

		return &Reader{Reader: bufio.NewReader(sr), Closer: resp.Body}, metadata, nil
	}

	if alg == "zip" {

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return &Reader{}, metadata, errors.Wrap(err, "error reading bytes from zip-compressed http file")
		}

		zr, err := zip.NewReader(bytes.NewReader(body), int64(len(body)))
		if err != nil {
			return &Reader{}, metadata, errors.Wrap(err, "error creating zip reader for file at uri "+uri+".")
		}

		if len(zr.File) != 1 {
			return &Reader{}, metadata, errors.New("error zip file has more than one internal file.")
		}

		zfr, err := zr.File[0].Open()
		if err != nil {
			return &Reader{}, metadata, errors.Wrap(err, "error opening internal file for zip.")
		}

		if cache {
			return &Cache{
				Reader:  &Reader{Reader: bufio.NewReader(zfr)},
				Content: &[]byte{},
			}, metadata, nil
		}

		return &Reader{Reader: bufio.NewReader(zfr)}, metadata, nil
	}

	if cache {
		return &Cache{
			Reader:  &Reader{Reader: bufio.NewReader(resp.Body), Closer: resp.Body},
			Content: &[]byte{},
		}, metadata, nil
	}

	return &Reader{Reader: bufio.NewReader(resp.Body), Closer: resp.Body}, metadata, nil

}
