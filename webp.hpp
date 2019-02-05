#ifndef WEBP_ANIMATION_HPP
#define WEBP_ANIMATION_HPP

#include <webp/encode.h>
#include <webp/mux.h>

#ifdef __cplusplus
extern "C" {
#endif

int webp_config_init(WebPConfig*);
int webp_anim_encoder_options_init(WebPAnimEncoderOptions*);
int webp_picture_init(WebPPicture*);
void webp_data_init(WebPData*);
int webp_validate_config(const WebPConfig*);
int webp_picture_alloc(WebPPicture*);
int webp_picture_copy(const WebPPicture*, WebPPicture*);
WebPAnimEncoder* webp_anim_encoder_new(int,int, const WebPAnimEncoderOptions*);
int webp_anim_encoder_add(WebPAnimEncoder*, struct WebPPicture*, int,const struct WebPConfig*);
int webp_anim_encoder_assemble(WebPAnimEncoder*, WebPData*);
const char* webp_anim_encoder_get_error(WebPAnimEncoder*);
WebPMux* webp_mux_create(const WebPData*, int);
void webp_data_clear(WebPData*);
WebPMuxError webp_mux_get_animation_params(const WebPMux*, WebPMuxAnimParams*);
WebPMuxError webp_mux_set_animation_params(WebPMux*, const WebPMuxAnimParams*);
WebPMuxError webp_mux_set_chunk(WebPMux*, const char fourcc[4], const WebPData*,int);
WebPMuxError webp_mux_assemble(WebPMux*, WebPData*);
void webp_mux_delete(WebPMux*);
void webp_picture_free(WebPPicture*);
void webp_anim_encoder_delete(WebPAnimEncoder*);

#ifdef __cplusplus
}
#endif

#endif