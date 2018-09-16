// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

import (
	"github.com/pkg/errors"
)

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

import (
	"github.com/spatialcurrent/go-reader/reader"
)

func connect_to_aws(aws_access_key_id string, aws_secret_access_key string, aws_session_token string, aws_region string) *session.Session {
	aws_session := session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Credentials: credentials.NewStaticCredentials(aws_access_key_id, aws_secret_access_key, aws_session_token),
			MaxRetries:  aws.Int(3),
			Region:      aws.String(aws_region),
		},
	}))
	return aws_session
}

func main() {

	start := time.Now()

	var aws_default_region string
	var aws_access_key_id string
	var aws_secret_access_key string
	var aws_session_token string

	var input_uri string
	var input_compression string
	var input_reader_buffer_size int

	var verbose bool
	var version bool
	var help bool

	flag.StringVar(&aws_default_region, "aws_default_region", "", "Defaults to value of environment variable AWS_DEFAULT_REGION.")
	flag.StringVar(&aws_access_key_id, "aws_access_key_id", "", "Defaults to value of environment variable AWS_ACCESS_KEY_ID")
	flag.StringVar(&aws_secret_access_key, "aws_secret_access_key", "", "Defaults to value of environment variable AWS_SECRET_ACCESS_KEY.")
	flag.StringVar(&aws_session_token, "aws_session_token", "", "Defaults to value of environment variable AWS_SESSION_TOKEN.")

	flag.StringVar(&input_uri, "uri", "", "\"stdin\" or uri to input file")
	flag.StringVar(&input_compression, "alg", "none", "Stream input compression algorithm for nodes, using: bzip2, gzip, snappy, or none.")
	flag.IntVar(&input_reader_buffer_size, "buffer_size", 4096, "The input reader buffer size") // default from https://golang.org/src/bufio/bufio.go

	flag.BoolVar(&version, "version", false, "Prints version to stdout")
	flag.BoolVar(&help, "help", false, "Print help")

	flag.Parse()

	if len(aws_default_region) == 0 {
		aws_default_region = os.Getenv("AWS_DEFAULT_REGION")
	}
	if len(aws_access_key_id) == 0 {
		aws_access_key_id = os.Getenv("AWS_ACCESS_KEY_ID")
	}
	if len(aws_secret_access_key) == 0 {
		aws_secret_access_key = os.Getenv("AWS_SECRET_ACCESS_KEY")
	}
	if len(aws_session_token) == 0 {
		aws_session_token = os.Getenv("AWS_SESSION_TOKEN")
	}

	if help {
		fmt.Println("Usage: reader -uri URI [-alg [bzip2|gzip|snappy|none]] [-verbose] [-version]")
		fmt.Println("Options:")
		flag.PrintDefaults()
		os.Exit(0)
	} else if len(os.Args) == 1 {
		fmt.Println("Error: Provided no arguments.")
		fmt.Println("Run \"reader -help\" for more information.")
		os.Exit(0)
	} else if len(os.Args) == 2 && os.Args[1] == "help" {
		fmt.Println("Usage: reader -input_uri INPUT_URI -endianness [little|big] [-filter INPUT] [-verbose] [-version] [-help] [A=1] [B=2]")
		fmt.Println("Options:")
		flag.PrintDefaults()
		os.Exit(0)
	}

	if version {
		fmt.Println(reader.VERSION)
		os.Exit(0)
	}

	var aws_session *session.Session
	var s3_client *s3.S3

	if strings.HasPrefix(input_uri, "s3://") {
		aws_session = connect_to_aws(aws_access_key_id, aws_secret_access_key, aws_session_token, aws_default_region)
		s3_client = s3.New(aws_session)
	}

	input_reader, _, err := reader.OpenResource(input_uri, input_compression, input_reader_buffer_size, false, s3_client)
	if err != nil {
		log.Fatal(errors.Wrap(err, "error opening resource from uri "+input_uri))
	}

	input_bytes, err := input_reader.ReadAll()
	if err != nil {
		log.Fatal(errors.Wrap(err, "Error reading from resource"))
	}

	input_string := string(input_bytes)

	fmt.Println(input_string)

	elapsed := time.Since(start)
	if verbose {
		fmt.Println("Done in " + elapsed.String())
	}

}
