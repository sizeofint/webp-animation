package webpanim

/*
#cgo pkg-config: libwebp libwebpmux
#include "webp.hpp"
*/
import "C"
import "unsafe"

type WebPImageHint int
type WebPMuxError int

const (
	WebpMuxOk              = WebPMuxError(C.WEBP_MUX_OK)
	WebpMuxNotFound        = WebPMuxError(C.WEBP_MUX_NOT_FOUND)
	WebpMuxInvalidArgument = WebPMuxError(C.WEBP_MUX_INVALID_ARGUMENT)
	WebpMuxBadData         = WebPMuxError(C.WEBP_MUX_BAD_DATA)
	WebpMuxMemoryError     = WebPMuxError(C.WEBP_MUX_MEMORY_ERROR)
	WebpMuxNotEnoughData   = WebPMuxError(C.WEBP_MUX_NOT_ENOUGH_DATA)
)

const (
	WebpHintDefault = WebPImageHint(C.WEBP_HINT_DEFAULT)
	WebpHintPicture = WebPImageHint(C.WEBP_HINT_PICTURE)
	WebpHintPhoto   = WebPImageHint(C.WEBP_HINT_PHOTO)
	WebpHintGraph   = WebPImageHint(C.WEBP_HINT_GRAPH)
	WebpHintLast    = WebPImageHint(C.WEBP_HINT_LAST)
)

type WebPAnimEncoder C.WebPAnimEncoder
type WebPMux C.WebPMux
type WebPMuxAnimParams C.WebPMuxAnimParams
type WebPPicture C.WebPPicture
type WebPConfig C.WebPConfig
type WebPAnimEncoderOptions C.WebPAnimEncoderOptions
type WebPData C.WebPData

func (webpCfg *WebPConfig) SetLossless(v int) {
	((*C.WebPConfig)(webpCfg)).lossless = (C.int)(v)
}

func (webpCfg WebPConfig) GetLossless() int {
	return int(((C.WebPConfig)(webpCfg)).lossless)
}

func (webpCfg *WebPConfig) SetMethod(v int) {
	((*C.WebPConfig)(webpCfg)).method = (C.int)(v)
}

func (webpCfg *WebPConfig) SetImageHint(v WebPImageHint) {
	((*C.WebPConfig)(webpCfg)).image_hint = (C.WebPImageHint)(v)
}

func (webpCfg *WebPConfig) SetTargetSize(v int) {
	((*C.WebPConfig)(webpCfg)).target_size = (C.int)(v)
}

func (webpCfg *WebPConfig) SetTargetPSNR(v float32) {
	((*C.WebPConfig)(webpCfg)).target_PSNR = (C.float)(v)
}

func (webpCfg *WebPConfig) SetSegments(v int) {
	((*C.WebPConfig)(webpCfg)).segments = (C.int)(v)
}

func (webpCfg *WebPConfig) SetSnsStrength(v int) {
	((*C.WebPConfig)(webpCfg)).sns_strength = (C.int)(v)
}

func (webpCfg *WebPConfig) SetFilterStrength(v int) {
	((*C.WebPConfig)(webpCfg)).filter_strength = (C.int)(v)
}

func (webpCfg *WebPConfig) SetFilterSharpness(v int) {
	((*C.WebPConfig)(webpCfg)).filter_sharpness = (C.int)(v)
}

func (webpCfg *WebPConfig) SetAutofilter(v int) {
	((*C.WebPConfig)(webpCfg)).autofilter = (C.int)(v)
}

func (webpCfg *WebPConfig) SetAlphaCompression(v int) {
	((*C.WebPConfig)(webpCfg)).alpha_compression = (C.int)(v)
}

func (webpCfg *WebPConfig) SetAlphaFiltering(v int) {
	((*C.WebPConfig)(webpCfg)).alpha_filtering = (C.int)(v)
}

func (webpCfg *WebPConfig) SetPass(v int) {
	((*C.WebPConfig)(webpCfg)).pass = (C.int)(v)
}

func (webpCfg *WebPConfig) SetShowCompressed(v int) {
	((*C.WebPConfig)(webpCfg)).show_compressed = (C.int)(v)
}

func (webpCfg *WebPConfig) SetPreprocessing(v int) {
	((*C.WebPConfig)(webpCfg)).preprocessing = (C.int)(v)
}

func (webpCfg *WebPConfig) SetPartitions(v int) {
	((*C.WebPConfig)(webpCfg)).partitions = (C.int)(v)
}

func (webpCfg *WebPConfig) SetPartitionLimit(v int) {
	((*C.WebPConfig)(webpCfg)).partition_limit = (C.int)(v)
}

func (webpCfg *WebPConfig) SetEmulateJpegSize(v int) {
	((*C.WebPConfig)(webpCfg)).emulate_jpeg_size = (C.int)(v)
}

