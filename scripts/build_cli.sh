#!/bin/bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
DEST=$(realpath ${1:-$DIR/../bin})

mkdir -p $DEST

echo "******************"
echo "Formatting $(realpath $DIR/../reader)"
cd $DIR/../reader
go fmt
echo "Formatting $(realpath $DIR/../cmd/reader)"
cd $DIR/../cmd/reader
go fmt
echo "Done formatting."
echo "******************"
echo "Building program for go-reader"
cd $DEST
for GOOS in darwin linux windows; do
  GOOS=${GOOS} GOARCH=amd64 go build -o "reader_${GOOS}_amd64" github.com/spatialcurrent/go-reader/cmd/reader
done
if [[ "$?" != 0 ]] ; then
    echo "Error building program for go-reader"
    exit 1
fi
echo "Executables built at $DEST"
