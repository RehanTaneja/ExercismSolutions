#include "resistor_color_trio.h"

resistor_value_t color_code(resistor_band_t r[3]){
    uint16_t answer = r[0]*10 + r[1];
    int numZeros = (int)r[2];
    if (r[1]==BLACK){
        numZeros++;
        answer = answer / 10;
    }
    resistor_unit_t u;
    if (numZeros>=9){
        u = GIGAOHMS;
        numZeros-=9;
    }else if (numZeros>=6){
        u = MEGAOHMS;
        numZeros -=6;
    }else if (numZeros>=3){
        u = KILOOHMS;
        numZeros -=3;
    }else{
        u = OHMS;
    }
    for(int i=0;i<numZeros;i++){
        answer *= 10;
    }
    resistor_value_t a = {answer,u};
    return a;
}
