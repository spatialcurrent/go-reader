#!/bin/bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
DEST=$(realpath ${1:-$DIR/../bin})

mkdir -p $DEST

echo "******************"
echo "Formatting $(realpath $DIR/../reader)"
cd $DIR/../reader
go fmt
echo "Formatting $(realpath $DIR/../cmd/reader.js)"
cd $DIR/../cmd/reader.js
go fmt
echo "Done formatting."
echo "******************"
echo "Building Javascript artifact for go-reader"
cd $DEST
gopherjs build -o dfl.js github.com/spatialcurrent/go-reader/cmd/reader.js
if [[ "$?" != 0 ]] ; then
    echo "Error building Javascript for go-reader"
    exit 1
fi
echo "Executable built at $DEST"
