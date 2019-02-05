#include "webp.hpp"

int webp_config_init(WebPConfig* webPConfig) {
    return WebPConfigInitInternal(webPConfig, WEBP_PRESET_DEFAULT, 75.f,
                                WEBP_ENCODER_ABI_VERSION);
}

int webp_anim_encoder_options_init(WebPAnimEncoderOptions* enc_options) {
    return WebPAnimEncoderOptionsInitInternal(enc_options, WEBP_MUX_ABI_VERSION);
}

int webp_picture_init(WebPPicture* webPPicture) {
    return WebPPictureInit(webPPicture);
}

void webp_data_init(WebPData* webPData) {
    return WebPDataInit(webPData);
}

int webp_validate_config(const WebPConfig* webPConfig) {
 return WebPValidateConfig(webPConfig);
}

int webp_picture_alloc(WebPPicture* webPPicture) {
    return WebPPictureAlloc(webPPicture);
}

int webp_picture_copy(const WebPPicture* webPPictureA, WebPPicture* webPPictureB) {
    return WebPPictureCopy(webPPictureA,webPPictureB);
}

WebPAnimEncoder* webp_anim_encoder_new(int width,int height, const WebPAnimEncoderOptions* webPAnimEncoderOptions){
    return WebPAnimEncoderNew(width,height,webPAnimEncoderOptions);
}

int webp_anim_encoder_add(WebPAnimEncoder* webPAnimEncoder, struct WebPPicture* webPPicture, int timestamp,const struct WebPConfig* webPConfig){
    return WebPAnimEncoderAdd(webPAnimEncoder, webPPicture, timestamp,webPConfig);
}

int webp_anim_encoder_assemble(WebPAnimEncoder* webPAnimEncoder, WebPData* webPData) {
    return WebPAnimEncoderAssemble(webPAnimEncoder,webPData);
}

const char* webp_anim_encoder_get_error(WebPAnimEncoder* webPAnimEncoder) {
    return WebPAnimEncoderGetError(webPAnimEncoder);
}

WebPMux* webp_mux_create(const WebPData* webPData, int copyData) {
    return WebPMuxCreate(webPData,copyData);
}

void webp_data_clear(WebPData* webPData) {
    return WebPDataClear(webPData);
}

WebPMuxError webp_mux_get_animation_params(const WebPMux* webPMux, WebPMuxAnimParams* webPMuxAnimParams) {
    return WebPMuxGetAnimationParams(webPMux,webPMuxAnimParams);
}

WebPMuxError webp_mux_set_animation_params(WebPMux* webPMux, const WebPMuxAnimParams* webPMuxAnimParams) {
    return WebPMuxSetAnimationParams(webPMux,webPMuxAnimParams);
}

WebPMuxError webp_mux_set_chunk(WebPMux* webPMux, const char fourcc[4], const WebPData* webPData,int copyData) {
    return WebPMuxSetChunk(webPMux,fourcc,webPData,copyData);
}

WebPMuxError webp_mux_assemble(WebPMux* webPMux, WebPData* webPData) {
    return WebPMuxAssemble(webPMux, webPData);
}

void webp_mux_delete(WebPMux* webPMux) {
    return WebPMuxDelete(webPMux);
}

void webp_picture_free(WebPPicture* webPPicture) {
    return WebPPictureFree(webPPicture);
}


void webp_anim_encoder_delete(WebPAnimEncoder* webPAnimEncoder) {
    return WebPAnimEncoderDelete(webPAnimEncoder);
}
