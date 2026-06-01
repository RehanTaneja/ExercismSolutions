#include "grains.h"

uint64_t square(uint8_t n){
    if (n==0){
        return 0;
    }
    uint64_t grain = 1;
    uint64_t i =1;
    while (i<n){
        grain *= 2;
        i++;
    }
    return grain;
}

uint64_t total(){
    uint64_t sum = 1;
    uint64_t grain = 1;
    uint64_t i = 1;
    while (i<=64){
        grain *= 2;
        sum += grain;
        i++;
    }
    return sum;
}