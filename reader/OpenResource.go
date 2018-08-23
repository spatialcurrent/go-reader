// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package reader

import (
	"strings"
)

import (
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/colinmarc/hdfs"
	"github.com/mitchellh/go-homedir"
	"github.com/pkg/errors"
)

func OpenResource(uri string, alg string, buffer_size int, cache bool, s3_client *s3.S3, hdfs_client *hdfs.Client) (ByteReadCloser, *Metadata, error) {

	if uri == "stdin" {
		brc, err := OpenStdin(alg, cache)
		return brc, nil, err
	}

	scheme, path := SplitUri(uri)
	if scheme == "http" || scheme == "https" {
		return OpenHTTPFile(uri, alg, cache)
	} else if scheme == "s3" {
		i := strings.Index(path, "/")
		if i == -1 {
			return &Reader{}, nil, errors.New("path missing bucket")
		}
		return OpenS3Object(path[0:i], path[i+1:], alg, cache, s3_client)
	} else if scheme == "hdfs" {
		brc, err := OpenHDFSFile(path, alg, cache, hdfs_client)
		return brc, nil, err
	} else if scheme == "" {
		pathExpanded, err := homedir.Expand(path)
		if err != nil {
			return nil, nil, errors.Wrap(err, "Error expanding resource file path "+path)
		}
		switch alg {
		case "snappy":
			brc, err := SnappyFile(pathExpanded, cache, buffer_size)
			return brc, nil, err
		case "gzip":
			brc, err := GzipFile(pathExpanded, cache, buffer_size)
			return brc, nil, err
		case "bzip2":
			brc, err := Bzip2File(pathExpanded, cache, buffer_size)
			return brc, nil, err
		case "none":
			brc, err := File(pathExpanded, cache, buffer_size)
			return brc, nil, err
		}
	}

	return nil, nil, errors.New("Unknown algorithm \"" + alg + "\"")
}