func (webpCfg *WebPConfig) SetThreadLevel(v int) {
	((*C.WebPConfig)(webpCfg)).thread_level = (C.int)(v)
}

func (webpCfg *WebPConfig) SetLowMemory(v int) {
	((*C.WebPConfig)(webpCfg)).low_memory = (C.int)(v)
}

func (webpCfg *WebPConfig) SetNearLossless(v int) {
	((*C.WebPConfig)(webpCfg)).near_lossless = (C.int)(v)
}

func (webpCfg *WebPConfig) SetExact(v int) {
	((*C.WebPConfig)(webpCfg)).exact = (C.int)(v)
}

func (webpCfg *WebPConfig) SetUseDeltaPalette(v int) {
	((*C.WebPConfig)(webpCfg)).use_delta_palette = (C.int)(v)
}

func (webpCfg *WebPConfig) SetUseSharpYuv(v int) {
	((*C.WebPConfig)(webpCfg)).use_sharp_yuv = (C.int)(v)
}

func (webpCfg *WebPConfig) SetAlphaQuality(v int) {
	((*C.WebPConfig)(webpCfg)).alpha_quality = (C.int)(v)
}

func (webpCfg *WebPConfig) SetFilterType(v int) {
	((*C.WebPConfig)(webpCfg)).filter_type = (C.int)(v)
}

func (webpCfg *WebPConfig) SetQuality(v float32) {
	((*C.WebPConfig)(webpCfg)).quality = (C.float)(v)
}

func (wpd WebPData) GetBytes() []byte {
	return C.GoBytes(unsafe.Pointer(((C.WebPData)(wpd)).bytes), (C.int)(((C.WebPData)(wpd)).size))
}

func (encOptions *WebPAnimEncoderOptions) GetAnimParams() WebPMuxAnimParams {
	return WebPMuxAnimParams(((*C.WebPAnimEncoderOptions)(encOptions)).anim_params)
}

func (encOptions *WebPAnimEncoderOptions) SetAnimParams(v WebPMuxAnimParams) {
	((*C.WebPAnimEncoderOptions)(encOptions)).anim_params = (C.WebPMuxAnimParams)(v)
}

func (encOptions *WebPAnimEncoderOptions) SetMinimizeSize(v int) {
	((*C.WebPAnimEncoderOptions)(encOptions)).minimize_size = (C.int)(v)
}

func (encOptions *WebPAnimEncoderOptions) SetKmin(v int) {
	((*C.WebPAnimEncoderOptions)(encOptions)).kmin = (C.int)(v)
}

func (encOptions *WebPAnimEncoderOptions) SetKmax(v int) {
	((*C.WebPAnimEncoderOptions)(encOptions)).kmax = (C.int)(v)
}

func (encOptions *WebPAnimEncoderOptions) SetAllowMixed(v int) {
	((*C.WebPAnimEncoderOptions)(encOptions)).allow_mixed = (C.int)(v)
}

func (encOptions *WebPAnimEncoderOptions) SetVerbose(v int) {
	((*C.WebPAnimEncoderOptions)(encOptions)).verbose = (C.int)(v)
}

func (wmap WebPMuxAnimParams) GetBgcolor() uint32 {
	return uint32((C.WebPMuxAnimParams)(wmap).bgcolor)
}

func (wmap *WebPMuxAnimParams) SetBgcolor(v uint32) {
	(*C.WebPMuxAnimParams)(wmap).bgcolor = (C.uint32_t)(v)
}

func (wmap *WebPMuxAnimParams) SetLoopCount(v int) {
	(*C.WebPMuxAnimParams)(wmap).loop_count = (C.int)(v)
}

func (wpp *WebPPicture) SetWidth(v int) {
	((*C.WebPPicture)(wpp)).width = (C.int)(v)
}

func (wpp *WebPPicture) SetHeight(v int) {
	((*C.WebPPicture)(wpp)).height = (C.int)(v)
}

func (wpp WebPPicture) GetWidth() int {
	return int(((C.WebPPicture)(wpp)).width)
}

func (wpp WebPPicture) GetHeight() int {
	return int(((C.WebPPicture)(wpp)).height)
}

func (wpp *WebPPicture) SetUseArgb(v int) {
	((*C.WebPPicture)(wpp)).use_argb = (C.int)(v)
}

func WebpConfigInit(config *WebPConfig) int {
	return int(C.webp_config_init((*C.WebPConfig)(unsafe.Pointer(config))))
}

func WebpAnimEncoderOptionsInit(webPAnimEncoderOptions *WebPAnimEncoderOptions) int {
	return int(C.webp_anim_encoder_options_init((*C.WebPAnimEncoderOptions)(unsafe.Pointer(webPAnimEncoderOptions))))
}

func WebPPictureInit(webPPicture *WebPPicture) int {
	return int(C.webp_picture_init((*C.WebPPicture)(unsafe.Pointer(webPPicture))))
}

