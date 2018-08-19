// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

// Package reader provides the interfaces, embedded structs, and implementing code
// for normalizing the reading a stream of bytes from compressed files.
// This package supports the gzip and snappy compression algorithms.  No compression is usually identified as "none".
// This package is used by the go-stream package.
//  - https://godoc.org/github.com/spatialcurrent/go-stream/stream
//
// Usage
//
// You can import reader as a package into your own Go project or use the command line interface.
//
//  import (
//    "github.com/spatialcurrent/go-reader/reader"
//  )
//
//  r, err := reader.OpenFile("data-for-2018.sz", "snappy")
//  if err != nil {
//    panic(err)
//  }
//  for {
//    b, err := input_reader.ReadBytes([]byte("\n")[0])
//    if err != nil {
//      if err != io.EOF {
//        fmt.Println(errors.Wrap(err, "Error reading bytes from file"))
//        os.Exit(1)
//      }
//    }
//    if len(b) > 0 {
//      fmt.Println(string(b))
//    }
//    if err != nil && err == io.EOF {
//      break
//    }
//  }
//
//
// See the github.com/go-reader/cmd/go-reader package for a command line tool for testing DFL expressions.
//
//  - https://godoc.org/github.com/spatialcurrent/go-reader/reader
//
// Projects
//
// go-reader is used by the go-stream project.
//  - https://godoc.org/github.com/spatialcurrent/go-stream/stream
//
package reader

import (
	"io"
	"io/ioutil"
	"os"
)

import (
	"github.com/pkg/errors"
)

// Reader is a struct for normalizing reading of bytes from files with arbitrary compression and for closing underlying resources.
// Reader implements the ByteReader interface by wrapping around a subordinate ByteReader.
type Reader struct {
	Reader ByteReader // the instance of ByteReader used for reading bytes
	Closer io.Closer  // Used for closing readers with footer metadata, e.g., gzip.  Not always needed, e.g., snappy
	File   *os.File   // underlying file, if any
}

// Read reads a maximum len(p) bytes from the reader and returns an error, if any.
func (r *Reader) Read(p []byte) (n int, err error) {

	if r.Reader != nil {
		return r.Reader.Read(p)
	}

	return 0, nil
}

// ReadByte returns a single byte from the underlying reader.
func (r *Reader) ReadByte() (byte, error) {

	if r.Reader != nil {
		return r.Reader.ReadByte()
	}

	return 0, nil
}

// Read returns all bytes up to an including the first occurence of the delimiter "delim" and an error, if any.
func (r *Reader) ReadBytes(delim byte) ([]byte, error) {

	if r.Reader != nil {
		return r.Reader.ReadBytes(delim)
	}

	return make([]byte, 0), nil
}

// ReadFirst is not implemented by Reader
func (r *Reader) ReadFirst() (byte, error) {
	return byte(0), errors.New("ReadFirst is not implemented by Reader")
}

// ReadAt is not implemented by Reader
func (r *Reader) ReadAt(i int) (byte, error) {
	return byte(0), errors.New("ReadAt is not implemented by Reader")
}

// ReadRange is not implemented by Reader
func (r *Reader) ReadRange(start int, end int) ([]byte, error) {
	return make([]byte, 0), errors.New("ReadRange is not implemented by Reader")
}

// ReadAll is not implemented by Reader
func (r *Reader) ReadAll() ([]byte, error) {
	return ioutil.ReadAll(r.Reader)
}

// Close closes the Closer and the underlying *os.File if not nil.
func (r *Reader) Close() error {

	if r.Closer != nil {
		err := r.Closer.Close()
		if err != nil {
			return errors.Wrap(err, "Error closing read closer.")
		}
	}

	if r.File != nil {
		err := r.File.Close()
		if err != nil {
			return errors.Wrap(err, "Error closing file.")
		}
	}

	return nil
}
