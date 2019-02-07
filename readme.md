
# webp-animation

Golang binding to **libwebp** and **giflib** library (https://github.com/webmproject/libwebp, https://github.com/rcancro/giflib)

This package implements functions (not all) which are required to convert GIF to WEBP animation, especially the functions that were used for this example: https://github.com/webmproject/libwebp/blob/0.6.1/examples/gif2webp.c

  
  

## Installation

Get package: ```go get github.com/sizeofint/webp-animation```

  

dependencies:

- libraries: libwebp (0.6.1), giflib 5.1


Package provides prebuilt dependencies for Linux, but if you want rebuild by your own, run [deps/build-deps-linux.sh](deps/build-deps-linux.sh)  
  


## Implementations

[gif-to-webp](https://github.com/sizeofint/gif-to-webp  "Golang convert GIF to WEBP") library based on this library