func WebPDataInit(webPData *WebPData) {
	C.webp_data_init((*C.WebPData)(unsafe.Pointer(webPData)))
}

func WebPValidateConfig(config *WebPConfig) int {
	return int(C.webp_validate_config((*C.WebPConfig)(unsafe.Pointer(config))))
}

func WebPPictureAlloc(webPPicture *WebPPicture) int {
	return int(C.webp_picture_alloc((*C.WebPPicture)(unsafe.Pointer(webPPicture))))
}

func WebPPictureCopy(webPPictureA *WebPPicture, webPPictureB *WebPPicture) int {
	return int(C.webp_picture_copy(
		(*C.WebPPicture)(unsafe.Pointer(webPPictureA)),
		(*C.WebPPicture)(unsafe.Pointer(webPPictureB)),
	))
}

func WebPAnimEncoderNew(width, height int, webPAnimEncoderOptions *WebPAnimEncoderOptions) *WebPAnimEncoder {
	return (*WebPAnimEncoder)(C.webp_anim_encoder_new(
		(C.int)(width),
		(C.int)(height),
		(*C.WebPAnimEncoderOptions)(unsafe.Pointer(webPAnimEncoderOptions)),
	))
}

func WebPAnimEncoderAdd(webPAnimEncoder *WebPAnimEncoder, webPPicture *WebPPicture, timestamp int, webPConfig *WebPConfig) int {
	return int(C.webp_anim_encoder_add(
		(*C.WebPAnimEncoder)(unsafe.Pointer(webPAnimEncoder)),
		(*C.WebPPicture)(unsafe.Pointer(webPPicture)),
		(C.int)(timestamp),
		(*C.WebPConfig)(unsafe.Pointer(webPConfig)),
	))
}

func WebPAnimEncoderAssemble(webPAnimEncoder *WebPAnimEncoder, webPData *WebPData) int {
	return int(C.webp_anim_encoder_assemble(
		(*C.WebPAnimEncoder)(unsafe.Pointer(webPAnimEncoder)),
		(*C.WebPData)(unsafe.Pointer(webPData)),
	))
}

func WebPAnimEncoderGetError(webPAnimEncoder *WebPAnimEncoder) string {
	return C.GoString(C.webp_anim_encoder_get_error(
		(*C.WebPAnimEncoder)(unsafe.Pointer(webPAnimEncoder)),
	))
}

func WebPMuxCreate(webPData *WebPData, copyData int) *WebPMux {
	return (*WebPMux)(C.webp_mux_create(
		(*C.WebPData)(unsafe.Pointer(webPData)),
		(C.int)(copyData),
	))
}

func WebPDataClear(webPData *WebPData) {
	C.webp_data_clear((*C.WebPData)(unsafe.Pointer(webPData)))
}

func WebPMuxGetAnimationParams(webPMux *WebPMux, webPMuxAnimParams *WebPMuxAnimParams) WebPMuxError {
	return (WebPMuxError)(C.webp_mux_get_animation_params(
		(*C.WebPMux)(unsafe.Pointer(webPMux)),
		(*C.WebPMuxAnimParams)(unsafe.Pointer(webPMuxAnimParams)),
	))
}

func WebPMuxSetAnimationParams(webPMux *WebPMux, webPMuxAnimParams *WebPMuxAnimParams) WebPMuxError {
	return (WebPMuxError)(C.webp_mux_set_animation_params(
		(*C.WebPMux)(unsafe.Pointer(webPMux)),
		(*C.WebPMuxAnimParams)(unsafe.Pointer(webPMuxAnimParams)),
	))
}

func WebPMuxSetChunk(webPMux *WebPMux, fourcc string, webPData *WebPData, copyData int) WebPMuxError {
	return (WebPMuxError)(C.webp_mux_set_chunk(
		(*C.WebPMux)(unsafe.Pointer(webPMux)),
		C.CString(fourcc),
		(*C.WebPData)(unsafe.Pointer(webPData)),
		(C.int)(copyData),
	))
}

func WebPMuxAssemble(webPMux *WebPMux, webPData *WebPData) WebPMuxError {
	return (WebPMuxError)(C.webp_mux_assemble(
		(*C.WebPMux)(unsafe.Pointer(webPMux)),
		(*C.WebPData)(unsafe.Pointer(webPData)),
	))
}

func WebPMuxDelete(webPMux *WebPMux) {
	C.webp_mux_delete((*C.WebPMux)(unsafe.Pointer(webPMux)))
}

func WebPPictureFree(webPPicture *WebPPicture) {
	C.webp_picture_free((*C.WebPPicture)(unsafe.Pointer(webPPicture)))
}

func WebPAnimEncoderDelete(webPAnimEncoder *WebPAnimEncoder) {
	C.webp_anim_encoder_delete((*C.WebPAnimEncoder)(unsafe.Pointer(webPAnimEncoder)))
}
