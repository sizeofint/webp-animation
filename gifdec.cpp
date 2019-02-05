#include "gifdec.hpp"

typedef struct  {
    GifByteType * image;
    int size;
    int fileHandle;
} ImageData;

int gif_file_input(GifFileType * gif, GifByteType * buf, int len) {

    ImageData* imageData = (ImageData*) gif->UserData;
    int readBytes = 0;

    
    for(int i = imageData->fileHandle; i<(imageData->fileHandle+len); i++) {
        if(imageData->size>i) {
            buf[readBytes] = (GifByteType)((GifByteType*)imageData->image)[i];
            readBytes++;
        } else {
            break;
        }
    }
    
    
    imageData->fileHandle += readBytes;
    
    
    return readBytes;
}

GifFileType * gif(GifByteType * gifByteType, int size, int *Error) {
    
    ImageData* imageData = new ImageData;

    imageData->image = gifByteType;
    imageData->size = size;
    imageData->fileHandle = 0;

    
    return DGifOpen(imageData,gif_file_input,Error);
}
