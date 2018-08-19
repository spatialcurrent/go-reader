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
	"github.com/pkg/errors"
)

import (
	"github.com/aws/aws-sdk-go/service/s3"
)

func OpenResource(uri string, alg string, cache bool, s3_client *s3.S3) (ByteReadCloser, error) {

	if uri == "stdin" {
		return OpenStdin(alg, cache)
	}

	scheme, path := SplitUri(uri)
	if scheme == "s3" {
		i := strings.Index(path, "/")
		if i == -1 {
			return &Reader{}, errors.New("path missing bucket")
		}
		return OpenS3Object(path[0:i], path[i+1:], alg, cache, s3_client)
	} else if scheme == "" {
		switch alg {
		case "snappy":
			return SnappyFile(uri, cache)
		case "gzip":
			return GzipFile(uri, cache)
		case "none":
			return File(uri, cache)
		}
	}

	return nil, errors.New("Unknown algorithm \"" + alg + "\"")
}
