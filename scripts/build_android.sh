#!/bin/bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
DEST=$(realpath ${1:-$DIR/../bin})

mkdir -p $DEST

echo "******************"
echo "Formatting $(realpath $DIR/../reader)"
cd $DIR/../reader
go fmt
echo "Done formatting."
echo "******************"
echo "Building AAR for go-reader"
cd $DEST
gomobile bind -target android -javapkg=com.spatialcurrent -o reader.aar github.com/spatialcurrent/go-reader/reader
if [[ "$?" != 0 ]] ; then
    echo "Error building program for go-reader"
    exit 1
fi
echo "Executable built at $DEST"
