// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

// reader.js is the Javascript version of go-reader.
//
package main

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/pkg/errors"
	"github.com/spatialcurrent/go-reader/reader"
)

func main() {
	js.Global.Set("reader", map[string]interface{}{
		"version": reader.VERSION,
		"fetch":   FetchResource,
	})
}

func FetchResource(uri string, alg string, callback func(...interface{}) *js.Object) {

	go func() {
		r, _, err := reader.OpenResource(uri, alg, 4096, false, nil, nil)
		if err != nil {
			callback("", errors.Wrap(err, "error opening resource from uri "+uri).Error())
			return
		}

		b, err := r.ReadAll()
		if err != nil {
			callback("", errors.Wrap(err, "error reading from resource at uri "+uri).Error())
			return
		}

		callback(string(b), nil)
	}()

}
