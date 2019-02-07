#!/bin/sh

set -e

BASEDIR=$(cd $(dirname "$0") && pwd)
PREFIX="$BASEDIR/linux"
BUILDDIR="$BASEDIR/build"
SRCDIR="$BASEDIR/dep-source"
export CFLAGS="$CFLAGS -fPIC"
export CXXFLAGS="$CXXFLAGS -fPIC"

mkdir -p $PREFIX/include
mkdir -p $PREFIX/lib

rm -rf $BUILDDIR
mkdir -p $BUILDDIR

rm -rf libwebp
rm -rf giflib
rm -rf $SRCDIR



mkdir -p $SRCDIR
echo "Downloading sources..."
curl -L -o $SRCDIR/giflib-5.1.2.tar.gz https://github.com/sizeofint/giflib/archive/5.1.2.tar.gz
curl -L -o $SRCDIR/libwebp-v0.6.1.tar.gz https://github.com/sizeofint/libwebp/archive/v0.6.1.tar.gz


echo "Installing giflib from source..."
mkdir -p $BASEDIR/giflib
tar -xzf $SRCDIR/giflib-5.1.2.tar.gz -C $BASEDIR/giflib --strip-components 1
###cd $BASEDIR/giflib
###./autogen.sh  --prefix=$PREFIX --disable-shared
mkdir -p $BUILDDIR/giflib
cd $BUILDDIR/giflib
/$BASEDIR/giflib/autogen.sh  --prefix=$PREFIX --disable-shared
make
make install

echo "Installing libwebp from source..."
mkdir -p $BASEDIR/libwebp
tar -xzf $SRCDIR/libwebp-v0.6.1.tar.gz -C $BASEDIR/libwebp --strip-components 1
cd $BASEDIR/libwebp
./autogen.sh
mkdir -p $BUILDDIR/libwebp
cd $BUILDDIR/libwebp
$BASEDIR/libwebp/configure --prefix=$PREFIX --disable-shared --enable-static --enable-libwebpmux --with-gifincludedir=$PREFIX/include --with-giflibdir=$PREFIX/lib
make
make install

echo "Done."
