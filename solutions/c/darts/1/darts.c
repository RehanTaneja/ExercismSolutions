#include "darts.h"

uint8_t score(coordinate_t landing_position){
    float d = (landing_position.x*landing_position.x)+(landing_position.y*landing_position.y);
    if (d<=1){
        return 10;
    }else if (d <= 25){
        return 5;
    }else if (d <= 100){
        return 1;
    }
    return 0;
}
