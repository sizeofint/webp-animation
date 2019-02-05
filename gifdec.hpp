#ifndef GIFDEC_HPP
#define GIFDEC_HPP

#include <gif_lib.h>

#ifdef __cplusplus
extern "C" {
#endif


GifFileType * gif(GifByteType*,int ,int *);


#ifdef __cplusplus
}
#endif

#endif