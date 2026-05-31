#include "resistor_color.h"

static const resistor_band_t band_array[] = {
    BLACK,BROWN,RED,ORANGE,YELLOW,GREEN,BLUE,VIOLET,GREY,WHITE
};

int color_code(resistor_band_t color){
    return (int) color;
}

const resistor_band_t* colors(){
    return band_array;
}