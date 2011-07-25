#!/bin/sh

DIRS="appengine appengine_internal"

# check for mercurial
which hg >/dev/null 2>&1
if [ $? -ne 0 ]; then
    echo "Mercurial not found in path $PATH." >&2
    exit 1
fi

# check parameters
if [ $# -ne 2 ]; then
    echo "Usage:  $0 <version> <commit>" >&2
    exit 2
fi

VERSION=$1
COMMIT=$2

# where to store the result
SOURCES=$(pwd)/gaesources

# check if available
if [ -d $SOURCES/$VERSION ]; then
    echo "Version $VERSION already fetched." >&2
    exit 3
fi

# generate temporary directory
echo "Creating temporary directory..."
TMP=/tmp/gaesources
rm -rf $TMP
mkdir $TMP || exit 4

# checkout
echo "Checking out given commit for the version..."
hg clone --rev $COMMIT https://code.google.com/p/appengine-go/ $TMP || exit 5

# copy
echo "Copying selected sources to plugin directory..."
mkdir -p $SOURCES/$VERSION || exit 6
for dir in $DIRS; do
    cp -r $TMP/$dir $SOURCES/$VERSION/ || exit 7
done
