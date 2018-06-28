package main

import (
  "flag"
  "io"
  "io/ioutil"
  "fmt"
  "time"
  "os"
)

import (
	"github.com/pkg/errors"
)

import (
  "github.com/spatialcurrent/go-reader/reader"
)

var GO_READER_VERSION = "0.0.1"

func main () {

  start := time.Now()

	var input_uri string
	var input_format string
  var input_compression_algorithm string

  var endianness string

  var header bool
  var delim string

  var verbose bool
	var version bool
	var help bool

  flag.StringVar(&input_uri, "input_uri", "", "\"stdin\" or uri to input file")
	flag.StringVar(&input_format, "input_format", "json", "Input format: json, or yaml.")
	flag.StringVar(&input_compression_algorithm, "input_alg", "none", "Stream input compression algorithm for nodes, using: snappy, gzip, or none.")

  flag.BoolVar(&header, "header", false, "Object size header")
  flag.StringVar(&delim, "delim", "", "Delimiter between objects, e.g., \\n")
  flag.StringVar(&endianness, "endianness", "", "Encode numbers using a big-endian byte order else encodes using littl-endian byte order.")

  flag.BoolVar(&version, "version", false, "Prints version to stdout")
  flag.BoolVar(&help, "help", false, "Print help")

  flag.Parse()

  if help {
    fmt.Println("Usage: go-reader -input_uri INPUT_URI -endianness [little|big] [-filter INPUT] [-verbose] [-version] [-help] [A=1] [B=2]")
    fmt.Println("Options:")
    flag.PrintDefaults()
    os.Exit(0)
  } else if len(os.Args) == 1 {
    fmt.Println("Error: Provided no arguments.")
    fmt.Println("Run \"go-reader -help\" for more information.")
    os.Exit(0)
  } else if len(os.Args) == 2 && os.Args[1] == "help" {
    fmt.Println("Usage: go-reader -input_uri INPUT_URI -endianness [little|big] [-filter INPUT] [-verbose] [-version] [-help] [A=1] [B=2]")
    fmt.Println("Options:")
    flag.PrintDefaults()
    os.Exit(0)
  }

  if version {
    fmt.Println(GO_READER_VERSION)
    os.Exit(0)
  }

  var input_reader reader.ByteReader
  if input_uri == "stdin" {
    input_bytes, err := ioutil.ReadAll(os.Stdin)
    if err != nil {
      fmt.Println("error reading from stdin")
      os.Exit(1)
    }
    r, err := reader.OpenBytes(input_bytes, input_compression_algorithm)
    if err != nil {
      fmt.Println(errors.Wrap(err, "Error opening file \""+input_uri+"\""))
      os.Exit(1)
    }
    input_reader = r
  } else {
    r, err := reader.OpenFile(input_uri, input_compression_algorithm)
    if err != nil {
      fmt.Println(errors.Wrap(err, "Error opening file \""+input_uri+"\""))
      os.Exit(1)
    }
    input_reader = r
  }


  if header {
    fmt.Println("Header not implemented yet")
    os.Exit(1)
  } else if len(delim) > 0 {

    for {
      b, err := input_reader.ReadBytes([]byte(delim)[0])
      if err != nil {
        if err != io.EOF {
          fmt.Println(errors.Wrap(err, "Error reading bytes from file"))
          os.Exit(1)
        }
      }
      if len(b) > 0 {
        fmt.Println(string(b))
      }
      if err != nil && err == io.EOF {
        break
      }

    }

  }

  elapsed := time.Since(start)
  if verbose {
    fmt.Println("Done in " + elapsed.String())
  }

}
