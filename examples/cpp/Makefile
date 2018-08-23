# =================================================================
#
# Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
# Released as open source under the MIT License.  See LICENSE file.
#
# =================================================================

build_so:
	go build -o reader.so -buildmode=c-shared github.com/spatialcurrent/go-reader/plugins/reader
build_cpp:
	mkdir -p bin && g++ -o bin/reader_test_cpp test.cpp -L. -l:reader.so
build: build_so build_cpp
run:
	LD_LIBRARY_PATH=. bin/reader_test_cpp
clean:
	rm -f reader.h reader.so bin/reader_test_cpp
