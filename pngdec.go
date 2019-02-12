package webpanim

// #cgo linux CFLAGS: -I${SRCDIR}/deps/linux/include
// #cgo linux CXXFLAGS: -I${SRCDIR}/deps/linux/include
// #cgo linux LDFLAGS: -L${SRCDIR}/deps/linux/lib -lpng -lwebp -lz
// #include "webp.hpp"
// #include "pngdec.h"
// #include "metadata.h"
import "C"
import "unsafe"

type Metadata C.Metadata

func ReadPNG(buffer []byte, webPPicture *WebPPicture, keepAlpha int, metadata *Metadata) int {
	return int(C.ReadPNG(
		(*C.uint8_t)(C.CBytes(buffer)),
		(C.size_t)(len(buffer)),
		(*C.WebPPicture)(unsafe.Pointer(webPPicture)),
		(C.int)(keepAlpha),
		(*C.Metadata)(unsafe.Pointer(metadata)),
	))
}
