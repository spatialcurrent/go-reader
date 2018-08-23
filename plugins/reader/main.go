// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

// reader.so creates a shared library of Go that can be called by C, C++, or Python
//

package main

import (
	"C"
	"github.com/pkg/errors"
	"github.com/spatialcurrent/go-reader/reader"
)

func main() {}

//export ReadAll
func ReadAll(uri *C.char, alg *C.char, result **C.char) *C.char {

	r, _, err := reader.OpenResource(C.GoString(uri), C.GoString(alg), 4096, false, nil, nil)
	if err != nil {
		return C.CString(errors.Wrap(err, "error opening resource from uri "+C.GoString(uri)).Error())
	}

	b, err := r.ReadAll()
	if err != nil {
		return C.CString(errors.Wrap(err, "Error reading from resource").Error())
	}

	*result = C.CString(string(b))

	return nil
}

//export Version
func Version() *C.char {
	return C.CString(reader.VERSION)
}
