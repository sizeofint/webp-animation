package webpanim

// #cgo pkg-config: libwebp libwebpmux
// #cgo linux CFLAGS: -I/usr/include
// #cgo linux LDFLAGS: -L/usr/lib/x86_64-linux-gnu -lgif
// #include "webp.hpp"
// #include "gifdec.h"
// #include "gifdec.hpp"
import "C"
import (
	"unsafe"
)

type GIFDisposeMethod C.GIFDisposeMethod
type ColorMapObject C.ColorMapObject
type GifImageDesc C.GifImageDesc
type GifFileType C.GifFileType
type GifRecordType C.GifRecordType
type GIFFrameRect C.GIFFrameRect

const TransparentIndex = -1

const (
	ContinueExtFuncCode    = 0x00
	CommentExtFuncCode     = 0xfe
	GraphicsExtFuncCode    = 0xf9
	PlaintextExtFuncCode   = 0x01
	ApplicationExtFuncCode = 0xff
)

const (
	GifError = 0
	GifOk    = 1
)

const (
	UndefinedRecordType  = GifRecordType(C.UNDEFINED_RECORD_TYPE)
	ScreenDescRecordType = GifRecordType(C.SCREEN_DESC_RECORD_TYPE)
	ImageDescRecordType  = GifRecordType(C.IMAGE_DESC_RECORD_TYPE)
	ExtensionRecordType  = GifRecordType(C.EXTENSION_RECORD_TYPE)
	TerminateRecordType  = GifRecordType(C.TERMINATE_RECORD_TYPE)
)

const (
	GifDisposeNone            = GIFDisposeMethod(C.GIF_DISPOSE_NONE)
	GifDisposeBackground      = GIFDisposeMethod(C.GIF_DISPOSE_BACKGROUND)
	GifDisposeRestorePrevious = GIFDisposeMethod(C.GIF_DISPOSE_RESTORE_PREVIOUS)
)

func (imgDesc *GifImageDesc) SetLeft(v int) {

	((*C.GifImageDesc)(imgDesc)).Left = (C.GifWord)(v)
}

func (imgDesc *GifImageDesc) SetTop(v int) {

	((*C.GifImageDesc)(imgDesc)).Top = (C.GifWord)(v)
}

func (imgDesc *GifImageDesc) SetWidth(v int) {

	((*C.GifImageDesc)(imgDesc)).Width = (C.GifWord)(v)
}

func (imgDesc *GifImageDesc) SetHeight(v int) {

	((*C.GifImageDesc)(imgDesc)).Height = (C.GifWord)(v)
}

func (imgDesc *GifImageDesc) GetWidth() int {

	return int(((*C.GifImageDesc)(imgDesc)).Width)
}

func (imgDesc *GifImageDesc) GetHeight() int {

	return int(((*C.GifImageDesc)(imgDesc)).Height)
}

func (g *GifFileType) GetImage() GifImageDesc {
	return (GifImageDesc)(((*C.GifFileType)(g)).Image)
}

func (g *GifFileType) SetSWidth(w int) {
	((*C.GifFileType)(g)).SWidth = (C.GifWord)(w)
}

func (g *GifFileType) SetSHeight(h int) {
	((*C.GifFileType)(g)).SHeight = (C.GifWord)(h)
}

func (g *GifFileType) GetSWidth() int {
	return int(((*C.GifFileType)(g)).SWidth)
}

func (g *GifFileType) GetSHeight() int {
	return int(((*C.GifFileType)(g)).SHeight)
}

func (g *GifFileType) GetSColorMap() *ColorMapObject {
	return (*ColorMapObject)(((*C.GifFileType)(g)).SColorMap)
}

func (g *GifFileType) GetSBackGroundColor() int {
	return int(((*C.GifFileType)(g)).SBackGroundColor)
}

func GIFGetBackgroundColor(colorMapObject *ColorMapObject, bgcolorIndex int, transparentIndex int, bgcolor *uint32) {
	C.GIFGetBackgroundColor(
		(*C.ColorMapObject)(unsafe.Pointer(colorMapObject)),
		(C.int)(bgcolorIndex),
		(C.int)(transparentIndex),
		(*C.uint32_t)(bgcolor),
	)
}

func DGifOpenFileName(gifFileName string, errCode *int) *GifFileType {
	return (*GifFileType)(C.DGifOpenFileName(
		C.CString(gifFileName),
		(*C.int)(unsafe.Pointer(errCode)),
	))
}

func DGifGetImageDesc(gifFileType *GifFileType) int {
	return int(C.DGifGetImageDesc(
		(*C.GifFileType)(unsafe.Pointer(gifFileType)),
	))
}

func DGifGetRecordType(gifFileType *GifFileType, gifRecordType *GifRecordType) int {
	return int(C.DGifGetRecordType(
		(*C.GifFileType)(unsafe.Pointer(gifFileType)),
		(*C.GifRecordType)(unsafe.Pointer(gifRecordType)),
	))
}

func GIFClearPic(webPPicture *WebPPicture, gIFFrameRect *GIFFrameRect) {
	C.GIFClearPic(
		(*C.WebPPicture)(unsafe.Pointer(webPPicture)),
		(*C.GIFFrameRect)(unsafe.Pointer(gIFFrameRect)),
	)
}

