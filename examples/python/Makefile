# =================================================================
#
# Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
# Released as open source under the MIT License.  See LICENSE file.
#
# =================================================================

build:
	go build -o reader.so -buildmode=c-shared github.com/spatialcurrent/go-reader/plugins/reader
run:
	LD_LIBRARY_PATH=. python test.py
clean:
	rm -f reader.h reader.so
