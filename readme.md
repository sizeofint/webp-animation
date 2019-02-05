# webp-animation
Binding to **libwebp** and **giflib** library (https://github.com/webmproject/libwebp, https://github.com/rcancro/giflib)
This package implements all functions (not all) which are required to convert GIF to WEBP animation, especially the functions that were used for this example: https://github.com/webmproject/libwebp/blob/0.6.1/examples/gif2webp.c


## Installation
Get package: ```go get github.com/sizeofint/webp-animation```

dependencies:
 - libraries: libwebp (0.6.1), giflib 5.1
 - tools: pkg-config

In gifdec.go replace with correct path values depending on your system configuration if needed:
```
// #cgo linux CFLAGS: -I/usr/include
// #cgo linux LDFLAGS: -L/usr/lib/x86_64-linux-gnu -lgif
```