func GIFReadFrame(gifFileType *GifFileType, transparentIndex int, gIFFrameRect *GIFFrameRect, webPPicture *WebPPicture) int {
	return int(C.GIFReadFrame(
		(*C.GifFileType)(unsafe.Pointer(gifFileType)),
		(C.int)(transparentIndex),
		(*C.GIFFrameRect)(unsafe.Pointer(gIFFrameRect)),
		(*C.WebPPicture)(unsafe.Pointer(webPPicture)),
	))
}

func GIFBlendFrames(webPPictureSrc *WebPPicture, gIFFrameRect *GIFFrameRect, webPPictureDst *WebPPicture) {
	C.GIFBlendFrames(
		(*C.WebPPicture)(unsafe.Pointer(webPPictureSrc)),
		(*C.GIFFrameRect)(unsafe.Pointer(gIFFrameRect)),
		(*C.WebPPicture)(unsafe.Pointer(webPPictureDst)),
	)
}

func GIFDisposeFrame(gIFDisposeMethod GIFDisposeMethod, gIFFrameRect *GIFFrameRect, webPPicturePrev *WebPPicture, webPPictureCurr *WebPPicture) {
	C.GIFDisposeFrame(
		(C.GIFDisposeMethod)(gIFDisposeMethod),
		(*C.GIFFrameRect)(unsafe.Pointer(gIFFrameRect)),
		(*C.WebPPicture)(unsafe.Pointer(webPPicturePrev)),
		(*C.WebPPicture)(unsafe.Pointer(webPPictureCurr)),
	)
}

func GIFCopyPixels(webPPictureSrc *WebPPicture, webPPictureDst *WebPPicture) {
	C.GIFCopyPixels(
		(*C.WebPPicture)(unsafe.Pointer(webPPictureSrc)),
		(*C.WebPPicture)(unsafe.Pointer(webPPictureDst)),
	)
}

func DGifGetExtension(gifFileType *GifFileType, gifExtCode *int, gifByteType *[]byte) int {

	pointerToArrayPointer := (**C.GifByteType)(unsafe.Pointer(C.CBytes(*gifByteType)))

	res := int(C.DGifGetExtension(
		(*C.GifFileType)(unsafe.Pointer(gifFileType)),
		(*C.int)(unsafe.Pointer(gifExtCode)),
		pointerToArrayPointer,
	))

	arr := unsafe.Pointer((*C.GifByteType)(*pointerToArrayPointer))
	if arr != nil {
		size := int(C.GoBytes(arr, (C.int)(1))[0])
		*gifByteType = C.GoBytes(arr, (C.int)((size + 1)))
	} else {
		*gifByteType = nil
	}

	return res
}

func GIFReadGraphicsExtension(gifByteType []byte, duration *int, gIFDisposeMethod *GIFDisposeMethod, transparentIndex *int) int {
	return int(C.GIFReadGraphicsExtension(
		(*C.GifByteType)(C.CBytes(gifByteType)),
		(*C.int)(unsafe.Pointer(duration)),
		(*C.GIFDisposeMethod)(unsafe.Pointer(gIFDisposeMethod)),
		(*C.int)(unsafe.Pointer(transparentIndex)),
	))
}

func GIFReadLoopCount(gifFileType *GifFileType, gifByteType []byte, loopCount *int) int {

	cArray := C.malloc(C.size_t(len(gifByteType)))

	a := (*[256]C.GifByteType)(cArray)

	for index, value := range gifByteType {
		a[index] = (C.GifByteType)(value)
	}

	res := int(C.GIFReadLoopCount(
		(*C.GifFileType)(unsafe.Pointer(gifFileType)),
		(**C.GifByteType)(unsafe.Pointer(&cArray)),
		(*C.int)(unsafe.Pointer(loopCount)),
	))

	return res
}

func GIFReadMetadata(gifFileType *GifFileType, gifByteType []byte, webPData *WebPData) int {
	return int(C.GIFReadMetadata(
		(*C.GifFileType)(unsafe.Pointer(gifFileType)),
		(**C.GifByteType)(unsafe.Pointer(C.CBytes(gifByteType))),
		(*C.WebPData)(unsafe.Pointer(webPData)),
	))
}

func DGifGetExtensionNext(gifFileType *GifFileType, gifByteType *[]byte) int {
	pointerToArrayPointer := (**C.GifByteType)(unsafe.Pointer(C.CBytes(*gifByteType)))

	res := int(C.DGifGetExtensionNext(
		(*C.GifFileType)(unsafe.Pointer(gifFileType)),
		pointerToArrayPointer,
	))

	arr := unsafe.Pointer((*C.GifByteType)(*pointerToArrayPointer))
	if arr != nil {
		size := int(C.GoBytes(arr, (C.int)(1))[0])
		*gifByteType = C.GoBytes(arr, (C.int)((size + 1)))
		if len(*gifByteType) == 0 {
			*gifByteType = nil
		}
	} else {
		*gifByteType = nil
	}

	return res
}

func GIFDisplayError(gifFileType *GifFileType, gifError int) {
	C.GIFDisplayError(
		(*C.GifFileType)(unsafe.Pointer(gifFileType)),
		(C.int)(gifError),
	)
}

func DGifCloseFile(gifFileType *GifFileType, errorCode *int) int {
	return int(C.DGifCloseFile(
		(*C.GifFileType)(unsafe.Pointer(gifFileType)),
		(*C.int)(unsafe.Pointer(errorCode)),
	))
}

func Gif(buffer []byte, errorCode *int) *GifFileType {

	return (*GifFileType)(C.gif(
		(*C.GifByteType)(C.CBytes(buffer)),
		(C.int)(len(buffer)),
		(*C.int)(unsafe.Pointer(errorCode)),
	))
}
