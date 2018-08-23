# =================================================================
#
# Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
# Released as open source under the MIT License.  See LICENSE file.
#
# =================================================================

build_so:
	go build -o reader.so -buildmode=c-shared github.com/spatialcurrent/go-reader/plugins/reader
build_c:
	mkdir -p bin && gcc -o bin/reader_test_c test.c -L. -l:reader.so
build: build_so build_c
run:
	LD_LIBRARY_PATH=. bin/reader_test_c
clean:
	rm -f reader.h reader.so bin/reader_test_c
