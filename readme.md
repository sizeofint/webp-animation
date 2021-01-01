⚠Before using this library read text below👇  
This library is direct binding to libwebp 0.6, originally created to convert gif animations to webp.  
Recently I created more generic and golang language friendly version to make webp animations supporting image.Image interface which will help to create webp animation not only from gif animations but  also from any image source that you will get in image.Image interface. so I highly recommend to use the [new version](https://github.com/sizeofint/webpanimation)
# webp-animation

Golang binding to **libwebp**, **libpng** and **giflib** library

This package implements functions (not all) which are required to convert GIF to WEBP animation, especially the functions that were used for this example: https://github.com/webmproject/libwebp/blob/0.6.1/examples/gif2webp.c.   
Binding to libpng helps to create new webp animation from png images.


## Installation

Get package: ```go get github.com/sizeofint/webp-animation```

dependencies:

- libraries: libwebp 0.6.1, giflib 5.1, zlib 1.2.11(required for libpng), libpng 1.6.36


Package provides prebuilt dependencies for Linux, but if you want to rebuild by your own, run [deps/build-deps-linux.sh](deps/build-deps-linux.sh)  
  


## Implementations

[gif-to-webp](https://github.com/sizeofint/gif-to-webp  "Golang convert gif to webp") Convert gif to webp 

[png-to-webp-animation](https://github.com/sizeofint/png-to-webp-animation  "Create webp animation from png files") Create webp animation from png files
