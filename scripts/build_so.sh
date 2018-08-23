#!/bin/bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
DEST=$(realpath ${1:-$DIR/../bin})

mkdir -p $DEST

echo "******************"
echo "Formatting $(realpath $DIR/../reader)"
cd $DIR/../reader
go fmt
echo "Formatting $(realpath $DIR/../plugins/reader)"
cd $DIR/../plugins/reader
go fmt
echo "Done formatting."
echo "******************"
echo "Building Shared Object (*.so) for go-reader"
cd $DEST
go build -o reader.so -buildmode=c-shared github.com/spatialcurrent/go-reader/plugins/reader
if [[ "$?" != 0 ]] ; then
    echo "Error Building Shared Object (*.so) for go-reader"
    exit 1
fi
echo "Shared Object (*.so) built at $DEST"
